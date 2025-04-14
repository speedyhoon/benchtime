package main

import (
	"bufio"
	"fmt"
	"github.com/speedyhoon/benchtime"
	"github.com/spf13/pflag"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	version := pflag.BoolP("version", "v", false, "Print version & exit.")
	decimalsQty := pflag.Uint8P("decimals", "d", benchtime.DecimalPlacesDefault, fmt.Sprintf("Decimal places. Maximum is %d.", benchtime.DecimalPlacesMax))
	// Sort table column options
	nameDesc := pflag.Bool("name", false, "Sort by column Name descending.")
	timeAverage := pflag.Bool("avg", false, "Sort by column Time Average.")
	timeMaximum := pflag.Bool("max", false, "Sort by column Time Maximum.")
	timeMinimum := pflag.Bool("min", false, "Sort by column Time Minimum.")
	timeTotal := pflag.Bool("tot", false, "Sort by column Time Total.")

	pflag.Usage = func() {
		_, _ = fmt.Fprintf(pflag.CommandLine.Output(), "benchtime summarises Go benchmark results into a table.\n\nUsage of %s: [files...]\n", os.Args[0])
		pflag.PrintDefaults()
		_, _ = fmt.Fprintln(pflag.CommandLine.Output(), "\nIf no files are specified as arguments then benchtime reads stdin for input.\nUsage example:\n\tgo test -bench . -benchmem -count=30 -shuffle=on | benchtime")
	}
	pflag.Parse()
	if *version {
		fmt.Println("0.3")
		return
	}

	paths := pflag.Args()
	if len(paths) == 0 {
		var stdin []byte
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			stdin = append(stdin, scanner.Bytes()...)
			stdin = append(stdin, byte('\n'))
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		sortCol := benchtime.SortNameAscending
		switch {
		case *nameDesc:
			sortCol = benchtime.SortNameDescending
		case *timeAverage:
			sortCol = benchtime.SortTimeAverage
		case *timeMaximum:
			sortCol = benchtime.SortTimeMaximum
		case *timeMinimum:
			sortCol = benchtime.SortTimeMinimum
		case *timeTotal:
			sortCol = benchtime.SortTimeTotal
		}

		fmt.Println(benchtime.Calculate(string(stdin), *decimalsQty, sortCol))
		return
	}

	for _, path := range paths {
		var err error
		path, err = cleanPath(path)
		if err != nil {
			fmt.Println(err)
			continue
		}

		var src []byte
		src, err = os.ReadFile(path)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(benchtime.Calculate(string(src), *decimalsQty, benchtime.SortNameAscending))
	}
}

func cleanPath(filePath string) (_ string, err error) {
	const homeDir = "~"
	if strings.HasPrefix(filePath, homeDir) && (runtime.GOOS == "linux" || runtime.GOOS == "darwin") {
		var usr *user.User
		usr, err = user.Current()
		if err != nil {
			return
		}
		filePath = strings.Replace(filePath, homeDir, usr.HomeDir, 1)
	}

	return filepath.Abs(filePath)
}
