package main

import (
	"fmt"
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"sync"
	"sort"
)

const Multiplier = 10000
var wg sync.WaitGroup

type HashesFound map[int64]string

var hashes HashesFound

func calcHash(DoorId string, x int64) {
	defer wg.Done()

	var i int64 = 0

	for i < Multiplier {
		index := x * Multiplier + i
		md5sum := md5.Sum([]byte(DoorId + strconv.FormatInt(index, 10)))
		md5str := hex.EncodeToString(md5sum[0:16])
		if md5str[:5] == "00000" {
			hashes[index] = md5str
			fmt.Println("Found md5 on index", index, ":", md5str)
		}
		i++
	}
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

	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", hashes[int64(k)])
	}
}
