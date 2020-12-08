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
	lines := strings.Split(string(data), "\n\n")
	reg, _ := regexp.Compile("[^a-zA-Z]+")

	result := 0
	for _, line := range lines {
		if len(line) > 0 {
			m := make(map[rune]bool)
			letters := reg.ReplaceAllString(line, "")
			for _, char := range letters {
				m[char] = true
			}
			result += len(m)
		}
	}

	fmt.Println(result)
}
