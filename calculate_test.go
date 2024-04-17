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
