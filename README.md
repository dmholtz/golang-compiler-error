# Bug Report

This repository contains a simple program that triggers an internal compiler error.
The bug is caused by a combination of the following language features:

- Generics
- Interfaces (to restrict the valid types for instantiating a generic type)
- Nested structs

## Reproducing the bug

The bug has been observed on both go versions `go1.19` (latest) and `go1.18.5`.
However, the program has compiled without any errors using go version `go1.18.2`.

### go1.19 (confirmed)

#### Check go version

```bash
david@w520:~/repos/gobug$ go version
go version go1.19 linux/amd64
```

#### Build the program

```
david@w520:~/repos/gobug$ go build main.go
# command-line-arguments
./main.go:38:13: internal compiler error: panic: runtime error: invalid memory address or nil pointer dereference

Please file a bug report including a short program that triggers the error.
https://go.dev/issue/new
```

### go1.18.2 (NOT confirmed)

#### Check go version

```bash
david@w520:~/repos/gobug$ go version
go version go1.19 linux/amd64
```

#### Build the program

```
david@w520:~/repos/gobug$ go build main.go
# command-line-arguments
./main.go:38:13: internal compiler error: panic: runtime error: invalid memory address or nil pointer dereference

Please file a bug report including a short program that triggers the error.
https://go.dev/issue/new
```

## Observations

- The bug does not occur if no method of the inner struct `Edge` is called. Please refer to [sanitized1.go](sanitized1/sanitized1.go), where we avoid the call of the `To()` method.
- The bug does not occur if the struct `WeightedEdge` does not use any generics. Please refer to [sanitized2.go](sanitized2/sanitized2.go), where we replace the generic `WeightType` by a concrete type.
