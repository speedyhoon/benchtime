# benchtime

A commandline tool to compare several [Go](https://go.dev/) benchmarks against each other. Handy for determining which function is the most performant.

If you want to compare the same benchmark between different versions use [`benchstat`](https://pkg.go.dev/golang.org/x/perf/cmd/benchstat) or [`benchmany`](https://pkg.go.dev/github.com/aclements/go-misc/benchmany) for running Go benchmarks across many git commits.

## Usage:

### Pipe

Pipe the standard output from `go test` directly into `benchtime`.

```shell
go test -bench . -benchmem -shuffle on -count 20 | benchtime
```

### Benchmark Data file

1. Generate some benchmark data to `bench.log` file:
	```shell
	go test -bench . -benchmem -shuffle on -count 20 > bench.log
	```
2. Then process the benchmark data in `bench.log`:
	```shell
	benchtime bench.log
	```

## Install

```shell
go install github.com/speedyhoon/benchtime/benchtime@latest
```
