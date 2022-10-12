package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"
)

func main() {
	flag.Usage = func() { printUsage(flag.CommandLine.Output()) }
	config := ""
	output := ""
	system := runtime.GOOS
	target := runtime.GOARCH
	version_opt := false
	intermediate := false
	flag.StringVar(&config, "config", config, "configuration file for target")
	flag.StringVar(&config, "c", config, "configuration file for target")
	flag.BoolVar(&intermediate, "intermediate", false, "generate intermediate (single node) .abu files")
	flag.BoolVar(&intermediate, "i", false, "generate intermediate (single node) .abu files")
	flag.StringVar(&output, "output", output, "output file name for the compiled source")
	flag.StringVar(&output, "o", output, "output file name for the compiled source")
	flag.StringVar(&system, "system", system, "target operating system")
	flag.StringVar(&system, "s", system, "target operating system")
	flag.StringVar(&target, "target", target, "target language or architecture")
	flag.StringVar(&target, "t", target, "target language or architecture")
	flag.BoolVar(&version_opt, "version", false, "print version information and exit")
	flag.BoolVar(&version_opt, "v", false, "print version information and exit")
	flag.Parse()
	if version_opt {
		fmt.Println("abuc version " + version())
		fmt.Println(license)
		os.Exit(0)
	}
	source := flag.Arg(0)
	if intermediate {
		target = "abu"
	}
	switch target {
	case "go":
	case "abu":
	case "amd64":
		fallthrough
	case "arm64":
		fmt.Printf("target %s is not yet available\n", target)
		os.Exit(0)
	default:
		fmt.Fprintf(os.Stderr, "unknown target %s\n", target)
		os.Exit(1)
	}
	preprocs, symbolTable := preprocess(source, func(errs []error) {
		for _, err := range errs {
			fmt.Fprintln(os.Stderr, err.Error())
		}
		os.Exit(2)
	})
	compiler, err := makeCompileStrategy(system, target, output, config)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(4)
	}
	outcomes := make(chan struct {
		device string
		errs   []error
	})
	for d, ts := range preprocs {
		d := d
		ts := ts
		go func(out chan<- struct {
			device string
			errs   []error
		}) {
			o := struct {
				device string
				errs   []error
			}{
				device: d,
			}
			st, err := symbolTable.DeviceSymbolTable(d)
			if err != nil {
				o.errs = []error{err}
			} else {
				o.errs = compiler.compile(d, ts, st)
			}
			out <- o
		}(outcomes)
	}
	remaining := len(preprocs)
	ok := true
	for remaining > 0 {
		o := <-outcomes
		remaining--
		if len(o.errs) > 0 {
			ok = false
			fmt.Fprintln(os.Stderr, o.device+":")
			for _, err := range o.errs {
				fmt.Fprintln(os.Stderr, `	`, err.Error())
			}
		}
	}
	err = compiler.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	if err != nil || !ok {
		os.Exit(5)
	}
}

type flagInfo struct {
	short string
	long  string
	typ   string
	usage string
	def   string
}

// printUsage prints abuc usage on an io.Writer.
// It should be called after flag.Parse()
func printUsage(out io.Writer) error {
	_, err := fmt.Fprintf(out, "Usage of %s:\n", os.Args[0])
	if err != nil {
		return err
	}
	for _, f := range getFlagInfos() {
		if f.typ == "boolean" {
			_, err = fmt.Fprintf(out, "  -%s, --%s :\t%s\n\n", f.short, f.long, f.usage)
		} else {
			_, err = fmt.Fprintf(out, "  -%s, --%s %s\n    \t%s %s\n\n", f.short, f.long, f.typ, f.usage, f.def)
		}
		if err != nil {
			return err
		}
	}
	_, err = fmt.Fprintln(out, "Get help/Report bugs/Provide suggestions at: https://github.com/abu-lang/abuc")
	return err
}

// getFlagInfos returns a []flagInfo with the information
// on all the specifiable flags of abuc. The values of the
// returned slice are ordered lexicographically on the "short"
// fields. It should be called after flag.Parse().
func getFlagInfos() []flagInfo {
	longFlag := map[string]string{
		"c": "config",
		"i": "intermediate",
		"o": "output",
		"s": "system",
		"t": "target",
		"v": "version",
	}
	res := make([]flagInfo, 0, len(longFlag))
	flag.VisitAll(func(f *flag.Flag) {
		if len(f.Name) > 1 {
			return
		}
		t, u := flag.UnquoteUsage(f)
		if t == "" {
			t = "boolean"
		}
		var d string
		if t == "string" {
			d = fmt.Sprintf(" (default %q)", f.DefValue)
		} else {
			d = fmt.Sprintf(" (default %v)", f.DefValue)
		}
		res = append(res, flagInfo{
			short: f.Name,
			long:  longFlag[f.Name],
			typ:   t,
			usage: u,
			def:   d,
		})
	})
	return res
}
