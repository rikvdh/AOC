package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	sum := int64(0)
	seen := map[int64]bool{0: true}
	s := bufio.NewScanner(f)
	for {
		for s.Scan() {
			i, err := strconv.ParseInt(s.Text(), 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			sum += i
			if seen[sum] {
				log.Println(sum)
				return
			}
			seen[sum] = true
		}
		if err := s.Err(); err != nil {
			log.Fatal(err)
		}
		f.Seek(0, 0)
		s = bufio.NewScanner(f)
	}
}
