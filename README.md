# benchtime

A commandline tool to compare several benchmarks against each other. Handy for determining when function is the most performant.

If you want to compare the same benchmark between different versions use [`benchstat`](https://pkg.go.dev/golang.org/x/perf/cmd/benchstat) or [`benchmany`](https://pkg.go.dev/github.com/aclements/go-misc/benchmany) for running Go benchmarks across many git commits.

## Usage:

### Pipe

Pipe the standard output from `go test` directly into `benchtime`.

```shell
go test -bench . -benchmem -shuffle on -count 20 | benchtime -i
```

### Benchmark Data file

Generate some benchmark data to `bench.log` file:

```shell
go test -bench . -benchmem -shuffle on -count 20 > bench.log
```

Process the benchmark data in `bench.log`:

```shell
benchtime bench.log
```
