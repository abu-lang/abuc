{{- /* Copyright 2022 Massimo Comuzzo, Michele Pasqua and Marino Miculan
       SPDX-License-Identifier: 0BSD */ -}}

// Code generated by abuc version {{.Version}}.
package main

import (
    "flag"
    "strings"
    "time"

    "github.com/abu-lang/goabu"
    "github.com/abu-lang/goabu/communication"
    "github.com/abu-lang/goabu/config"
    {{range .Imports}}"{{.}}"
    {{end}}
)

var port_opt = flag.Int("port", 0, "set listening port")
var join_opt = flag.String("join", "", "set nodes (\"host:port,...\") to join")

func main() {
    flag.Parse()

    dev_state := {{.StateInitializer}}
    {{range .Resources}}dev_state{{.}}
    {{end}}
    knowledge_base := []string{
        {{range .Rules}}{{.}},
        {{end}}
    }

    lg_cfg := config.LogConfig{Level: config.LogInfo}
    dev_agent := communication.NewMemberlistAgent("{{.Id}}", *port_opt, lg_cfg, split_join_opt()...)
    device, err := goabu.NewExecuter(dev_state, knowledge_base, dev_agent, lg_cfg{{if .Invariant}}, `{{.Invariant}}`{{end}})
    if err != nil {
        panic(err)
    }
    for {
        time.Sleep({{.Tick}} * time.Second)
        device.Exec()
    }
}

// split_join_opt possibly separates the comma separated
// parts of *join_opt in a []string, trimming whitespaces.
// It returns []string{} if *join_opt == "".
func split_join_opt() []string {
    if *join_opt == "" {
        return []string{}
    }
    res := strings.Split(*join_opt, ",")
    for i, s := range res {
        res[i] = strings.TrimSpace(s)
    }
    return res
}
