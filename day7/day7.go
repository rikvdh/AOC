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

func ContainsAba(in string) (bool, []string) {
	end := len(in) - 2
	i := 0
	var strs []string

	for i < end {
		if in[i] == in[i+2] && in[i] != in[i+1] {
			strs = append(strs, in[i:i+3])
		}
		i++
	}
	if len(strs) > 0 {
		return true, strs
	} else {
		return false, nil
	}
}

func ContainsBab(in string, abas []string) bool {
	end := len(in) - 2
	i := 0

	for _, aba := range abas {
		i = 0
		for i < end {
			if in[i] == in[i+2] && in[i] != in[i+1] &&
				in[i] == aba[1] && in[i+1] == aba[0] {
				return true
			}
			i++
		}
	}

	return false
}

func ContainsAbba(in string) bool {
	end := len(in) - 3
	i := 0

	for i < end {
		if in[i] == in[i+3] && in[i+1] == in[i+2] && in[i] != in[i+1] {
			return true
		}
		i++
	}
	return false
}

func main() {
	validIps := 0
	validIps2 := 0
	dat, err := ioutil.ReadFile("input.txt")
	check(err)

	lines := strings.Split(string(dat), "\n")

	for _, line := range lines {
		insidePart := false
		insideStr := ""
		outsideStr := ""
		fields := strings.FieldsFunc(line, func(c rune) bool {
			return c == '[' || c == ']'
		})
		for _, part := range fields {
			if insidePart {
				insideStr += "   " + part
				insidePart = false
			} else {
				outsideStr += "   " + part
				insidePart = true
			}
		}

		if !ContainsAbba(insideStr) && ContainsAbba(outsideStr) {
			validIps++
		}

		_, abas := ContainsAba(outsideStr)
		if ContainsBab(insideStr, abas) {
			validIps2++
			fmt.Println(line)
		}
	}

	fmt.Println("Valid TLS IPs", validIps)
	fmt.Println("Valid SSL IPs", validIps2)
}
