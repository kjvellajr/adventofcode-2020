package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
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
	data, err := ioutil.ReadFile("input/d2.input")
	check(err)
	re := regexp.MustCompile("([0-9]+)-([0-9]+) ([a-z]+): ([a-z]+)")
	array := strings.Split(string(data), "\n")
	var count = 0
	for _, line := range array {
		res := re.FindAllStringSubmatch(line, -1)
		for i := range res {
			min, _ := strconv.Atoi(res[i][1])
			max, _ := strconv.Atoi(res[i][2])
			key := res[i][3]
			password := res[i][4]
			var x = 0
			if string(password[min-1]) == key {
				x++
			}
			if string(password[max-1]) == key {
				x++
			}
			if x == 1 {
				count++
			}
		}
	}
	fmt.Println(count)
}
