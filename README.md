# AbU compiler
A compiler for the AbU language.

`abuc` is a compiler translating programs written in the AbU DSL `abudsl` to different platforms. At the moment, the compiler supports the following target platforms: *Golang* code, *amd64* executables and *arm64* executables.

## Usage
```
Usage: abuc [-o <output>] -t <target> <source>
  <source>                 filename of the source code to compile
  -o, --output <output>    output filename for the compiled source
  -t, --target <target>    target language or architecture

Available options for <target>:
  go       compile into 'Golang' code
  amd64    compile into 'amd64' executable
  arm64    compile into 'arm64' executable
```

As an example, consider the following snippet of an `abudsl` program `test.abu`:
```
# AbU devices definition.

dev1 : "A first test device" {
    # Resources declaration.
    ...
}

dev2 : "A second test device" {
    # Resources declaration.
    ...
}

dev3 : "A third test device" {
    # Resources declaration.
    ...
}
```
We can compile such program to *Golang* code by typing: `abuc -o testgo -t go test.abu`. <br>
The command outputs three *Golang* source code file `testgo-dev1.go`, `testgo-dev2.go` and `testgo-dev3.go`.
