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

	var max = 0
	seatIds := make([]int, len(lines))
	for i, line := range lines {
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
			seatIds[i] = seatId
			if max < seatId {
				max = seatId
			}
		}
	}

	seats := make([]int, max+1)
	for _, id := range seatIds {
		seats[id] = 1
	}

	var result = 0
	for i, x := range seats {
		if x == 0 && i > 0 && seats[i-1] == 1 && len(seats) > i+1 && seats[i+1] == 1 {
			result = i
			break
		}
	}
	fmt.Println(result)
}
