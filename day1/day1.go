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

var vertical int = 0
var horizontal int = 0
var direction int = 0

func main() {
	dat, err := ioutil.ReadFile("input.txt")
	check(err)

	strs := strings.Split(string(dat), ",")
	for k, str := range strs {
		strs[k] = strings.Trim(str, "\n ")
	}

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

		switch direction {
		case 0:
			vertical += dist
		case 180:
			vertical -= dist
		case 90:
			horizontal += dist
		case 270:
			horizontal -= dist
		}
	}
	fmt.Printf("Horizontal: %d\nVertical: %d\n", horizontal, vertical)
	fmt.Printf("Easter Bunny HQ is %.0f blocks away\n",
		math.Abs(float64(horizontal))+math.Abs(float64(vertical)))
}
