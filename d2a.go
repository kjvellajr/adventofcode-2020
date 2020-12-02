package main

import (
	"fmt"
	"io/ioutil"
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
	fmt.Println("---day 2a---")
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
			x := strings.Count(password, key)
			if x >= min && x <= max {
				count++
			}
		}
	}
	fmt.Println(count)
}
