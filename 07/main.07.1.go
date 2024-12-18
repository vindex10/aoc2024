package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/stat/combin"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fpath := os.Args[1]
	tasks := parseInput(fpath)
	tot := 0
	for task, parts := range tasks {
		tot += solve(task, parts)
	}
	fmt.Println(tot)
}

func parseInput(fpath string) map[int]([]int) {
	f, err := os.Open(fpath)
	defer f.Close()
	check(err)

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	res := make(map[int]([]int))
	for scanner.Scan() {
		buf := scanner.Text()
		task, parts := parseLine(buf)
		res[task] = parts
	}
	return res
}

func parseLine(buf string) (int, []int) {
	buf = strings.TrimSpace(buf)
	parts := strings.SplitN(buf, ": ", 2)
	task, err := strconv.ParseInt(parts[0], 10, 64)
	check(err)
	var taskParts []int
	taskPartsStr := strings.Split(parts[1], " ")
	for _, v := range taskPartsStr {
		parsed, err := strconv.ParseInt(v, 10, 64)
		check(err)
		taskParts = append(taskParts, int(parsed))
	}
	return int(task), taskParts
}

func solve(task int, parts []int) int {
	signs := len(parts) - 1
	for numProds := range signs + 1 {
		for i := range combin.Combinations(signs, numProds) {
			comb := combin.IndexToCombination(nil, i, signs, numProds)
			buf := parts[0]
			signsI := 0
			for j := range signs {
				if signsI < len(comb) && j == comb[signsI] {
					buf *= parts[j+1]
					signsI += 1
				} else {
					buf += parts[j+1]
				}
			}
			if buf == task {
				fmt.Println("***", comb, parts, buf)
				return task
			}
			//fmt.Println(task, comb, parts, buf)
		}
	}
	return 0
}
