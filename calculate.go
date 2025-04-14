package benchtime

import (
	"bytes"
	"fmt"
	"math"
	"sort"
	"strconv"
)

const (
	SortNameAscending uint8 = iota
	SortNameDescending
	SortTimeAverage
	SortTimeMaximum
	SortTimeMinimum
	SortTimeTotal
	DecimalPlacesDefault = 3
	DecimalPlacesMax     = 40
)

type info struct {
	os         string
	arch       string
	pkg        string
	cpu        string
	benchmarks []*benchmark
}

type benchmark struct {
	name        string
	runs        []run
	timeMaximum float64
	TimeMinimum float64
	timeAverage float64
	timeTotal   float64
}

type run struct {
	runs        uint64
	nanoseconds float64 // Per operation.
	bytes       uint64  // Bytes per operation.
	allocations uint64  // Allocations per operation.
}

func Calculate(benchmarkData []byte, decimalPlaces, sortColumn uint8) string {
	decimalPlaces = min(decimalPlaces, DecimalPlacesMax)

	inf := info{benchmarks: []*benchmark{}}
	var buf = bytes.NewBuffer(nil)
	var maxNameLen int
	for _, line := range bytes.Split(benchmarkData, []byte("\n")) {
		line = bytes.TrimSpace(line)

		switch {
		case len(line) == 0, ignoreLine(line):
			continue
		case bytes.HasPrefix(line, []byte("goos: ")):
			inf.os = string(bytes.TrimPrefix(line, []byte("goos: ")))
		case bytes.HasPrefix(line, []byte("goarch: ")):
			inf.arch = string(bytes.TrimPrefix(line, []byte("goarch: ")))
		case bytes.HasPrefix(line, []byte("pkg: ")):
			inf.pkg = string(bytes.TrimPrefix(line, []byte("pkg: ")))
		case bytes.HasPrefix(line, []byte("cpu: ")):
			inf.cpu = string(bytes.TrimPrefix(line, []byte("cpu: ")))
		default:
			data := bytes.Split(line, []byte("  "))
			var bench benchmark
			var r run
			var err error
			for _, item := range data {
				item = bytes.TrimSpace(item)
				if len(item) == 0 {
					continue
				}

				switch {
				case bytes.HasPrefix(item, []byte("Benchmark")):
					bench.name = string(item)
					maxNameLen = max(maxNameLen, len(item))
					err = nil
				case bytes.HasSuffix(item, []byte(" ns/op")):
					item = bytes.TrimSuffix(item, []byte(" ns/op"))
					r.nanoseconds, err = strconv.ParseFloat(string(item), 64)
				case bytes.HasSuffix(item, []byte(" B/op")):
					item = bytes.TrimSuffix(item, []byte(" B/op"))
					r.bytes, err = strconv.ParseUint(string(item), 10, 64)
				case bytes.HasSuffix(item, []byte(" allocs/op")):
					item = bytes.TrimSuffix(item, []byte(" allocs/op"))
					r.allocations, err = strconv.ParseUint(string(item), 10, 64)
				default:
					r.runs, err = strconv.ParseUint(string(item), 10, 64)
				}
				if err != nil {
					buf.WriteString(err.Error() + "\n")
				}
			}

			inf.Add(bench, r)
		}
	}

	buf.WriteString(fmt.Sprintf("arch: %s\nos: %s\npkg: %s\ncpu: %s\n", inf.arch, inf.os, inf.pkg, inf.cpu))

	var decimalWidth int
	if decimalPlaces != 0 {
		decimalWidth = int(decimalPlaces) + 1 // Add +1 for the width of the decimal points "."
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
	Heading(buf, c, maxNameLen)

	inf.Sort(sortColumn)
	for _, benchmarks := range inf.benchmarks {
		// %-*s	= right padding spaces to `maxNameLen`.
		// %*.*f	= left pad the float up to X spaces, then truncate float to X decimal places.
		buf.WriteString(fmt.Sprintf("%-*s  %*.*f  %*.*f  %*.*f  %*.*f  %*d  %*d\n",
			maxNameLen, benchmarks.name,
			c.maximum, decimalPlaces, benchmarks.timeMaximum,
			c.minimum, decimalPlaces, benchmarks.TimeMinimum,
			c.average, decimalPlaces, benchmarks.timeAverage,
			c.total, decimalPlaces, benchmarks.timeTotal,
			c.bytes, benchmarks.runs[0].bytes,
			c.allocations, benchmarks.runs[0].allocations,
		))
	}

	return buf.String()
}

func ignoreLine(line []byte) bool {
	return bytes.EqualFold(line, []byte("PASS")) ||
		bytes.HasPrefix(line, []byte("ok ")) ||
		bytes.HasPrefix(line, []byte("-test.shuffle ")) ||
		bytes.HasPrefix(line, []byte("Benchmarking "))
}

func (inf *info) Add(bench benchmark, r run) {
	for _, b := range inf.benchmarks {
		if b.name == bench.name {
			b.runs = append(b.runs, r)
			b.timeMaximum = max(b.timeMaximum, r.nanoseconds)
			b.TimeMinimum = min(b.TimeMinimum, r.nanoseconds)
			b.timeTotal += r.nanoseconds
			return
		}
	}

	bench.runs = []run{r}
	bench.timeMaximum = r.nanoseconds
	bench.TimeMinimum = r.nanoseconds
	bench.timeTotal = r.nanoseconds
	inf.benchmarks = append(inf.benchmarks, &bench)
}

func (bench *benchmark) Calc() {
	if len(bench.runs) == 0 {
		return
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

func Heading(b *bytes.Buffer, c columnWidths, l int) {
	b.WriteString(fmt.Sprintf("%-*s  %*s  %*s  %*s  %*s  %*s  %*s\n",
		l, " ",
		c.maximum, ColMax,
		c.minimum, ColMin,
		c.average, ColAvg,
		c.total, ColTotal,
		c.bytes, ColBytesPerOp,
		c.allocations, ColAllocationsPerOp,
	))
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

func (inf *info) Sort(column uint8) {
	switch column {
	default: // SortNameAscending
		sort.Slice(inf.benchmarks, func(i, j int) bool {
			return inf.benchmarks[i].name < inf.benchmarks[j].name
		})
	case SortNameDescending:
		sort.Slice(inf.benchmarks, func(i, j int) bool {
			return inf.benchmarks[i].name > inf.benchmarks[j].name
		})
	case SortTimeAverage:
		sort.Slice(inf.benchmarks, func(i, j int) bool {
			return inf.benchmarks[i].timeAverage < inf.benchmarks[j].timeAverage
		})
	case SortTimeMaximum:
		sort.Slice(inf.benchmarks, func(i, j int) bool {
			return inf.benchmarks[i].timeMaximum < inf.benchmarks[j].timeMaximum
		})
	case SortTimeMinimum:
		sort.Slice(inf.benchmarks, func(i, j int) bool {
			return inf.benchmarks[i].TimeMinimum < inf.benchmarks[j].TimeMinimum
		})
	case SortTimeTotal:
		sort.Slice(inf.benchmarks, func(i, j int) bool {
			return inf.benchmarks[i].timeTotal < inf.benchmarks[j].timeTotal
		})
	}
}
