package main

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
	benchmarks map[string][]benchmark
}

type benchmark struct {
	name        string
	runs        uint64
	nanoseconds float64 // Per operation.
	bytes       uint64  // Bytes per operation.
	allocations uint64  // Allocations per operation.
}

func main() {
	var inf info
	lines := strings.Split(testData, "\n")
	for _, line := range lines {
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
		case strings.EqualFold(line, "PASS"), strings.HasPrefix(line, "ok "):
			// Ignore line
		default:
			data := strings.Split(line, "  ")
			var bench benchmark
			var err error
			for _, item := range data {
				item = strings.TrimSpace(item)
				if item == "" {
					continue
				}

				switch {
				case strings.HasPrefix(item, "Benchmark"):
					bench.name = item
					err = nil
				case strings.HasSuffix(item, " ns/op"):
					item = strings.TrimSuffix(item, " ns/op")
					bench.nanoseconds, err = strconv.ParseFloat(item, 64)
				case strings.HasSuffix(item, " B/op"):
					item = strings.TrimSuffix(item, " B/op")
					bench.bytes, err = strconv.ParseUint(item, 10, 64)
				case strings.HasSuffix(item, " allocs/op"):
					item = strings.TrimSuffix(item, " allocs/op")
					bench.allocations, err = strconv.ParseUint(item, 10, 64)
				default:
					bench.runs, err = strconv.ParseUint(item, 10, 64)
				}
				if err != nil {
					log.Println(err)
				}
			}
			if len(inf.benchmarks) == 0 {
				inf.benchmarks = make(map[string][]benchmark)
			}
			inf.benchmarks[bench.name] = append(inf.benchmarks[bench.name], bench)
		}
	}

	fmt.Println(inf.os)
	fmt.Println(inf.arch)
	fmt.Println(inf.pkg)
	fmt.Println(inf.cpu)
	for _, benchmarks := range inf.benchmarks {
		fmt.Println(benchmarks[0].name, benchmarks[0].nanoseconds, benchmarks[0].bytes, benchmarks[0].allocations)
	}
}
