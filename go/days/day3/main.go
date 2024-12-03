package main

import (
	"fmt"
	"os"

	"github.com/kKar1503/aoc-2024/days/day3/solutions"
	"github.com/kKar1503/aoc-2024/helper"
)

const day = 3

var parts = map[string]solutionFn{
	"part1": solutions.Part1,
	"part2": solutions.Part2,
}

type solutionFn = func(string) []byte

func main() {
	input, err := helper.ImportAsString(fmt.Sprintf("days/day%d/input", day))
	if err != nil {
		fmt.Printf("Failed reading the file: %s", err.Error())
		os.Exit(1)
	}

	part := os.Args[1]

	output := parts[part](input)

	helper.Output(output, fmt.Sprintf("days/day%d/output-%s", day, part))
}
