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
	var processed int
	modeOn := true
	for scanner.Scan() {
		buf := scanner.Text()
		processed, modeOn = parseLine(buf, modeOn)
		tot += processed
	}

	f.Close()
	fmt.Println(tot)
}

func parseLine(line string, modeOn bool) (int, bool) {
	tot := 0
	buf := line
	for buf != "" {
		buf1, one, mode1 := parseOne(buf, modeOn)
		buf = buf1
		modeOn = mode1
		if one == -1 {
			continue
		}
		tot += one
	}
	return tot, modeOn
}

func parseOne(buf string, modeOn bool) (string, int, bool) {
	//runtime.Breakpoint()
	prefix, rest := chopUntilToken(buf, "mul(")
	if rest == "" {
		return "", -1, modeOn
	}
	modeOn = parseMode(prefix, modeOn)
	if !modeOn {
		return rest, -1, modeOn
	}
	fmt.Println(crop(rest))
	rest, int1 := parseInt(rest)
	fmt.Println(int1, crop(rest))
	if int1 == -1 {
		return rest, -1, modeOn
	}
	if (rest == "") || (rest[0] != ',') {
		return rest, -1, modeOn
	}
	rest = rest[1:]
	rest, int2 := parseInt(rest)
	fmt.Println(int2, crop(rest))
	if int2 == -1 {
		return rest, -1, modeOn
	}
	if (rest == "") || (rest[0] != ')') {
		return rest, -1, modeOn
	}
	rest = rest[1:]
	return rest, int1 * int2, modeOn
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

func parseMode(prefix string, currentMode bool) bool {
	indexDo := strings.LastIndex(prefix, "do()")
	indexDont := strings.LastIndex(prefix, "don't()")
	if indexDo == indexDont {
		return currentMode
	}
	if indexDo > indexDont {
		return true
	}
	return false
}
