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

	claimsValid := map[int]bool{}
	fabric := map[int]map[int]int{}
	for s.Scan() {
		var id, offx, offy, szx, szy int
		_, err := fmt.Sscanf(s.Text(), "#%d @ %d,%d: %dx%d", &id, &offx, &offy, &szx, &szy)
		if err != nil {
			log.Fatal(err)
		}
		if id == 0 {
			log.Fatal("id 0 unsupported")
		}
		valid := true
		for x := offx; x < offx+szx; x++ {
			for y := offy; y < offy+szy; y++ {
				if _, ok := fabric[x]; !ok {
					fabric[x] = make(map[int]int)
				}
				if fabric[x][y] != 0 {
					claimsValid[fabric[x][y]] = false
					valid = false
				}
				fabric[x][y] = id
			}
		}
		claimsValid[id] = valid
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	for id, valid := range claimsValid {
		if valid {
			fmt.Println(id)
		}
	}
}
