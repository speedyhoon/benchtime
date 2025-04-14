package benchtime

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestCalculate2NameAsc(t *testing.T) {
	stdIn, err := os.ReadFile("testdata2.log")
	require.NoError(t, err)

	tests := []struct {
		name     string
		decimals uint8
		expected string
	}{
		{decimals: 0, expected: `arch: amd64
os: linux
pkg: github.com/speedyhoon/jay/generate/testdata/boolDef
cpu: Intel(R) Core(TM)2 Quad CPU    Q8300  @ 2.50GHz
                                 max    min    avg    total   B/op  allocs/op
Benchmark23_Marshal_Gob-4      32085  18190  20729   621882   2768         19
Benchmark23_Marshal_JSON-4      2966   2325   2543    76298    376          2
Benchmark23_Marshal_Jay-4         64     46     49     1482      3          1
Benchmark23_Unmarshal_Gob-4   155635  76386  89071  2672116  10448        238
Benchmark23_Unmarshal_JSON-4   19506  10809  12146   364368    216          4
Benchmark23_Unmarshal_Jay-4       33     18     20      606      0          0
`},
		{decimals: 1, expected: `arch: amd64
os: linux
pkg: github.com/speedyhoon/jay/generate/testdata/boolDef
cpu: Intel(R) Core(TM)2 Quad CPU    Q8300  @ 2.50GHz
                                   max      min      avg      total   B/op  allocs/op
Benchmark23_Marshal_Gob-4      32085.0  18190.0  20729.4   621882.0   2768         19
Benchmark23_Marshal_JSON-4      2966.0   2325.0   2543.3    76298.0    376          2
Benchmark23_Marshal_Jay-4         64.1     46.0     49.4     1481.8      3          1
Benchmark23_Unmarshal_Gob-4   155635.0  76386.0  89070.5  2672116.0  10448        238
Benchmark23_Unmarshal_JSON-4   19506.0  10809.0  12145.6   364368.0    216          4
Benchmark23_Unmarshal_Jay-4       33.4     17.7     20.2      606.0      0          0
`},
		{decimals: 2, expected: `arch: amd64
os: linux
pkg: github.com/speedyhoon/jay/generate/testdata/boolDef
cpu: Intel(R) Core(TM)2 Quad CPU    Q8300  @ 2.50GHz
                                    max       min       avg       total   B/op  allocs/op
Benchmark23_Marshal_Gob-4      32085.00  18190.00  20729.40   621882.00   2768         19
Benchmark23_Marshal_JSON-4      2966.00   2325.00   2543.27    76298.00    376          2
Benchmark23_Marshal_Jay-4         64.07     46.03     49.39     1481.85      3          1
Benchmark23_Unmarshal_Gob-4   155635.00  76386.00  89070.53  2672116.00  10448        238
Benchmark23_Unmarshal_JSON-4   19506.00  10809.00  12145.60   364368.00    216          4
Benchmark23_Unmarshal_Jay-4       33.40     17.70     20.20      606.01      0          0
`},
		{decimals: 3, expected: `arch: amd64
os: linux
pkg: github.com/speedyhoon/jay/generate/testdata/boolDef
cpu: Intel(R) Core(TM)2 Quad CPU    Q8300  @ 2.50GHz
                                     max        min        avg        total   B/op  allocs/op
Benchmark23_Marshal_Gob-4      32085.000  18190.000  20729.400   621882.000   2768         19
Benchmark23_Marshal_JSON-4      2966.000   2325.000   2543.267    76298.000    376          2
Benchmark23_Marshal_Jay-4         64.070     46.030     49.395     1481.850      3          1
Benchmark23_Unmarshal_Gob-4   155635.000  76386.000  89070.533  2672116.000  10448        238
Benchmark23_Unmarshal_JSON-4   19506.000  10809.000  12145.600   364368.000    216          4
Benchmark23_Unmarshal_Jay-4       33.400     17.700     20.200      606.010      0          0
`},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("decimals: %d", tt.decimals), func(t *testing.T) {
			assert.Equal(t, tt.expected, Calculate(stdIn, tt.decimals, SortNameAscending))
		})
	}
}

func TestCalculate2NameDsc(t *testing.T) {
	stdIn, err := os.ReadFile("testdata2.log")
	require.NoError(t, err)

	tests := []struct {
		name     string
		decimals uint8
		expected string
	}{
		{decimals: 0, expected: `arch: amd64
os: linux
pkg: github.com/speedyhoon/jay/generate/testdata/boolDef
cpu: Intel(R) Core(TM)2 Quad CPU    Q8300  @ 2.50GHz
                                 max    min    avg    total   B/op  allocs/op
Benchmark23_Unmarshal_Jay-4       33     18     20      606      0          0
Benchmark23_Unmarshal_JSON-4   19506  10809  12146   364368    216          4
Benchmark23_Unmarshal_Gob-4   155635  76386  89071  2672116  10448        238
Benchmark23_Marshal_Jay-4         64     46     49     1482      3          1
Benchmark23_Marshal_JSON-4      2966   2325   2543    76298    376          2
Benchmark23_Marshal_Gob-4      32085  18190  20729   621882   2768         19
`},
		{decimals: 1, expected: `arch: amd64
os: linux
pkg: github.com/speedyhoon/jay/generate/testdata/boolDef
cpu: Intel(R) Core(TM)2 Quad CPU    Q8300  @ 2.50GHz
                                   max      min      avg      total   B/op  allocs/op
Benchmark23_Unmarshal_Jay-4       33.4     17.7     20.2      606.0      0          0
Benchmark23_Unmarshal_JSON-4   19506.0  10809.0  12145.6   364368.0    216          4
Benchmark23_Unmarshal_Gob-4   155635.0  76386.0  89070.5  2672116.0  10448        238
Benchmark23_Marshal_Jay-4         64.1     46.0     49.4     1481.8      3          1
Benchmark23_Marshal_JSON-4      2966.0   2325.0   2543.3    76298.0    376          2
Benchmark23_Marshal_Gob-4      32085.0  18190.0  20729.4   621882.0   2768         19
`},
		{decimals: 2, expected: `arch: amd64
os: linux
pkg: github.com/speedyhoon/jay/generate/testdata/boolDef
cpu: Intel(R) Core(TM)2 Quad CPU    Q8300  @ 2.50GHz
                                    max       min       avg       total   B/op  allocs/op
Benchmark23_Unmarshal_Jay-4       33.40     17.70     20.20      606.01      0          0
Benchmark23_Unmarshal_JSON-4   19506.00  10809.00  12145.60   364368.00    216          4
Benchmark23_Unmarshal_Gob-4   155635.00  76386.00  89070.53  2672116.00  10448        238
Benchmark23_Marshal_Jay-4         64.07     46.03     49.39     1481.85      3          1
Benchmark23_Marshal_JSON-4      2966.00   2325.00   2543.27    76298.00    376          2
Benchmark23_Marshal_Gob-4      32085.00  18190.00  20729.40   621882.00   2768         19
`},
		{decimals: 3, expected: `arch: amd64
os: linux
pkg: github.com/speedyhoon/jay/generate/testdata/boolDef
cpu: Intel(R) Core(TM)2 Quad CPU    Q8300  @ 2.50GHz
                                     max        min        avg        total   B/op  allocs/op
Benchmark23_Unmarshal_Jay-4       33.400     17.700     20.200      606.010      0          0
Benchmark23_Unmarshal_JSON-4   19506.000  10809.000  12145.600   364368.000    216          4
Benchmark23_Unmarshal_Gob-4   155635.000  76386.000  89070.533  2672116.000  10448        238
Benchmark23_Marshal_Jay-4         64.070     46.030     49.395     1481.850      3          1
Benchmark23_Marshal_JSON-4      2966.000   2325.000   2543.267    76298.000    376          2
Benchmark23_Marshal_Gob-4      32085.000  18190.000  20729.400   621882.000   2768         19
`},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("decimals: %d", tt.decimals), func(t *testing.T) {
			assert.Equal(t, tt.expected, Calculate(stdIn, tt.decimals, SortNameDescending))
		})
	}
}

func TestCalculate2Maximum(t *testing.T) {
	stdIn, err := os.ReadFile("testdata2.log")
	require.NoError(t, err)

	tests := []struct {
		name     string
		decimals uint8
		expected string
	}{
		{decimals: 0, expected: `arch: amd64
os: linux
pkg: github.com/speedyhoon/jay/generate/testdata/boolDef
cpu: Intel(R) Core(TM)2 Quad CPU    Q8300  @ 2.50GHz
                                 max    min    avg    total   B/op  allocs/op
Benchmark23_Unmarshal_Jay-4       33     18     20      606      0          0
Benchmark23_Marshal_Jay-4         64     46     49     1482      3          1
Benchmark23_Marshal_JSON-4      2966   2325   2543    76298    376          2
Benchmark23_Unmarshal_JSON-4   19506  10809  12146   364368    216          4
Benchmark23_Marshal_Gob-4      32085  18190  20729   621882   2768         19
Benchmark23_Unmarshal_Gob-4   155635  76386  89071  2672116  10448        238
`},
		{decimals: 1, expected: `arch: amd64
os: linux
pkg: github.com/speedyhoon/jay/generate/testdata/boolDef
cpu: Intel(R) Core(TM)2 Quad CPU    Q8300  @ 2.50GHz
                                   max      min      avg      total   B/op  allocs/op
Benchmark23_Unmarshal_Jay-4       33.4     17.7     20.2      606.0      0          0
Benchmark23_Marshal_Jay-4         64.1     46.0     49.4     1481.8      3          1
Benchmark23_Marshal_JSON-4      2966.0   2325.0   2543.3    76298.0    376          2
Benchmark23_Unmarshal_JSON-4   19506.0  10809.0  12145.6   364368.0    216          4
Benchmark23_Marshal_Gob-4      32085.0  18190.0  20729.4   621882.0   2768         19
Benchmark23_Unmarshal_Gob-4   155635.0  76386.0  89070.5  2672116.0  10448        238
`},
		{decimals: 2, expected: `arch: amd64
os: linux
pkg: github.com/speedyhoon/jay/generate/testdata/boolDef
cpu: Intel(R) Core(TM)2 Quad CPU    Q8300  @ 2.50GHz
                                    max       min       avg       total   B/op  allocs/op
Benchmark23_Unmarshal_Jay-4       33.40     17.70     20.20      606.01      0          0
Benchmark23_Marshal_Jay-4         64.07     46.03     49.39     1481.85      3          1
Benchmark23_Marshal_JSON-4      2966.00   2325.00   2543.27    76298.00    376          2
Benchmark23_Unmarshal_JSON-4   19506.00  10809.00  12145.60   364368.00    216          4
Benchmark23_Marshal_Gob-4      32085.00  18190.00  20729.40   621882.00   2768         19
Benchmark23_Unmarshal_Gob-4   155635.00  76386.00  89070.53  2672116.00  10448        238
`},
		{decimals: 3, expected: `arch: amd64
os: linux
pkg: github.com/speedyhoon/jay/generate/testdata/boolDef
cpu: Intel(R) Core(TM)2 Quad CPU    Q8300  @ 2.50GHz
                                     max        min        avg        total   B/op  allocs/op
Benchmark23_Unmarshal_Jay-4       33.400     17.700     20.200      606.010      0          0
Benchmark23_Marshal_Jay-4         64.070     46.030     49.395     1481.850      3          1
Benchmark23_Marshal_JSON-4      2966.000   2325.000   2543.267    76298.000    376          2
Benchmark23_Unmarshal_JSON-4   19506.000  10809.000  12145.600   364368.000    216          4
Benchmark23_Marshal_Gob-4      32085.000  18190.000  20729.400   621882.000   2768         19
Benchmark23_Unmarshal_Gob-4   155635.000  76386.000  89070.533  2672116.000  10448        238
`},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("decimals: %d", tt.decimals), func(t *testing.T) {
			assert.Equal(t, tt.expected, Calculate(stdIn, tt.decimals, SortTimeMaximum))
		})
	}
}
