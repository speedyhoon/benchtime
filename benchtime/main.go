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
	const flagStdIn = "i"
	readStdIn := flag.Bool(flagStdIn, false, "read from standard input (stdin).")
	flag.Parse()

	if *readStdIn {
		var stdin []byte
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			stdin = append(stdin, scanner.Bytes()...)
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		benchtime.Calculate(string(stdin))
		return
	}

	paths := flag.Args()
	if len(paths) == 0 {
		_, _ = fmt.Fprintf(os.Stderr, "no paths specified or standard input flag -%s\n", flagStdIn)
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
