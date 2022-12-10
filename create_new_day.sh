#!/bin/bash
mkdir $1
touch $1/input.txt $1/input_test.txt
echo $'package main\n\nimport (\n\"aoc_2021/helpers\"\n)\n\nfunc one(data []string) {\n\n}\n\nfunc two(data []string) {\n\n}\n\nfunc main() {\n	// data := helpers.ReadFile(\"input.txt\")\n	data := helpers.ReadFile(\"input_test.txt\")\n	one(data)\n	two(data)\n}' > $1/main.go