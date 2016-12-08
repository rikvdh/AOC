package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const ScreenWidth = 50
const ScreenHeight = 6

var screen [ScreenWidth][ScreenHeight]int

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func doRectangle(rect string) {
	wh := strings.Split(rect, "x")
	w, _ := strconv.Atoi(wh[0])
	h, _ := strconv.Atoi(wh[1])
	x := 0
	y := 0
	for x < w {
		y = 0
		for y < h {
			screen[x][y] = 1
			y++
		}
		x++
	}
}

func rotateCol(column int, n int) {
	//fmt.Println("Move x=", column, "by", n)
	prev := 0
	i := 0
	for i < n {
		y := 0
		for y < ScreenHeight {
			oldPrev := prev
			prev = screen[column][y]
			screen[column][y] = oldPrev
			y++
		}
		// Overflow the end to the beginning
		screen[column][0] = prev
		i++
	}
}
func rotateRow(row int, n int) {
	//fmt.Println("Move y=", row, "by", n)
	prev := 0
	i := 0
	for i < n {
		x := 0
		for x < ScreenWidth {
			oldPrev := prev
			prev = screen[x][row]
			screen[x][row] = oldPrev
			x++
		}
		// Overflow the end to the beginning
		screen[0][row] = prev
		i++
	}
}

func doRotate(rowcol string, dat []string) {
	n, err := strconv.Atoi(dat[0][2:])
	check(err)
	amount, err := strconv.Atoi(dat[2])
	check(err)

	if rowcol == "column" {
		rotateCol(n, amount)
	} else if rowcol == "row" {
		rotateRow(n, amount)
	}
}

func renderRect() {
	y := 0
	ones := 0
	for y < ScreenHeight {
		x := 0
		for x < ScreenWidth {
			if screen[x][y] == 1 {
				fmt.Printf("X")
				ones++
			} else {
				fmt.Printf(" ")
			}
			x++
		}
		y++
		fmt.Printf("\n")
	}
	fmt.Println("Got", ones, "ones")
}

func main() {
	dat, err := ioutil.ReadFile("input.txt")
	check(err)

	lines := strings.Split(string(dat), "\n")

	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) > 0 {
			switch fields[0] {
			case "rect":
				doRectangle(fields[1])
			case "rotate":
				doRotate(fields[1], fields[2:])
			default:
				fmt.Println("Something else:", fields)
			}
		}
	}
	renderRect()
}
