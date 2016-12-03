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

//     1
//   2 3 4
// 5 6 7 8 9
//   A B C    = 10, 11, 12
//     D      = 13
// The lookup-table for part 2
//                 x  1  2  3  4  5  6  7  8  9  A  B  C  D
var up2    = []int{0, 1, 2, 1, 4, 5, 2, 3, 4, 9, 6, 7, 8,11}
var right2 = []int{0, 1, 3, 4, 4, 6, 7, 8, 9, 9,11,12,12,13}
var down2  = []int{0, 3, 6, 7, 8, 5,10,11,12, 9,10,13,12,13}
var left2  = []int{0, 1, 2, 2, 3, 5, 5, 6, 7, 8,10,10,11,13}

var char int = 5
var char2 int = 5
var code2 []byte

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
					char2 = up2[char2]
				case 'D':
					char = down[char]
					char2 = down2[char2]
				case 'L':
					char = left[char]
					char2 = left2[char2]
				case 'R':
					char = right[char]
					char2 = right2[char2]
				default:
					fmt.Printf("Unwanted character: %c\n", ch)
					panic("Argh")
				}
			}
			fmt.Print(char)
			switch char2 {
			case 10:
				code2 = append(code2, 'A')
			case 11:
				code2 = append(code2, 'B')
			case 12:
				code2 = append(code2, 'C')
			case 13:
				code2 = append(code2, 'D')
			default:
				code2 = append(code2, byte(0x30+char2))
			}
		}
	}
	fmt.Print("\n")
	fmt.Println("Code 2 is ", string(code2))
}
