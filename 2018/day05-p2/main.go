package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func resolve(txt string) int {
	str := txt
	oldStr := ""
	for len(str) != len(oldStr) {
		oldStr = str
		for i := 0; i < len(str)-1; i++ {
			if strings.ToLower(string(str[i])) == string(str[i]) {
				// it is lowercase
				if strings.ToUpper(string(str[i])) == string(str[i+1]) {
					str = str[:i] + str[i+2:]
					i--
				}
			} else {
				// it is uppercase
				if strings.ToLower(string(str[i])) == string(str[i+1]) {
					str = str[:i] + str[i+2:]
					i--
				}
			}
		}
	}
	return len(str)
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	chars := map[string]int{
		"a": 0,
		"b": 0,
		"c": 0,
		"d": 0,
		"e": 0,
		"f": 0,
		"g": 0,
		"h": 0,
		"i": 0,
		"j": 0,
		"k": 0,
		"l": 0,
		"m": 0,
		"n": 0,
		"o": 0,
		"p": 0,
		"q": 0,
		"r": 0,
		"s": 0,
		"t": 0,
		"u": 0,
		"v": 0,
		"w": 0,
		"x": 0,
		"y": 0,
		"z": 0,
	}
	min := 99999999999999999
	for s.Scan() {
		str := s.Text()
		for char := range chars {
			strToResolve := strings.Replace(strings.Replace(str, char, "", -1), strings.ToUpper(char), "", -1)
			chars[char] = resolve(strToResolve)
			if chars[char] < min {
				min = chars[char]
				fmt.Println(char, chars[char])
			}
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(min)
}
