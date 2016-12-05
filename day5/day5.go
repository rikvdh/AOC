package main

import (
	"fmt"
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"sync"
	"sort"
)

const Multiplier = 50000
var wg sync.WaitGroup

type HashesFound map[int64]string

var hashes HashesFound

func calcHash(DoorId string, x int64) {
	defer wg.Done()

	var i int64 = 0

	for i < 1000 {
		index := x * 1000 + i
		md5sum := md5.Sum([]byte(DoorId + strconv.FormatInt(index, 10)))
		md5str := hex.EncodeToString(md5sum[0:16])
		if md5str[:5] == "00000" {
			hashes[index] = md5str
		}
		i++
	}
}

func replaceAtIndex(str string, replacement byte, index int) string {
	return str[:index] + string(replacement) + str[index+1:]
}

func main() {
	var DoorId = "uqwqemis"
	var i int64 = 0

	hashes = make(HashesFound)

	for i < Multiplier {
		wg.Add(1)
		go calcHash(DoorId, i)
		i++
	}
	wg.Wait()

	var keys []int
	for k := range hashes {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)

	var code string
	var code2 string = "----------"

	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", hashes[int64(k)])
		code += hashes[int64(k)][5:6]
		pos,err := strconv.Atoi(hashes[int64(k)][5:6])
		if err == nil {
			val := hashes[int64(k)][6:7]
			if code2[pos] == '-' {
				code2 = replaceAtIndex(code2, val[0], pos)
			}
		}
	}
	fmt.Println("Code is:", code[:8])
	fmt.Println("2nd code is:",code2[:8])
}
