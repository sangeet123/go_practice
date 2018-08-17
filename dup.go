package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	m := make(map[string]int)

	if len(files) < 1 {
		countDuplicate(os.Stdin, m)
	} else {
		for _, file := range files {
			f, err := os.Open(file)

			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v", err)
				continue
			}
			countDuplicate(f, m)
		}
	}

	for k, v := range m {
		if v > 1 {
			fmt.Printf("%s:<%d>\n", k, v)
		}
	}
}

func countDuplicate(f *os.File, m map[string]int) {
	b := bufio.NewScanner(f)
	for b.Scan() {
		m[b.Text()]++
	}
}
