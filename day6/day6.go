package main

import (
	"fmt"
	_ "sort"
	"io/ioutil"
	_ "strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type StringArray []string

func (s StringArray) Len() int {
    return len(s)
}
func (s StringArray) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s StringArray) Less(i, j int) bool {
    return s[i][0] < s[j][0]
}

type Columns struct {
	letters map[byte]int
}

func getMaxLetter(col Columns) string {
	var maxLetter byte = '-'
	var maxFreq int = 0
	for letter, freq := range col.letters {
		if freq > maxFreq {
			maxLetter = letter
			maxFreq = freq
		}
	}

	return string(maxLetter)
}

func getMinLetter(col Columns) string {
	var minLetter byte = '-'
	var minFreq int = 999
	for letter, freq := range col.letters {
		if freq < minFreq {
			minLetter = letter
			minFreq = freq
		}
	}

	return string(minLetter)
}

func main() {
	dat, err := ioutil.ReadFile("input.txt")
	check(err)

	lines := strings.Split(string(dat), "\n")

	var cols [20]Columns

	var i int = 0
	for i < len(lines[0]) {
		cols[i].letters = make(map[byte]int)
		i++;
	}

	for _, line := range lines {
		if len(line) > 0 {
			var i int = 0
			for i < len(line) {
				if _, ok := cols[i].letters[line[i]]; ok {
					cols[i].letters[line[i]]++
				} else {
					cols[i].letters[line[i]] = 1
				}
				i++;
			}
		}
	}

	i = 0
	var msg string
	var minmsg string
	for i < len(lines[0]) {
		msg += getMaxLetter(cols[i])
		minmsg += getMinLetter(cols[i])
		i++;
	}
	fmt.Println(msg)
	fmt.Println(minmsg)
}
