//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fpath := os.Args[1]
	f, err := os.Open(fpath)
	check(err)

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	tot := 0
	for scanner.Scan() {
		buf := scanner.Text()
		processed := parseLine(buf)
		tot += processed
	}

	f.Close()
	fmt.Println(tot)
}

func parseLine(line string) int {
	tot := 0
	buf := line
	for buf != "" {
		buf1, one := parseOne(buf)
		buf = buf1
		if one == -1 {
			continue
		}
		tot += one
	}
	return tot
}

func parseOne(buf string) (string, int) {
	//runtime.Breakpoint()
	_, rest := chopUntilToken(buf, "mul(")
	if rest == "" {
		return "", -1
	}
	fmt.Println(crop(rest))
	rest, int1 := parseInt(rest)
	fmt.Println(int1, crop(rest))
	if int1 == -1 {
		return rest, -1
	}
	if (rest == "") || (rest[0] != ',') {
		return rest, -1
	}
	rest = rest[1:]
	rest, int2 := parseInt(rest)
	fmt.Println(int2, crop(rest))
	if int2 == -1 {
		return rest, -1
	}
	if (rest == "") || (rest[0] != ')') {
		return rest, -1
	}
	rest = rest[1:]
	return rest, int1 * int2
}

func chopUntilToken(line string, token string) (string, string) {
	loc := strings.Index(line, token)
	if loc == -1 {
		return line, ""
	}
	return line[:loc], line[loc+len(token):]
}

func parseInt(line string) (string, int) {
	digitEnds := -1
	for i, v := range line {
		if (!unicode.IsDigit(v)) || i > 3 {
			break
		}
		digitEnds = i
	}
	if digitEnds == -1 {
		return line, -1
	}
	num, _ := strconv.ParseInt(line[:digitEnds+1], 10, 64)
	num32 := int(num)
	return line[digitEnds+1:], num32
}

func crop(s string) string {
	if len(s) >= 10 {
		return s[:10]
	}
	return s
}
