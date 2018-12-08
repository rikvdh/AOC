package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)

	twos := 0
	threes := 0
	for s.Scan() {
		seen := map[rune]int{}
		for _, ch := range s.Text() {
			seen[ch]++
		}
		var two, three bool
		for _, n := range seen {
			if n == 2 && !two {
				twos++
				two = true
			}
			if n == 3 && !three {
				threes++
				three = true
			}
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(twos * threes)
}
