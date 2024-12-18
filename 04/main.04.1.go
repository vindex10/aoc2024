//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
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

	var lines []string
	for scanner.Scan() {
		buf := scanner.Text()
		buf = strings.TrimSpace(buf)
		lines = append(lines, buf)
	}
	f.Close()

	tot := 0
	tot += countXmas(lines)
	lines90 := rotate90(lines)
	tot += countXmas(lines90)
	lines45 := rotate45(lines)
	tot += countXmas(lines45)
	lines45R := rotate45R(lines)
	tot += countXmas(lines45R)

	fmt.Println(tot)
}

func countXmas(lines []string) int {
	res := 0
	for _, line := range lines {
		res += countXmasOne(line)
	}
	return res
}

func countXmasOne(line string) int {
	var res int
	re := regexp.MustCompile("XMAS")
	matches := re.FindAllString(line, -1)
	if matches != nil {
		res += len(matches)
	}
	re = regexp.MustCompile("SAMX")
	matches = re.FindAllString(line, -1)
	if matches != nil {
		res += len(matches)
	}
	return res
}

func rotate90(lines []string) []string {
	var res []string
	length := len(lines)
	if length == 0 {
		return res
	}

	resArr := make([][]rune, len(lines[0]))
	for _, line := range lines {
		for i, ch := range line {
			resArr[i] = append(resArr[i], ch)
		}
	}
	for _, oneArr := range resArr {
		res = append(res, string(oneArr))
	}
	return res
}

func rotate45(lines []string) []string {
	var res []string
	length := len(lines)
	if length == 0 {
		return res
	}

	resArr := make([][]byte, 2*length-1)
	for j := range length {
		for i := range j + 1 {
			real_i := i
			real_j := length - j + i - 1
			resArr[j] = append(resArr[j], lines[real_i][real_j])
		}
	}
	for j := range length - 1 {
		for i := range length - j - 1 {
			real_i := j + 1 + i
			real_j := i
			resArr[length+j] = append(resArr[length+j], lines[real_i][real_j])
		}
	}
	for _, oneArr := range resArr {
		res = append(res, string(oneArr))
	}
	return res
}

func rotate45R(lines []string) []string {
	var res []string
	length := len(lines)
	if length == 0 {
		return res
	}

	resArr := make([][]byte, 2*length-1)
	for j := range length {
		for i := range j + 1 {
			real_i := j - i
			real_j := i
			resArr[j] = append(resArr[j], lines[real_i][real_j])
		}
	}
	for j := range length - 1 {
		for i := range length - j - 1 {
			real_i := length - i - 1
			real_j := j + i + 1
			resArr[length+j] = append(resArr[length+j], lines[real_i][real_j])
		}
	}
	for _, oneArr := range resArr {
		res = append(res, string(oneArr))
	}
	return res
}

func printLines(lines []string) {
	for _, line := range lines {
		fmt.Println(line)
	}
}
