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
	xslope := [...]int{1, 3, 5, 7, 1}
	yslope := [...]int{1, 1, 1, 1, 2}
	var result = 1
	for i, x := range xslope {
		result *= trees(lines, x, yslope[i])
	}
	fmt.Println(result)
}

func trees(lines []string, right int, down int) int {
	var x = 0
	var count = 0
	for i, line := range lines {
		if len(line) > 0 && i%down == 0 {
			if line[x*right%len(line)] == '#' {
				count++
			}
			x++
		}
	}
	fmt.Printf("%d, %d = %d\n", right, down, count)
	return count
}
