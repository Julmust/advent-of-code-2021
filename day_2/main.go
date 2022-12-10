package main

import (
	"aoc_2021/helpers"
	"fmt"
	"strconv"
	"strings"
)

func one(data []string) int {
	var horPos, depth int

	for _, e := range data {
		command := strings.Split(e, " ")
		change, _ := strconv.Atoi(command[1])

		switch command[0] {
		case "forward":
			horPos += change
		case "down":
			depth += change
		case "up":
			depth -= change
		}
	}

	return horPos * depth
}

func two(data []string) int {
	var horPos, depth, aim int

	for _, e := range data {
		command := strings.Split(e, " ")
		change, _ := strconv.Atoi(command[1])

		switch command[0] {
		case "forward":
			horPos += change
			depth += aim * change
		case "down":
			aim += change
		case "up":
			aim -= change
		}
	}

	return horPos * depth
}

func main() {
	data := helpers.ReadFile("input.txt")
	// data := helpers.ReadFile("input_test.txt")
	fmt.Println("Solution one:", one(data))
	fmt.Println("Solution two:", two(data))
}
