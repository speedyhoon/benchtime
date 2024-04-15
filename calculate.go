package benchtime

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

type info struct {
	os         string
	arch       string
	pkg        string
	cpu        string
	benchmarks map[string]*benchmark
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
	inf := info{benchmarks: map[string]*benchmark{}}
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

	var decimalWidth int
	var DecimalPlaces = 3
	if DecimalPlaces != 0 {
		decimalWidth = DecimalPlaces + 1
	}

	c := columnWidths{
		maximum:     max(len(ColMax), decimalWidth),
		minimum:     max(len(ColMin), decimalWidth),
		average:     max(len(ColAvg), decimalWidth),
		total:       max(len(ColTotal), decimalWidth),
		bytes:       len(ColBytesPerOp),
		allocations: len(ColAllocationsPerOp),
	}
	for _, benchmarks := range inf.benchmarks {
		benchmarks.Calc()
		c.ColumnSizes(*benchmarks, decimalWidth)
	}
	Heading(c, maxNameLen)

	for _, benchmarks := range inf.benchmarks {
		// %-*s	= right padding spaces to `maxNameLen`.
		// %*.*f	= left pad the float up to X spaces, then truncate float to X decimal places.
		fmt.Printf("%-*s  %*.*f  %*.*f  %*.*f  %*.*f  %*d  %*d\n",
			maxNameLen, benchmarks.name,
			c.maximum, DecimalPlaces, benchmarks.timeMaximum,
			c.minimum, DecimalPlaces, benchmarks.TimeMinimum,
			c.average, DecimalPlaces, benchmarks.timeAverage,
			c.total, DecimalPlaces, benchmarks.timeTotal,
			c.bytes, benchmarks.runs[0].bytes,
			c.allocations, benchmarks.runs[0].allocations,
		)
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
		inf.benchmarks[bench.name] = &bench
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

type columnWidths struct {
	total, minimum, maximum, average, bytes, allocations int
}

func (c *columnWidths) ColumnSizes(bench benchmark, decimalWidth int) {
	c.total = max(c.total, size(uint64(bench.timeTotal))+decimalWidth)
	c.minimum = max(c.minimum, size(uint64(bench.TimeMinimum))+decimalWidth)
	c.maximum = max(c.maximum, size(uint64(bench.timeMaximum))+decimalWidth)
	c.average = max(c.average, size(uint64(bench.timeAverage))+decimalWidth)
	c.bytes = max(c.bytes, size(bench.runs[0].bytes))
	c.allocations = max(c.allocations, size(bench.runs[0].allocations))
}

func Heading(c columnWidths, l int) {
	fmt.Printf("%-*s  %*s  %*s  %*s  %*s  %*s  %*s\n",
		l, " ",
		c.maximum, ColMax,
		c.minimum, ColMin,
		c.average, ColAvg,
		c.total, ColTotal,
		c.bytes, ColBytesPerOp,
		c.allocations, ColAllocationsPerOp,
	)
}

// Column names
var (
	ColMax              = "max"
	ColMin              = "min"
	ColAvg              = "avg"
	ColTotal            = "total"
	ColBytesPerOp       = "B/op"
	ColAllocationsPerOp = "allocs/op"
)

func size(f uint64) (count int) {
	if f < 10 {
		return 1
	}
	count = int(math.Log10(float64(f))) + 1
	return
}
