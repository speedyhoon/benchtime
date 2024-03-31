package benchtime

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type info struct {
	os         string
	arch       string
	pkg        string
	cpu        string
	benchmarks map[string]benchmark
}

type benchmark struct {
	name        string
	runs        []run
	timeMaximum float64
	TimeMinimum float64
	timeMedian  float64
	timeAverage float64
	timeTotal   float64
}

type run struct {
	runs        uint64
	nanoseconds float64 // Per operation.
	bytes       uint64  // Bytes per operation.
	allocations uint64  // Allocations per operation.
}

func Calculate(benchmarkData string) {
	inf := info{benchmarks: map[string]benchmark{}}
	var maxNameLen int
	for _, line := range strings.Split(benchmarkData, "\n") {
		line = strings.TrimSpace(line)

		switch {
		case line == "":
			continue
		case strings.HasPrefix(line, "goos: "):
			inf.os = strings.TrimPrefix(line, "goos: ")
		case strings.HasPrefix(line, "goarch: "):
			inf.arch = strings.TrimPrefix(line, "goarch: ")
		case strings.HasPrefix(line, "pkg: "):
			inf.pkg = strings.TrimPrefix(line, "pkg: ")
		case strings.HasPrefix(line, "cpu: "):
			inf.cpu = strings.TrimPrefix(line, "cpu: ")
		case ignoreLine(line):
		default:
			data := strings.Split(line, "  ")
			var bench benchmark
			var r run
			var err error
			for _, item := range data {
				item = strings.TrimSpace(item)
				if item == "" {
					continue
				}

				switch {
				case strings.HasPrefix(item, "Benchmark"):
					bench.name = item
					maxNameLen = max(maxNameLen, len(item))
					err = nil
				case strings.HasSuffix(item, " ns/op"):
					item = strings.TrimSuffix(item, " ns/op")
					r.nanoseconds, err = strconv.ParseFloat(item, 64)
				case strings.HasSuffix(item, " B/op"):
					item = strings.TrimSuffix(item, " B/op")
					r.bytes, err = strconv.ParseUint(item, 10, 64)
				case strings.HasSuffix(item, " allocs/op"):
					item = strings.TrimSuffix(item, " allocs/op")
					r.allocations, err = strconv.ParseUint(item, 10, 64)
				default:
					r.runs, err = strconv.ParseUint(item, 10, 64)
				}
				if err != nil {
					log.Println(err)
				}
			}

			inf.Add(bench, r)
		}
	}

	fmt.Println("arch:", inf.arch)
	fmt.Println("os:", inf.os)
	fmt.Println("pkg:", inf.pkg)
	fmt.Println("cpu:", inf.cpu)
	Heading(maxNameLen)
	for _, benchmarks := range inf.benchmarks {
		benchmarks.Calc()

		// %-*s	= right padding spaces to `maxNameLen`.
		// %.3f	= truncate float 3 decimal places.
		fmt.Printf("%-*s\t%.3f\t%.3f\t%.3f\t%.3f\t%.3f\t%d\t%d\n", maxNameLen, benchmarks.name, benchmarks.timeMaximum, benchmarks.TimeMinimum, benchmarks.timeAverage, benchmarks.timeTotal, benchmarks.runs[0].nanoseconds, benchmarks.runs[0].bytes, benchmarks.runs[0].allocations)
	}
}

func ignoreLine(line string) bool {
	return strings.EqualFold(line, "PASS") ||
		strings.HasPrefix(line, "ok ") ||
		strings.HasPrefix(line, "-test.shuffle ")
}

func (inf *info) Add(bench benchmark, r run) {
	b, ok := inf.benchmarks[bench.name]
	if ok {
		b.runs = append(b.runs, r)
		b.timeMaximum = max(b.timeMaximum, r.nanoseconds)
		b.TimeMinimum = min(b.TimeMinimum, r.nanoseconds)
		inf.benchmarks[bench.name] = b
	} else {
		bench.runs = []run{r}
		bench.timeMaximum = r.nanoseconds
		bench.TimeMinimum = r.nanoseconds
		inf.benchmarks[bench.name] = bench
	}
}

func (bench *benchmark) Calc() {
	if len(bench.runs) == 0 {
		return
	}

	for _, r := range bench.runs {
		bench.timeTotal += r.nanoseconds
	}
	bench.timeAverage = bench.timeTotal / float64(len(bench.runs))
}

func Heading(l int) {
	fmt.Printf("%-*s\tmax\tmin\tavg\ttotal\n", l, " ")
}
