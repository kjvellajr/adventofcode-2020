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

var regexp_byr = regexp.MustCompile(`byr:(\d{4})(?:\s|\z)`)
var regexp_iyr = regexp.MustCompile(`iyr:(\d{4})(?:\s|\z)`)
var regexp_eyr = regexp.MustCompile(`eyr:(\d{4})(?:\s|\z)`)
var regexp_hgt = regexp.MustCompile(`hgt:(\d+)(cm|in)(?:\s|\z)`)
var regexp_hcl = regexp.MustCompile(`hcl:(#[0-9a-f]{6})(?:\s|\z)`)
var regexp_ecl = regexp.MustCompile(`ecl:(amb|blu|brn|gry|grn|hzl|oth)(?:\s|\z)`)
var regexp_pid = regexp.MustCompile(`pid:(\d{9})(?:\s|\z)`)
var regexp_cid = regexp.MustCompile(`cid:(\S+)`)

func byr(x string) (string, string, bool) {
	res := regexp_byr.FindStringSubmatch(x)
	if len(res) > 1 {
		year, _ := strconv.Atoi(res[1])
		if year >= 1920 && year <= 2002 {
			return "byr", res[1], true
		}
	}
	return "byr", "", false
}

func iyr(x string) (string, string, bool) {
	res := regexp_iyr.FindStringSubmatch(x)
	if len(res) > 1 {
		year, _ := strconv.Atoi(res[1])
		if year >= 2010 && year <= 2020 {
			return "iyr", res[1], true
		}
	}
	return "iyr", "", false
}

func eyr(x string) (string, string, bool) {
	res := regexp_eyr.FindStringSubmatch(x)
	if len(res) > 1 {
		year, _ := strconv.Atoi(res[1])
		if year >= 2020 && year <= 2030 {
			return "eyr", res[1], true
		}
	}
	return "eyr", "", false
}

func hgt(x string) (string, string, bool) {
	res := regexp_hgt.FindStringSubmatch(x)
	if len(res) > 2 {
		height, _ := strconv.Atoi(res[1])
		unit := res[2]
		if unit == "cm" && height >= 150 && height <= 193 {
			return "hgt", res[1] + res[2], true
		}
		if unit == "in" && height >= 59 && height <= 76 {
			return "hgt", res[1] + res[2], true
		}
	}
	return "hgt", "", false
}

func hcl(x string) (string, string, bool) {
	res := regexp_hcl.FindStringSubmatch(x)
	if len(res) > 1 {
		return "hcl", res[1], true
	}
	return "hcl", "", false
}

func ecl(x string) (string, string, bool) {
	res := regexp_ecl.FindStringSubmatch(x)
	if len(res) > 1 {
		return "ecl", res[1], true
	}
	return "ecl", "", false
}

func pid(x string) (string, string, bool) {
	res := regexp_pid.FindStringSubmatch(x)
	if len(res) > 1 {
		return "pid", res[1], true
	}
	return "pid", "", false
}

func cid(x string) (string, string, bool) {
	res := regexp_cid.FindStringSubmatch(x)
	if len(res) > 1 {
		return "cid", res[1], true
	}
	return "cid", "", false
}

func main() {
	fmt.Printf("---%s---\n", filepath.Base(os.Args[0]))
	data, err := ioutil.ReadFile("input/d4.input")
	check(err)
	lines := strings.Split(string(data), "\n\n")

	fns := []func(string) (string, string, bool){byr, iyr, eyr, hgt, hcl, ecl, pid, cid}
	required := []bool{true, true, true, true, true, true, true, false}

	var required_amount = 0
	for _, required := range required {
		if required {
			required_amount++
		}
	}

	var result = 0
	for _, line := range lines {
		if len(line) > 0 {
			attributes := make(map[string]string)
			var checks = 0
			for j, fn := range fns {
				key, value, passed_rule := fn(line)
				attributes[key] = value
				if passed_rule && required[j] {
					checks++
				}
			}
			if checks == required_amount {
				result++
				//fmt.Printf("%d: %d\t%d\t%s:%s\t%s:%s\t%s:%s\t%s:%s\t%s:%s\t%s:%s\t%s:%s\t%s:%s\n", i, result, checks,
				//	"byr", attributes["byr"],
				//	"iyr", attributes["iyr"],
				//	"eyr", attributes["eyr"],
				//	"hgt", attributes["hgt"],
				//	"hcl", attributes["hcl"],
				//	"ecl", attributes["ecl"],
				//	"pid", attributes["pid"],
				//	"cid", attributes["cid"],
				//)
			}
		}
	}
	fmt.Println(result)
}
