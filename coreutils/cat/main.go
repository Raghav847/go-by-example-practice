package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type Opts struct {
	Number    bool
	Squeeze   bool
	Filenames []string
}

func ParseOpts() Opts {
	var opts Opts

	flag.BoolVar(&opts.Number, "n", false, "Number all nonempty output lines, starting at 1")
	flag.BoolVar(&opts.Squeeze, "s", false, "Suppress repeated adjacent blank lines")
	flag.Parse()

	if flag.NArg() < 1 {
		log.Fatal("USAGE: cat [-n] [-s] file0 [fileN]")
		os.Exit(1)
	}

	opts.Filenames = flag.Args() //non flag cli arguments
	return opts
}

func main() {
	opts := ParseOpts()

	for _, filename := range opts.Filenames {
		in_file, err := os.ReadFile(filename)
		if err != nil {
			log.Fatalf("cannot open %s : %v", filename, err)
		}

		if opts.Number {
			count := 1
			for line := range strings.Lines(string(in_file)) {
				if opts.Squeeze && len(line) <= 1 {
					continue
				}

				fmt.Printf("%0.4d: %s", count, line)
				count++
			}
		} else {
			fmt.Print(string(in_file))
		}
	}
}
