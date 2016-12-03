package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// 1 2 3
// 4 5 6
// 7 8 9
// The lookup-table based on the keypad above
var up = []int{0, 1, 2, 3, 1, 2, 3, 4, 5, 6}
var right = []int{0, 2, 3, 3, 5, 6, 6, 8, 9, 9}
var down = []int{0, 4, 5, 6, 7, 8, 9, 7, 8, 9}
var left = []int{0, 1, 1, 2, 4, 4, 5, 7, 7, 8}

var char int = 5

func main() {
	dat, err := ioutil.ReadFile("input.txt")
	check(err)

	strs := strings.Split(string(dat), "\n")

	fmt.Printf("The code is: ")
	for _, str := range strs {
		if len(str) > 0 {
			for _, ch := range str {
				switch ch {
				case 'U':
					char = up[char]
				case 'D':
					char = down[char]
				case 'L':
					char = left[char]
				case 'R':
					char = right[char]
				default:
					fmt.Printf("Unwanted character: %c\n", ch)
					panic("Argh")
				}
			}
			fmt.Print(char)
		}
	}
	fmt.Print("\n")
}
