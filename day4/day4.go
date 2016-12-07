package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
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

type Room struct {
	letters  map[string]int
	sector   string
	checksum string
	realname string
}

func getChecksum(letters map[string]int) string {
	var maxfreq int
	var checksumResult string
	var sortBuffer StringArray

	for _, freq := range letters {
		if freq > maxfreq {
			maxfreq = freq
		}
	}

	for maxfreq > 0 {
		sortBuffer = sortBuffer[:0]
		for letter, freq := range letters {
			if freq == maxfreq {
				sortBuffer = append(sortBuffer, letter)
			}
		}
		sort.Sort(sortBuffer)
		for _, letter := range sortBuffer {
			checksumResult += letter
		}
		if len(checksumResult) >= 5 {
			break
		}
		maxfreq--
	}
	return checksumResult[:5]
}

func getRoom(roomstr string) Room {
	var r Room
	r.letters = make(map[string]int)
	parts := strings.Split(roomstr, "-")
	for _, part := range parts {
		for len(part) > 0 {
			if part[:1] == "[" {
				r.checksum = part[1:6]
				break
			} else if part[:1] >= "0" && part[:1] <= "9" {
				r.sector += part[:1]
			} else {
				if _, ok := r.letters[part[:1]]; ok {
					r.letters[part[:1]]++
				} else {
					r.letters[part[:1]] = 1
				}
			}
			part = part[1:]
		}
	}
	sectn, err := strconv.Atoi(r.sector)
	check(err)

	for _, letter := range roomstr {
		if letter == '-' {
			r.realname += " "
		} else if letter >= '0' && letter <= '9' {
			break
		} else {
			b := byte((int(letter)-'a'+sectn)%26 + 'a')
			r.realname = string(append([]byte(r.realname), b))
		}
	}

	return r
}

var total int64
var roomCount int64
var validRoomCount int64

func main() {
	dat, err := ioutil.ReadFile("input.txt")
	check(err)

	rooms := strings.Split(string(dat), "\n")

	for _, roomstr := range rooms {
		if len(roomstr) > 0 {
			r := getRoom(roomstr)
			roomCount++
			if getChecksum(r.letters) == r.checksum {
				validRoomCount++
				n, _ := strconv.Atoi(r.sector)
				total += int64(n)
				if r.realname == "northpole object storage " {
					fmt.Println("Valid room", r.realname, r.sector)
				}
			}
		}
	}
	fmt.Println("Sectorsum:", total, "in", validRoomCount, "of", roomCount, "rooms")
}
