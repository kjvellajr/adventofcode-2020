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
	data, err := ioutil.ReadFile("input/d4.input")
	check(err)
	lines := strings.Split(string(data), "\n\n")

	fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"}
	required := []bool{true, true, true, true, true, true, true, false}
	regexs := make([]*regexp.Regexp, len(fields))
	var required_amount = 0
	for i, x := range fields {
		regexs[i] = regexp.MustCompile(fmt.Sprintf("%s:(?P<value>\\S+)", x))
		if required[i] {
			required_amount++
		}
	}

	var result = 0
	for _, line := range lines {
		if len(line) > 0 {
			var checks = 0
			for i, re := range regexs {
				res := re.FindStringSubmatch(line)
				if len(res) > 0 && required[i] {
					checks++
				}
			}
			if checks == required_amount {
				result++
			}
		}
	}
	fmt.Println(result)
}
