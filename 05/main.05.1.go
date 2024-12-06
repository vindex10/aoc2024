//go:build ignore

package main

import (
	"bufio"
	"fmt"
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
	fname := os.Args[1]
	f, err := os.Open(fname)
	defer f.Close()
	check(err)
	scanner := bufio.NewScanner(f)
	var rules [][]int
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			break
		}
		partsStr := strings.SplitN(line, "|", 2)
		var parts []int
		parsed, _ := strconv.ParseInt(partsStr[0], 10, 64)
		parts = append(parts, int(parsed))
		parsed, _ = strconv.ParseInt(partsStr[1], 10, 64)
		parts = append(parts, int(parsed))
		rules = append(rules, parts)
	}
	var lines [](map[int]int)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		partsStr := strings.Split(line, ",")
		parts := make(map[int]int)
		for i, one := range partsStr {
			parsed, _ := strconv.ParseInt(one, 10, 64)
			parts[int(parsed)] = i
		}
		lines = append(lines, parts)
	}
	tot := 0
	for _, line := range lines {
		if validateLine(rules, line) {
			tot += midpoint(line)
		}
	}

	fmt.Println(tot)
}

func validateLine(rules [][]int, line map[int]int) bool {
	for _, rule := range rules {
		v1, e1 := line[rule[0]]
		v2, e2 := line[rule[1]]
		if !e1 || !e2 {
			continue
		}
		if v1 >= v2 {
			return false
		}
	}
	return true
}

func midpoint(line map[int]int) int {
	length := len(line)
	mid := length / 2
	for v, idx := range line {
		if idx == mid {
			return v
		}
	}
	panic("error")
}
