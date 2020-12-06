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
	data, err := ioutil.ReadFile("input/d5.input")
	check(err)
	lines := strings.Split(string(data), "\n")

	modifier := []int{64, 32, 16, 8, 4, 2, 1, 4, 2, 1}

	var result = 0
	for _, line := range lines {
		if len(line) > 0 {
			var row = 0
			var column = 0
			for j, x := range line {
				if j < 7 {
					if x == 'B' {
						row += modifier[j]
					}
				} else {
					if x == 'R' {
						column += modifier[j]
					}
				}
			}
			seatId := row*8 + column
			if result < seatId {
				result = seatId
				//fmt.Printf("%d:{ row:%d column:%d seatId:%d }\n", i, row, column, row*8+column)
			}
		}
	}
	fmt.Println(result)
}
