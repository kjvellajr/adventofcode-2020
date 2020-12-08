package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Printf("---%s---\n", filepath.Base(os.Args[0]))
	data, err := ioutil.ReadFile("input/d6.input")
	check(err)

	reg, _ := regexp.Compile(`\n$`)
	dataNoEOF := reg.ReplaceAllString(string(data), "")

	groups := strings.Split(dataNoEOF, "\n\n")

	result := 0
	for _, group := range groups {
		if len(group) > 0 {
			persons := strings.Split(group, "\n")
			m := make(map[rune]int)
			for i, person := range persons {
				for _, answer := range person {
					m[answer]++
					if i == len(persons)-1 && m[answer] == len(persons) {
						result++
					}
				}
			}
		}
	}

	fmt.Println(result)
}
