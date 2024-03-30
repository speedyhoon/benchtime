package main

const testData = `goos: linux
goarch: amd64
pkg: github.com/speedyhoon/jay/generate/testdata/string
cpu: Intel(R) Core(TM)2 Duo CPU     T6400  @ 2.00GHz
Benchmark_22_UnmarshalJ-2       164655706                6.819 ns/op           0 B/op          0 allocs/op
Benchmark_22a_UnmarshalJ-2      107506057               13.60 ns/op            0 B/op          0 allocs/op
Benchmark_22b_UnmarshalJ-2       4693540               226.9 ns/op             0 B/op          0 allocs/op
Benchmark_22c_UnmarshalJ-2        200798              8998 ns/op            1808 B/op         22 allocs/op
Benchmark_23_UnmarshalJ-2       188121460                6.476 ns/op           0 B/op          0 allocs/op
Benchmark_23a_UnmarshalJ-2      132235557                9.096 ns/op           0 B/op          0 allocs/op
Benchmark_23b_UnmarshalJ-2        311857              3488 ns/op            1808 B/op         22 allocs/op
Benchmark_23_UnmarshalJ3-2      183589686                6.256 ns/op           0 B/op          0 allocs/op
Benchmark_23a_UnmarshalJ3-2     127253947                9.144 ns/op           0 B/op          0 allocs/op
Benchmark_23b_UnmarshalJ3-2     23446874                45.96 ns/op            0 B/op          0 allocs/op
Benchmark_23c_UnmarshalJ3-2       328951              4184 ns/op            1808 B/op         22 allocs/op
PASS
ok      github.com/speedyhoon/jay/generate/testdata/string      21.700s
`
