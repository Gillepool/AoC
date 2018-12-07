package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const size = 1000

type clothStruct struct {
	id    int
	count int
}

func main() {
	inputs := read_file("day3.txt")
	start := time.Now()
	partA(inputs)
	fmt.Println(time.Since(start))
}

func partA(inputs []string) {

	cloth := make([][]clothStruct, size)
	overlapped := make(map[int]bool, 64)
	for i := range cloth {
		cloth[i] = make([]clothStruct, size)
	}
	for _, input := range inputs {
		var id, x, y, w, h int
		_, err := fmt.Sscanf(input, "#%d @ %d,%d: %dx%d\n", &id, &x, &y, &w, &h)
		check(err)
		for i := x; i < x+w; i++ {
			for j := y; j < y+h; j++ {
				if cloth[i][j].id > 0 {
					overlapped[cloth[i][j].id] = true
					overlapped[id] = true
				}
				cloth[i][j].count++
				cloth[i][j].id = id
			}
		}
	}
	overlap := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if cloth[i][j].count > 1 {
				overlap++
			}
			if _, did_overlap := overlapped[cloth[i][j].id]; cloth[i][j].count == 1 && !did_overlap {
				overlapped[cloth[i][j].id] = true
			}
		}
	}
	fmt.Println(overlap)
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
