package main

import (
	"aoc_2021/helpers"
	"fmt"
	"strconv"
)

func one(data []int) int {
	var output int
	for i, e := range data {
		if i != 0 { // Skipping first
			if e > data[i-1] {
				output++
			}
		}
	}

	return output
}

func two(data []int) int {
	var output, lastVal int

	for i := 0; i < (len(data) - 2); i++ {
		currVal := data[i] + data[i+1] + data[i+2]
		if lastVal != 0 {
			if currVal > lastVal {
				output++
			}
			lastVal = currVal
		}
		lastVal = currVal
	}

	return output
}

func parseData(data []string) []int {
	var output []int
	for _, e := range data {
		val, _ := strconv.Atoi(e)
		output = append(output, val)
	}

	return output
}

func main() {
	// data := helpers.ReadFile("input_test.txt")
	data := helpers.ReadFile("input.txt")
	parsedData := parseData(data)
	fmt.Println("First solution:", one(parsedData))
	fmt.Println("Second solution:", two(parsedData))
}
