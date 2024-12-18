package main

import "os"
import "bufio"
import "fmt"
import "strings"
import "strconv"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func validateLine(line []int) int {
	failedStatus := 0
	if len(line) <= 1 {
		return (1-failedStatus)  // successStatus
	}
	prev := line[0]
	directionUpwards := (line[1] - line[0]) > 0
	for _, v := range line[1:] {
		if v == prev {
			failedStatus = 1
			break
		}
		diff := v - prev
		thisDirectionUpwards := diff > 0
		if thisDirectionUpwards != directionUpwards  {
			failedStatus = 1
			break
		}
		var dist int
		if thisDirectionUpwards {
			dist = v - prev
		} else {
			dist = prev - v
		}
		if ! ((dist >= 1) && (dist <= 3)) {
			failedStatus = 1
			break
		}
		prev = v
	}
	return (1 - failedStatus)
}

func parseLine(buf string) ([]int, error) {
		line := strings.TrimSuffix(buf, "\n")
		parts := strings.Split(line, "\t")
		var parts_int []int
		for _, one := range parts {
			parsed, err := strconv.Atoi(one)
			if err != nil {
				return nil, err
			}
			parts_int = append(parts_int, parsed)
		}
		return parts_int, nil
}

func main() {
	fpath := os.Args[1]
	f, err := os.Open(fpath)
	check(err)

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var lines [][]int;

	for scanner.Scan() {
		buf := scanner.Text()
		parsed, err := parseLine(buf)
		check(err)
		lines = append(lines, parsed)
	}

	f.Close()
	
	tot := 0;
	for _, v := range lines {
		tot += validateLine(v)
	}

	fmt.Println(tot)
}
