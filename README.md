# AbU compiler

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/abu-lang/abuc/blob/main/LICENSE)

A compiler for the AbU language.

`abuc` is a compiler translating programs written with AbU DSL [`abudsl`](https://github.com/abu-lang/abudsl) to different platforms. At the moment, the compiler supports the following target languages/architectures: *Go* code, *amd64* executables and *arm64* executables.

### Requirements
An [adequate](https://go.dev/doc/devel/release#policy) version of the Go language is required for installing and using `abuc` (it calls the Go language compiler in its execution).

### Installation
`abuc` can be installed with the go install command:
```
$ go install github.com/abu-lang/abuc@latest
```

## Usage
```
Usage: abuc [-ivh] [-o <output>] [-s <system>] [-t <target>] [-c <config>] [<source>]
  <source>                 filename of the source code to compile
  -o, --output <output>    output filename for the compiled source
  -s, --system <system>    target operating system [default: the operating system where abuc was installed]
  -t, --target <target>    target language or architecture [default: the architecture where abuc was installed]
  -c, --config <config>    configuration file for <target>
  -i, --intermediate       intermediate (single node) .abu files are generated
  -v, --version            print version information and exit
  -h, --help               print usage

Available values for <target>
  go       compile into 'Go' code
  amd64    compile into 'amd64' executable
  arm64    compile into 'arm64' executable

Available values for <system>
  run the following command to see a list of possible values for <system>:
  go tool dist list | grep 'amd64\|arm64' | sed -e 's/\/.*//' | sed -nr '$!N;/^(.*)\n\1$/!P;D'
```
If `<output>` ends with a directory separator, the output is placed in the directory indicated by `<output>` (the directory is created if it does not exist).

If `<source>` is not specified, the `abudsl` input is read from the standard input. If provided, the `<source>` option is to be specified as the last argument.

### Examples

Consider the following very simple `abudsl` program `test.abu`:
```
# AbU devices definition.

dev1 : "A first test device" {
    logical string name = "dev1"
    physical output string test = ""
    physical input boolean activate
} has notifyPeers

dev2 : "A second test device" {
    logical string name = "dev2"
    physical output string test = ""
    physical input boolean activate
} has notifyPeers

dev3 : "A third test device" {
    logical string name = "dev3"
    physical output string test = ""
    physical input boolean activate
} has notifyPeers

# AbU rules definition.

rule notifyPeers
    on activate
    for all (this.activate == true)
        do ext.test = this.name
```
We can compile such program to *Go* code by typing:
```
$ abuc -o testgo -t go test.abu
```
The command outputs three *Go* source code files `testgo-dev1.go`, `testgo-dev2.go` and `testgo-dev3.go` containing (the translation of) device *dev1* code, device *dev2* code and device *dev3* code, respectively. All files contain (the translation of) the rule *notifyPeers* code. <br><br>

Consider now the `abudsl` example program `raspberry-pi.abu`, that you can find in the [raspberry-pi](https://github.com/abu-lang/abudsl/tree/master/examples/raspberry-pi) directory of the `abudsl` repository. We can compile such program to *arm64* executables by typing:
```
$ abuc -o testraspberry -s linux -t arm64 -c config.json raspberry-pi.abu
```
in the `raspberry-pi` directory.
The command outputs two *arm64* executables `testraspberry-controls` and `testraspberry-wheel`, that can be directly deployed on the Raspberry Pi.
The executables accept the `-port` and `-join` command line options allowing the devices to join a system through the specification of a known member's address.

## License
`abuc` is [licensed](https://github.com/abu-lang/abuc/blob/main/LICENSE) under the Apache-2.0 License.

The binary executables generated by `abuc` shall be treated abiding the terms and conditions of the Apache-2.0 License as they statically include code under such license or under other compliant licenses.

Meanwhile, the non-binary files generated by `abuc`, if not otherwise stated, do not fall under the Apache-2.0 License and may be treated abiding to their original terms.
In particular, You can distribute the Go sources and the .abu intermediate files generated from Your original work under terms of Your choice.
