package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"
)

type stack struct {
	lock sync.Mutex
	s    []byte
}

func NewStack() *stack {
	return &stack{
		sync.Mutex{},
		make([]byte, 0),
	}
}

func (s *stack) Push(v byte) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.s = append(s.s, v)
}

func (s *stack) Size() int {
	return len(s.s)
}

func (s *stack) GetTop() byte {
	return s.s[s.Size()-1]
}

func (s *stack) Pop() (byte, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	length := len(s.s)
	if length <= 0 {
		return 0, errors.New("Empty Stack")
	}

	res := s.s[length-1]
	s.s = s.s[:length-1]
	return res, nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func changeCase(char byte) byte {
	return byte(char) ^ 32
}

func lowerCase(char byte) byte {
	return byte(char) | 32
}

func solve(s *stack, lines string, ignore byte) {
	for i := 0; i < len(lines); i++ {
		if lowerCase(lines[i]) == ignore {
			continue
		}
		if s.Size() == 0 {
			s.Push(lines[i])
			continue
		} else if changeCase(lines[i]) == s.s[len(s.s)-1] {
			s.Pop()
			continue
		}
		s.Push(lines[i])
	}
}

func smallest(lines string) int {
	var i byte
	smallestLine := len(lines)
	for i = 'a'; i <= 'z'; i++ {
		s := NewStack()
		solve(s, lines, i)
		if s.Size() < smallestLine {
			smallestLine = s.Size()
		}
	}
	return smallestLine
}

func main() {
	lines, err := readLines("day5.txt")
	input := lines[0]
	s := NewStack()
	if err != nil {
		panic(err)
	}
	start := time.Now()
	solve(s, input, 0)
	fmt.Println(s.Size())
	fmt.Println(smallest(input))
	t := time.Now()
	elapsed := t.Sub(start)

	fmt.Println(elapsed)
}
