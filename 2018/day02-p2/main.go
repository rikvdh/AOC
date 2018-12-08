package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/agnivade/levenshtein"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)

	strs := []string{}
	for s.Scan() {
		strs = append(strs, s.Text())
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	for _, a := range strs {
		for _, b := range strs {
			if levenshtein.ComputeDistance(a, b) == 1 {
				for k, c := range a {
					if c == rune(b[k]) {
						fmt.Print(string(c))
					}
				}
				fmt.Println("")
				return
			}
		}
	}
}
