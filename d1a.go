package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Printf("---%s---\n", filepath.Base(os.Args[0]))
	data, err := ioutil.ReadFile("input/d1.input")
	check(err)
	array_string := strings.Split(string(data), "\n")
	array := make([]int, len(array_string))
	for i := range array_string {
		array[i], _ = strconv.Atoi(array_string[i])
	}

	for i, foo := range array {
		for _, bar := range array[i+1:] {
			if foo+bar == 2020 {
				fmt.Printf("%d + %d = %d\n", foo, bar, foo+bar)
				fmt.Printf("%d * %d = %d\n", foo, bar, foo*bar)
				os.Exit(0)
			}
		}
	}
}
