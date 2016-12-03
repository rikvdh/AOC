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

var validTriangles int = 0

func main() {
	dat, err := ioutil.ReadFile("input.txt")
	check(err)

	triangles := strings.Split(string(dat), "\n")

	for n, triangle := range triangles {
		numbers := strings.Fields(triangle)
		if len(numbers) == 3 {
			num1, err := strconv.Atoi(numbers[0])
			check(err)
			num2, err := strconv.Atoi(numbers[1])
			check(err)
			num3, err := strconv.Atoi(numbers[2])
			check(err)
			if num1+num2 > num3 && num2+num3 > num1 &&
				num1+num3 > num2 {
				fmt.Println("  Valid triangle", numbers)
				validTriangles++
			} else {
				fmt.Println("Invalid triangle", numbers)
			}
		} else if len(numbers) > 0 {
			fmt.Println("Invalid triangle", n, numbers)
			panic("Boem")
		}
	}
	fmt.Println("Valid triangles: ", validTriangles)
}
