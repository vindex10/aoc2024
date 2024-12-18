package main

import "os"
import "bufio"
import "fmt"
import "strings"
import "strconv"
import "sort"

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

	var list1 []int;
	var list2 []int;

	for scanner.Scan() {
		buf := scanner.Text()
		buf = strings.TrimSuffix(buf, "\n")
		parts := strings.SplitN(buf, "\t", 2)
		parsed, err := strconv.Atoi(parts[0])
		check(err)
		list1 = append(list1, parsed)
		parsed, err = strconv.Atoi(parts[1])
		check(err)
		list2 = append(list2, parsed)
	}

	f.Close()
	sort.Ints(list1)
	sort.Ints(list2)
	
	tot := 0;
	for i, v1 := range list1 {
		dist := list2[i] - v1
		if dist >= 0 {
			tot += dist
		} else {
			tot -= dist
		}
	}

	fmt.Println(tot)
}
