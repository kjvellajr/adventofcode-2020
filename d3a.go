package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Printf("---%s---\n", filepath.Base(os.Args[0]))
	data, err := ioutil.ReadFile("input/d3.input")
	check(err)
	lines := strings.Split(string(data), "\n")
	var count = 0
	for i, line := range lines {
		if len(line) > 0 {
			var xpos = i * 3 % len(line)
			if line[xpos] == '#' {
				count++
			}
		}
	}
	fmt.Println(count)
}
