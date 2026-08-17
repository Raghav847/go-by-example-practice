package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, file := range counts {
		if len(file) > 1 {
			fmt.Printf("%d\t%s\t%s\n", len(files), line, files)
		}
	}

}

func countLines(f *os.File, counts map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()] = append(counts[input.Text()], f.Name())
	}
}
