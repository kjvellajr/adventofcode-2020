package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("---day 1b---")
	data, err := ioutil.ReadFile("input/d1.input")
	check(err)
	array_string := strings.Split(string(data), "\n")
	array := make([]int, len(array_string))
	for i := range array_string {
		array[i], _ = strconv.Atoi(array_string[i])
	}

	for i, foo := range array {
		for j, bar := range array[i+1:] {
			for _, baz := range array[j+1:] {
				if foo+bar+baz == 2020 {
					fmt.Printf("%d + %d + %d = %d\n", foo, bar, baz, foo+bar+baz)
					fmt.Printf("%d * %d * %d = %d\n", foo, bar, baz, foo*bar*baz)
					os.Exit(0)
				}
			}
		}
	}
}
