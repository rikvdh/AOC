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

	fabric := map[int]map[int]int{}
	for s.Scan() {
		var id, offx, offy, szx, szy int
		_, err := fmt.Sscanf(s.Text(), "#%d @ %d,%d: %dx%d", &id, &offx, &offy, &szx, &szy)
		if err != nil {
			log.Fatal(err)
		}
		for x := offx; x < offx+szx; x++ {
			for y := offy; y < offy+szy; y++ {
				if _, ok := fabric[x]; !ok {
					fabric[x] = make(map[int]int)
				}
				fabric[x][y]++
			}
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	doubles := 0
	for _, xclaims := range fabric {
		for _, n := range xclaims {
			if n > 1 {
				doubles++
			}
		}
	}
	fmt.Println(doubles)
}
