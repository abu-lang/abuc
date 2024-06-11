// Copyright 2022 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

package compiler

import (
	"errors"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"sync"

	"github.com/abu-lang/abuc/internal/compiler/config"
	"github.com/abu-lang/abuc/preprocessor"
)

type machineCodeCompiler struct {
	goCompiler  compileStrategy
	ready       <-chan error
	tidyOnce    *sync.Once
	lockGoBuild *sync.Mutex
	commonCompileInfo
	goDest          string
	workDir         string
	additionalFiles []string
}

func makeMachineCodeCompiler(comm commonCompileInfo, cfgPath string) (compileStrategy, error) {
	conf, err := config.ReadFromFile(cfgPath)
	if err != nil {
		return nil, err
	}
	wd, err := os.MkdirTemp("", "abuc-")
	if err != nil {
		return nil, err
	}
	goDest := wd + string(os.PathSeparator)
	if comm.output != "" && !os.IsPathSeparator(comm.output[len(comm.output)-1]) {
		goDest = filepath.Join(wd, filepath.Base(comm.output))
	}
	gs, err := MakeCompileStrategy(comm.system, "go", goDest, cfgPath)
	if err != nil {
		os.Remove(wd)
		return nil, err
	}
	return machineCodeCompiler{
		commonCompileInfo: comm,
		additionalFiles:   basenames(conf.AdditionalFiles),
		goDest:            goDest,
		goCompiler:        gs,
		workDir:           wd,
		ready:             prepareWorkDir(wd, conf.AdditionalFiles),
		tidyOnce:          &sync.Once{},
		lockGoBuild:       &sync.Mutex{},
	}, nil
}

func (m machineCodeCompiler) Close() error {
	err := m.goCompiler.Close()
	if err != nil {
		return err
	}
	<-m.ready
	// RemoveAll fails silently with "" as argument
	// See https://github.com/golang/go/issues/28830
	if m.workDir == "" {
		return errors.New("invalid temporary working directory")
	}
	return os.RemoveAll(m.workDir)
}

func (m machineCodeCompiler) Compile(device string, stream preprocessor.TrivialStream, st preprocessor.DeviceSymbolTable) []error {
	// transpile from abudsl to Go
	errs := m.goCompiler.Compile(device, stream, st)
	if len(errs) > 0 {
		return errs
	}
	// wait for working directory preparations
	err := <-m.ready
	if err != nil {
		return []error{err}
	}
	// fetch dependencies
	m.tidyOnce.Do(func() {
		if !runCommand(m.workDir, "go", "mod", "tidy") {
			err = errors.New("error in fetching the required dependencies")
		}
	})
	if err != nil {
		return []error{err}
	}
	// construct go build command
	comp := exec.Command("go")
	comp.Dir = m.workDir
	comp.Env = append(os.Environ(), "GOOS="+m.system, "GOARCH="+m.target)
	comp.Args = make([]string, 0, 5+len(m.additionalFiles))
	dst, err := filepath.Abs(filepath.Dir(m.output))
	if err != nil {
		return []error{err}
	}
	dst = dst + string(os.PathSeparator) // (create directory if missing)
	comp.Args = append(comp.Args, "go", "build", "-o", dst, filepath.Base(outputFile(m.goDest, device, ".go")))
	comp.Args = append(comp.Args, m.additionalFiles...)
	// compile go sources
	m.lockGoBuild.Lock()
	_, err = comp.Output()
	m.lockGoBuild.Unlock()
	if err != nil {
		return []error{errors.New(string(err.(*exec.ExitError).Stderr))}
	}
	return nil
}

// basenames returns the base names of the argument's elements.
func basenames(paths []string) []string {
	res := make([]string, 0, len(paths))
	for _, p := range paths {
		res = append(res, filepath.Base(p))
	}
	return res
}

func prepareWorkDir(wd string, files []string) <-chan error {
	res := make(chan error)
	go func(done chan<- error) {
		var err error
		defer func() {
			for {
				done <- err
			}
		}()
		if !runCommand(wd, "go", "mod", "init", filepath.Base(wd)) {
			err = errors.New("could not initialize Go module")
			return
		}
		if len(files) > 0 {
			err = copyFiles(wd, files)
			if err != nil {
				return
			}
			// possibly fetch dependencies from additional files (not mandatory)
			runCommand(wd, "go", "mod", "tidy")
		}
		if !runCommand(wd, "go", "get", "github.com/abu-lang/goabu@"+goabuVersion) {
			err = errors.New("could not fetch GoAbU")
		}
	}(res)
	return res
}

// runCommand takes as argument a directory path and the arguments of a command.
// The command will be executed in the specified directory and the returned bool
// indicates wheter the command completed successfully or not.
func runCommand(workDir, name string, arg ...string) bool {
	cmd := exec.Command(name, arg...)
	cmd.Dir = workDir
	return nil == cmd.Run()
}

// copyFiles takes as argument a directory path and a slice of
// file paths and tries to copy the specified files in the indicated
// directory.
func copyFiles(dest string, files []string) error {
	for _, p := range files {
		r, err := os.Open(p)
		if err != nil {
			return err
		}
		w, err := os.Create(filepath.Join(dest, p))
		if err != nil {
			return err
		}
		_, err = io.Copy(w, r)
		if err != nil {
			return err
		}
	}
	return nil
}
