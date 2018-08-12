package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	files := os.Args[1:]
	m := make(map[string]int)

	for _, file := range files {
		data, err := ioutil.ReadFile(file)

		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2:%v", err)
		}

		for _, line := range strings.Split(string(data), "\n") {
			m[line]++
		}
	}

	for k, v := range m {
		if v > 1 {
			fmt.Printf("%s:<%d>\n", k, v)
		}
	}
}
