package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		str := s.Text()
		oldStr := ""
		for len(str) != len(oldStr) {
			oldStr = str
			fmt.Println(len(str))
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
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}
