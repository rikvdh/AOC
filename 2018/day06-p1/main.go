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

	inf := make(map[Input]bool)
	m := make(map[Input]int)

	for _, input := range inputList {
		if input.Y > height {
			height = input.Y
		}

		if input.X > width {
			width = input.X
		}
	}

	for y := float64(-200); y < height+200; y++ {
		for x := float64(-200); x < width+200; x++ {
			mc := Input{0, 0, 0}
			min := float64(-1)

			for _, c := range inputList {
				dist := math.Abs(x-c.X) + math.Abs(y-c.Y)
				if dist < min || min == -1 {
					min = dist
					mc = c
				} else if dist == min {
					mc = Input{-1, -1, -1}
				}
			}

			if x == 0 || y == 0 || x == width || y == height {
				inf[mc] = true
			}

			m[mc]++
		}
	}

	max := 0
	for k, v := range m {
		if _, found := inf[k]; v > max && !found {
			max = v
		}
	}

	fmt.Println(max)
}
