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

type Bot struct {
	val1      int
	val2      int
	highIsBin bool
	high      int
	lowIsBin  bool
	low       int
}

type Output struct {
	values []int
}

var bots map[int]Bot
var bins map[int]Output

func doParseValue(fields []string) {
	value, _ := strconv.Atoi(fields[0])
	bot, _ := strconv.Atoi(fields[4])
	b, ok := bots[bot]
	if ok {
		if b.val1 != -1 && b.val2 == -1 {
			b.val2 = value
		} else if b.val1 == -1 {
			b.val1 = value
		} else {
			panic("Bot full")
		}
		bots[bot] = b
	} else {
		var b Bot
		b.val1 = value
		b.val2 = -1
		b.high = -1
		b.low = -1
		bots[bot] = b
	}
}

func doParseHiLo(fields []string) {
	bot, _ := strconv.Atoi(fields[0])
	hl := fields[2]
	lowBin := fields[4] == "output"
	low, _ := strconv.Atoi(fields[5])
	highBin := fields[9] == "output"
	high, _ := strconv.Atoi(fields[10])

	if hl != "low" {
		panic("Not low")
	}

	b, ok := bots[bot]
	if ok {
		b.lowIsBin = lowBin
		b.low = low
		b.highIsBin = highBin
		b.high = high
		bots[bot] = b
	} else {
		var b Bot
		b.val1 = -1
		b.val2 = -1
		b.lowIsBin = lowBin
		b.low = low
		b.highIsBin = highBin
		b.high = high
		bots[bot] = b
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func valueToBin(bin, value int) {
	b := bins[bin]
	b.values = append(b.values, value)
	bins[bin] = b
}

func giveVal(bot, value int) {
	b := bots[bot]

	if b.val1 != -1 && b.val2 == -1 {
		b.val2 = value
	} else if b.val1 == -1 {
		b.val1 = value
	} else {
		fmt.Println(bot, value)
		panic("Bot full")
	}
	bots[bot] = b
}

func runBotCycle() bool {
	proc := false
	for bot, b := range bots {
		if b.val1 != -1 && b.val2 != -1 {
			if min(b.val1, b.val2) == 17 &&
				max(b.val1, b.val2) == 61 {
				fmt.Println("Bot found processing 17 & 61, bot", bot)
			}
			if !b.highIsBin {
				giveVal(b.high, max(b.val1, b.val2))
			} else {
				valueToBin(b.high, max(b.val1, b.val2))
			}
			if !b.lowIsBin {
				giveVal(b.low, min(b.val1, b.val2))
			} else {
				valueToBin(b.low, min(b.val1, b.val2))
			}
			b.val1 = -1
			b.val2 = -1
			bots[bot] = b
			proc = true
		}
	}
	return proc
}

func main() {
	dat, err := ioutil.ReadFile("input.txt")
	check(err)

	lines := strings.Split(string(dat), "\n")
	bots = make(map[int]Bot)
	bins = make(map[int]Output)

	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) > 0 {
			switch fields[0] {
			case "value":
				doParseValue(fields[1:])
			case "bot":
				doParseHiLo(fields[1:])
			default:
				fmt.Println("Something else:", fields)
			}
		}
	}

	i := 0
	for {
		i++
		if !runBotCycle() {
			fmt.Println("Finished in", i, "cycles")
			break
		}
	}
	i = 0
	out := 0
	for i < 3 {
		for _, v := range bins[i].values {
			if out == 0 {
				out = v
			} else {
				out *= v
			}
		}
		i++
	}
	fmt.Println("Multiplication of bin 1-2-3 is:", out)
}
