package main

import (
	"aoc_2021/helpers"
	"fmt"
	"strconv"
	"strings"
)

func one(data [][]int) int64 {
	var tmpStringGamma, tmpStringEpislon string

	for i := 0; i < len(data[0]); i++ {
		var tmp int
		for _, e := range data {
			tmp += e[i]
		}
		if tmp >= len(data)/2 {
			tmpStringGamma += "1"
			tmpStringEpislon += "0"
		} else {
			tmpStringGamma += "0"
			tmpStringEpislon += "1"
		}
	}

	gammaRate, _ := strconv.ParseInt(tmpStringGamma, 2, 64)
	epsilonRate, _ := strconv.ParseInt(tmpStringEpislon, 2, 64)

	return gammaRate * epsilonRate
}

func two(data [][]int) int {
	return 0
}

func parseData(data []string) [][]int {
	var out [][]int

	for _, e := range data {
		var tmp []int
		vals := strings.Split(e, "")
		for _, i := range vals {
			val, _ := strconv.Atoi(i)
			tmp = append(tmp, val)
		}
		out = append(out, tmp)
	}

	return out
}

func main() {
	data := helpers.ReadFile("input.txt")
	// data := helpers.ReadFile("input_test.txt")
	parsedData := parseData(data)
	fmt.Println("One:", one(parsedData))
	fmt.Println("Two:", two(parsedData))
}
