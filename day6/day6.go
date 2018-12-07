package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type coordinate struct {
	x, y int
}

type Grid struct {
	minX, maxX, minY, maxY int
	Points                 map[int]coordinate
}

func main() {
	grid, err := readLines("day6.txt")
	check(err)
	grid.solve()
	grid.solve2(10000)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func readLines(path string) (*Grid, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	points := make(map[int]coordinate)
	var i = 0
	maxX, maxY := 0, 0
	minX, minY := 9999999, 9999999
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ", ")
		x, _ := strconv.Atoi(line[0])
		y, _ := strconv.Atoi(line[1])
		points[i] = coordinate{x: x, y: y}
		i++
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
		if y < minY {
			minY = y
		}
		if x < minX {
			minX = x
		}
	}
	fmt.Println(len(points))
	return &Grid{minX: minX, minY: minY, maxX: maxX, maxY: maxY, Points: points}, nil
}

// ManhattanDistance returns the manhattan distance
func ManhattanDistance(x1, y1, x2, y2 int) int {
	return int(math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2)))
}

func (grid *Grid) printGrid() {
	for x := grid.minX; x < grid.maxX; x++ {
		for y := grid.minY; y < grid.maxY; y++ {
			for _, p := range grid.Points {
				if p.x == x && p.y == y {
					fmt.Print("P")
				}
			}
			fmt.Print("*")
		}
		fmt.Println("*")
	}
}

func (grid *Grid) solve() {
	sizes := make(map[int]int)
	for x := grid.minX; x < grid.maxX; x++ {
		for y := grid.minY; y < grid.maxY; y++ {
			minDistance := math.MaxInt64
			equallyClose := false
			placeId := 0
			for id, p := range grid.Points {
				if dist := ManhattanDistance(x, y, p.x, p.y); dist <= minDistance {
					equallyClose = minDistance == dist
					placeId = id
					minDistance = dist
				}
			}
			if equallyClose == false {
				sizes[placeId]++
			}
		}
	}

	size := 0

	for id, p := range grid.Points {
		if p.x == grid.maxX || p.x == grid.minX || p.y == grid.minY || p.y == grid.maxY {
			continue
		}

		if sizes[id] > size {
			size = sizes[id]
		}
	}

	fmt.Println("Biggest area is", size)
}

func (grid *Grid) solve2(maxSize int) {
	size := 0
	for x := grid.minX; x < grid.maxX; x++ {
		for y := grid.minY; y < grid.maxY; y++ {
			total := 0

			for _, p := range grid.Points {
				total += ManhattanDistance(x, y, p.x, p.y)
			}

			if total < int(maxSize) {
				size++
			}
		}
	}

	fmt.Println(size)

}
