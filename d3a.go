package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("---day 3a---")
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
