package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isValid(x string, y string, z string) bool {
	num1, err := strconv.Atoi(x)
	check(err)
	num2, err := strconv.Atoi(y)
	check(err)
	num3, err := strconv.Atoi(z)
	check(err)
	if num1+num2 > num3 && num2+num3 > num1 &&
		num1+num3 > num2 {
		return true
	}
	return false
}

var validTriangles int = 0

func main() {
	dat, err := ioutil.ReadFile("input.txt")
	check(err)

	triangles := strings.Split(string(dat), "\n")

	var a, b, c string
	var calc int = 0

	for n, triangle := range triangles {
		switch calc {
		case 2:
			c = triangle
			calc = 0
		case 1:
			b = triangle
			calc = 2
		case 0:
			r1 := strings.Fields(a)
			r2 := strings.Fields(b)
			r3 := strings.Fields(c)
			a = triangle
			calc = 1

			if len(r1) == 3 {
				if len(r2) == 3 && len(r3) == 3 {
					for _, i := range []int{0, 1, 2} {
						if isValid(r1[i], r2[i], r3[i]) {
							validTriangles++
						}
					}
				}
			} else if len(r1) > 0 {
				fmt.Println("Invalid triangle", n, r1)
				panic("Boem")
			}
		}
	}
	fmt.Println("Valid triangles: ", validTriangles)
}
