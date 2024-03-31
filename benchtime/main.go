package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/speedyhoon/benchtime"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	version := flag.Bool("v", false, "Print version & exit.")
	flag.Usage = func() {
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s: [files...]\n", os.Args[0])
		flag.PrintDefaults()
		_, _ = fmt.Fprintln(flag.CommandLine.Output(), "\nIf no files are specified as arguments then benchtime reads stdin for input.")
	}
	flag.Parse()
	if *version {
		fmt.Println("0.1")
		return
	}

	paths := flag.Args()
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

		benchtime.Calculate(string(stdin))
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

		benchtime.Calculate(string(src))
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
