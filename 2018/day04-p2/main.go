package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type dmy struct {
	year, month, day int
}

func (d dmy) String() string {
	return fmt.Sprintf("%02d-%02d-%04d", d.day, d.month, d.year)
}

func (d *dmy) addOne() {
	d.day++
	switch d.month {
	case 1, 3, 5, 7, 8, 10, 12:
		if d.day == 32 {
			d.month++
		}

	case 4, 6, 9, 11:
		if d.day == 31 {
			d.month++
			d.day = 1
		}
	case 2:
		if d.day >= 29 {
			if (d.year%4) == 0 && d.day >= 30 {
				d.month++
				d.day = 1
			} else {
				d.month++
				d.day = 1
			}
		}
	}
	if d.month == 13 {
		d.month = 1
	}
}

type dat struct {
	guard     int
	sleepMins map[int]bool
	wakeMins  map[int]bool
}

func (d dat) countSleepMins() int {
	mins := 0
	wake := true
	for i := -60; i < 60; i++ {
		if wake && d.sleepMins[i] {
			wake = false
		}
		if !wake && d.wakeMins[i] {
			wake = true
		}
		// only between 0:00 and 0:59
		if i >= 0 && !wake {
			mins++
		}
	}
	return mins
}

func (d dat) awakeAt(min int) bool {
	wake := true
	for i := -60; i < 60; i++ {
		if wake && d.sleepMins[i] {
			wake = false
		}
		if !wake && d.wakeMins[i] {
			wake = true
		}
		if i == min {
			return wake
		}
	}
	return wake
}

func (d dat) String() string {
	ret := fmt.Sprintf("% 5d ", d.guard)
	wake := true
	for i := -60; i < 60; i++ {
		if wake && d.sleepMins[i] {
			wake = false
		}
		if !wake && d.wakeMins[i] {
			wake = true
		}
		if i >= 0 {
			if wake {
				ret += "."
			} else {
				ret += "#"
			}
		}
	}
	return ret
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)

	data := map[dmy]*dat{}
	guards := []int{}

	for s.Scan() {
		var d, m, y, h, i, guard, ev int

		if strings.Contains(s.Text(), "Guard #") {
			ev = 1
			// [1518-06-09 00:04] Guard #2137 begins shift
			fmt.Sscanf(s.Text(), "[%d-%d-%d %d:%d] Guard #%d", &y, &m, &d, &h, &i, &guard)
			guards = append(guards, guard)
		} else if strings.Contains(s.Text(), "wakes up") {
			ev = 2
			// [1518-02-20 00:52] wakes up
			fmt.Sscanf(s.Text(), "[%d-%d-%d %d:%d]", &y, &m, &d, &h, &i)

		} else if strings.Contains(s.Text(), "falls asleep") {
			ev = 3
			// [1518-06-21 00:36] falls asleep
			fmt.Sscanf(s.Text(), "[%d-%d-%d %d:%d]", &y, &m, &d, &h, &i)
		} else {
			log.Fatal("unknown:" + s.Text())
		}
		day := dmy{y, m, d}
		if h == 23 {
			// 23:58 --> :-2
			i -= 60
			day.addOne()
		}

		if data[day] == nil {
			data[day] = &dat{sleepMins: make(map[int]bool), wakeMins: make(map[int]bool)}
		}
		if ev == 1 {
			data[day].guard = guard
		} else if ev == 2 {
			data[day].wakeMins[i] = true
		} else if ev == 3 {
			data[day].sleepMins[i] = true
		} else {
			log.Fatal("invalid ev")
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	max := 0
	guard := -1
	minute := -1

	for _, g := range guards {
		for m := 0; m < 60; m++ {
			current := 0
			for _, d := range data {
				if d.guard == g {
					if !d.awakeAt(m) {
						current++
					}
					if m == 0 {
						//fmt.Println(day, d)
					}
				}
			}
			if current > max {
				max = current
				minute = m
				guard = g
			}
		}
	}
	fmt.Println(minute * guard)
}
