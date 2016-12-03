package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Visited struct {
	coord string
	x     int
	y     int
}

var vertical int = 0
var horizontal int = 0
var direction int = 0
var visited []Visited

func main() {
	dat, err := ioutil.ReadFile("input.txt")
	check(err)

	strs := strings.Split(string(dat), ",")
	for k, str := range strs {
		strs[k] = strings.Trim(str, "\n ")
	}

	visited = append(visited, Visited{"0x0", 0, 0})

	fmt.Println(strs)
	for _, str := range strs {
		switch str[:1] {
		case "R":
			direction += 90
		case "L":
			direction -= 90
		default:
			fmt.Println(str)
			panic("Invalid instruction")
		}
		if direction < 0 {
			direction = 360 + direction
		}
		if direction == 360 {
			direction = 360 - direction
		}
		dist, err := strconv.Atoi(str[1:])
		check(err)

		for {
			switch direction {
			case 0:
				vertical += 1
			case 180:
				vertical -= 1
			case 90:
				horizontal += 1
			case 270:
				horizontal -= 1
			}

			visited = append(visited, Visited{
				strconv.Itoa(horizontal) + "x" + strconv.Itoa(vertical),
				horizontal,
				vertical})

			dist--
			if dist == 0 {
				break
			}
		}
	}
	fmt.Printf("Horizontal: %d\nVertical: %d\n", horizontal, vertical)
	fmt.Printf("Easter Bunny HQ is %.0f blocks away\n",
		math.Abs(float64(horizontal))+math.Abs(float64(vertical)))

	var visCount map[string]int = make(map[string]int)

	for _, coord := range visited {
		if _, ok := visCount[coord.coord]; !ok {
			visCount[coord.coord] = 1
		} else if visCount[coord.coord] == 1 {
			fmt.Println("Location ", coord.coord, " is first visited twice")
			fmt.Printf("Easter Bunny HQ is really %.0f blocks away\n",
				math.Abs(float64(coord.x))+math.Abs(float64(coord.y)))
			break
		}
	}
}
