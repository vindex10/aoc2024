//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
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
	var x int
	var y int
	for scanner.Scan() {
		buf := scanner.Text()
		buf = strings.TrimSpace(buf)
		idx := strings.IndexByte(buf, '^')
		if idx != -1 {
			x = idx
			y = len(lines)
			buf = buf[:idx] + "." + buf[idx+1:]
		}
		lines = append(lines, buf)
	}
	f.Close()

	direction := '^'
	var tot int
	for {
		x, y, direction = doStep(lines, x, y, direction)
		//fmt.Scanln()
		tot = countStars(lines)
		//fmt.Println(tot)
		if direction == '.' {
			break
		}
	}

	show(lines, x, y, direction)
	fmt.Println(tot)
}

func doStep(lines []string, x int, y int, direction rune) (int, int, rune) {
	width := len(lines[0])
	height := len(lines)
	if direction == '^' {
		if y == 0 {
			return -1, -1, '.'
		}
		if lines[y-1][x] == '#' {
			return x, y, '>'
		}
		starStep(lines, x, y)
		return x, y - 1, '^'
	}
	if direction == '>' {
		if x == width-1 {
			return -1, -1, '.'
		}
		if lines[y][x+1] == '#' {
			return x, y, 'v'
		}
		starStep(lines, x, y)
		return x + 1, y, '>'
	}
	if direction == 'v' {
		if y == height-1 {
			return -1, -1, '.'
		}
		if lines[y+1][x] == '#' {
			return x, y, '<'
		}
		starStep(lines, x, y)
		return x, y + 1, 'v'
	}
	if direction == '<' {
		if x == 0 {
			return -1, -1, '.'
		}
		if lines[y][x-1] == '#' {
			return x, y, '^'
		}
		starStep(lines, x, y)
		return x - 1, y, '<'
	}
	fmt.Println(x, y, direction)
	panic("error")
}

func show(lines []string, x int, y int, direction rune) {
	for i, line := range lines {
		if i == y {
			line = line[:x] + string(direction) + line[x+1:]
		}
		fmt.Println(line)
	}
	fmt.Println()
}

func starStep(lines []string, x int, y int) {
	if lines[y][x] != '*' {
		lines[y] = lines[y][:x] + "*" + lines[y][x+1:]
	}
}

func countStars(lines []string) int {
	res := 0
	for _, line := range lines {
		res += strings.Count(line, "*")
	}
	return res
}
