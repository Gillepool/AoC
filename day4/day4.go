package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"time"
)

// Entry for storing information about the time and activity
type Entry struct {
	Time     time.Time
	activity string
}

type Entries []*Entry

// Tired Guard af
type Guard struct {
	ID                int
	MinuteSleeps      []int
	TotalTimeSleeping int
}

func (entry *Entry) GuardID() (id int) {
	_, err := fmt.Sscanf(entry.activity, "Guard #%d begins shift", &id)
	if err != nil {
		return -1
	}
	return id
}

func CreateEntries(lines []string) Entries {
	entries := make(Entries, 0, len(lines))
	for _, line := range lines {
		entries = append(entries, LoadEntry(line))
	}
	sort.Slice(entries, func(i, j int) bool {
		time1 := entries[i].Time
		time2 := entries[j].Time
		return time1.Before(time2)
	})
	return entries
}

func LoadEntry(line string) *Entry {
	var reg = regexp.MustCompile("^\\[([0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2})] (.+)$")
	matches := reg.FindStringSubmatch(line)
	if matches == nil {
		panic("Did not find matches for line")
	}
	time, err := time.Parse("2006-01-02 15:04", matches[1])
	if err != nil {
		panic(err)
	}
	activity := matches[2]

	return &Entry{
		Time:     time,
		activity: activity,
	}
}

// NewGuard creates a new Guard instance
func NewGuard(id int) *Guard {
	return &Guard{
		ID:                id,
		MinuteSleeps:      make([]int, 60),
		TotalTimeSleeping: 0,
	}
}

func (guard *Guard) SleepingMinutesGuard(startSleeping, endSleeping time.Time) {
	for i := startSleeping.Minute(); i < endSleeping.Minute(); i++ {
		guard.MinuteSleeps[i]++
	}
	guard.TotalTimeSleeping += endSleeping.Minute() - startSleeping.Minute()
}

func (g *Guard) MostSleptMinute() (minute, max int) {
	for m, count := range g.MinuteSleeps {
		if count > max {
			minute = m
			max = count
		}
	}
	return
}

func (guardMap GuardMap) SleepiestGuard() (g *Guard) {
	for _, v := range guardMap {
		if g == nil {
			g = v
		}
		if v.TotalTimeSleeping > g.TotalTimeSleeping {
			g = v
		}
	}
	return g
}

func (m GuardMap) SleepingPart2() (g *Guard) {
	var count int
	for _, v := range m {
		_, c := v.MostSleptMinute()
		if c > count {
			count = c
			g = v
		}
	}
	return
}

// GuardMap contains all Guard instances
type GuardMap map[int]*Guard

func CreatingMapOfGuards(entries Entries) GuardMap {
	var guard *Guard
	var startSleep time.Time
	var endSleep time.Time
	guardMap := make(GuardMap)
	for _, entry := range entries {
		switch entry.activity {
		case "falls asleep":
			startSleep = entry.Time
		case "wakes up":
			endSleep = entry.Time
			guard.SleepingMinutesGuard(startSleep, endSleep)
		default:
			guard = guardMap.getGuard(entry.GuardID())

		}
	}
	return guardMap
}

func (guardMap GuardMap) getGuard(id int) *Guard {
	if _, ok := guardMap[id]; !ok {
		guardMap[id] = NewGuard(id)

	}
	return guardMap[id]
}

func main() {
	lines := get_input("day4.txt")
	entris := CreateEntries(lines)
	guardMap := CreatingMapOfGuards(entris)
	guard := guardMap.SleepiestGuard()
	mostMin, _ := guard.MostSleptMinute()
	fmt.Println(guard.ID * mostMin)

	guard2 := guardMap.SleepingPart2()
	mostMin2, _ := guard2.MostSleptMinute()

	fmt.Println(guard2.ID * mostMin2)
}

func get_input(file_name string) []string {
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}
