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

func decodeDecomp(str string) (int, int) {
	s := strings.Split(str, "x")
	nchars, err := strconv.Atoi(s[0])
	check(err)
	rep, err := strconv.Atoi(s[1])
	check(err)
	return nchars, rep
}

var part2 = true

func decompress(s string) int {
	idx := strings.Index(s, "(")
	if idx == -1 {
		return len(s)
	}
	ret := 0

	for idx != -1 {
		ret += idx
		idx2 := strings.Index(s, ")")
		nchars, rep := decodeDecomp(s[idx+1 : idx2])

		s = s[idx2+1:]

		if part2 {
			ret += decompress(s[:nchars]) * rep
		} else {
			ret += nchars * rep
		}
		s = s[nchars:]

		idx = strings.Index(s, "(")
	}
	ret += len(s)
	return ret
}

func main() {
	dat, err := ioutil.ReadFile("input.txt")
	check(err)

	fmt.Println(decompress(strings.TrimSpace(string(dat))))
}
