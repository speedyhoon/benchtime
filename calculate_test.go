package benchtime

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestCalculate(t *testing.T) {
	stdIn, err := os.ReadFile("testdata.log")
	require.NoError(t, err)

	tests := []struct {
		name     string
		decimals uint
		expected string
	}{
		{decimals: 0, expected: `arch: amd64
os: linux
pkg: github.com/speedyhoon/jay/bench/uint16s
cpu: Intel(R) Core(TM)2 Duo CPU     T6400  @ 2.00GHz
                               max  min  avg  total  B/op  allocs/op
BenchmarkReadUint16s-2          65   59   61   2453     8          1
BenchmarkReadUint16s2-2         68   58   62   2475     8          1
BenchmarkReadUint16sUnsafe-2    61   55   57   2297     8          1
BenchmarkReadUint16sUnsafe2-2  198   54   71   2820     8          1
`},
		{decimals: 1, expected: `arch: amd64
os: linux
pkg: github.com/speedyhoon/jay/bench/uint16s
cpu: Intel(R) Core(TM)2 Duo CPU     T6400  @ 2.00GHz
                                 max   min   avg   total  B/op  allocs/op
BenchmarkReadUint16s-2          64.6  58.8  61.3  2452.8     8          1
BenchmarkReadUint16s2-2         67.9  58.0  61.9  2474.6     8          1
BenchmarkReadUint16sUnsafe-2    60.9  54.9  57.4  2297.4     8          1
BenchmarkReadUint16sUnsafe2-2  197.9  53.7  70.5  2820.1     8          1
`},
		{decimals: 2, expected: `arch: amd64
os: linux
pkg: github.com/speedyhoon/jay/bench/uint16s
cpu: Intel(R) Core(TM)2 Duo CPU     T6400  @ 2.00GHz
                                  max    min    avg    total  B/op  allocs/op
BenchmarkReadUint16s-2          64.59  58.76  61.32  2452.81     8          1
BenchmarkReadUint16s2-2         67.93  58.04  61.87  2474.63     8          1
BenchmarkReadUint16sUnsafe-2    60.85  54.92  57.43  2297.37     8          1
BenchmarkReadUint16sUnsafe2-2  197.90  53.74  70.50  2820.09     8          1
`},
		{decimals: 3, expected: `arch: amd64
os: linux
pkg: github.com/speedyhoon/jay/bench/uint16s
cpu: Intel(R) Core(TM)2 Duo CPU     T6400  @ 2.00GHz
                                   max     min     avg     total  B/op  allocs/op
BenchmarkReadUint16s-2          64.590  58.760  61.320  2452.810     8          1
BenchmarkReadUint16s2-2         67.930  58.040  61.866  2474.630     8          1
BenchmarkReadUint16sUnsafe-2    60.850  54.920  57.434  2297.370     8          1
BenchmarkReadUint16sUnsafe2-2  197.900  53.740  70.502  2820.090     8          1
`},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("decimals: %d", tt.decimals), func(t *testing.T) {
			assert.Equal(t, tt.expected, Calculate(string(stdIn), tt.decimals))
		})
	}
}
