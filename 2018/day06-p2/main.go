package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type Input struct {
	N int
	X float64
	Y float64
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)

	inputList := []Input{}

	n := 0

	for s.Scan() {
		line := s.Text()
		if len(line) > 0 {
			i := Input{}

			if _, err := fmt.Sscanf(line, "%b, %b", &i.X, &i.Y); err != nil {
				log.Fatal(err)
			}
			i.N = n
			n++

			inputList = append(inputList, i)
		}
	}

	height := float64(0)
	width := float64(0)

	for _, input := range inputList {
		if input.Y > height {
			height = input.Y
		}

		if input.X > width {
			width = input.X
		}
	}
	area := 0
	for y := float64(-200); y < height+200; y++ {
		for x := float64(-200); x < width+200; x++ {
			min := float64(-1)
			sumDist := float64(0)

			for _, c := range inputList {
				dist := math.Abs(x-c.X) + math.Abs(y-c.Y)
				if dist < min || min == -1 {
					min = dist
				}
				sumDist += dist
			}
			if sumDist < 10000 {
				area++
			}
		}
	}

	fmt.Println(area)
}
