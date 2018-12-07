package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

func main() {
	input := read_file("day2.txt")
	resultA := part_a(input)
	start := time.Now()
	result := part_b(input)
	fmt.Println(time.Since(start))
	fmt.Println(result)
	fmt.Println(resultA)

}

func part_a(inputs []string) int {
	twos := 0
	threes := 0
	for i := range inputs {
		found_two := false
		found_three := false
		for _, c := range inputs[i] {
			counts := strings.Count(inputs[i], string(c))

			if counts == 2 {
				found_two = true
			}
			if counts == 3 {
				found_three = true
			}
		}
		if found_two {
			twos += 1
		}
		if found_three {
			threes += 1
		}
	}

	return twos * threes
}

func mismatch_count(s1 string, s2 string) (int, int) {
	c := 0
	index := 0
	for i := range s1 {
		if s1[i] != s2[i] {
			c++
			index = i
		}
	}
	return c, index
}

func part_b(inputs []string) string {
	sort.Strings(inputs)
	for id, _ := range inputs[:len(inputs)-1] {
		c, index := mismatch_count(inputs[id], inputs[id+1])
		if c == 1 {
			result := string(inputs[id][0:index]) + string(inputs[id][index+1:len(inputs[id])])
			return result
		}
	}
	return ""
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func read_file(filename string) (data []string) {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	return data
}
