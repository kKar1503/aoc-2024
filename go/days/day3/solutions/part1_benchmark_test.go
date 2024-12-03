package solutions_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/kKar1503/aoc-2024/days/day3/solutions"
	"github.com/kKar1503/aoc-2024/helper"
)

func BenchmarkD3P1StringSearch(b *testing.B) {
	input, err := helper.ImportAsString("../input")
	if err != nil {
		fmt.Printf("Failed reading the file: %s", err.Error())
		os.Exit(1)
	}

	for i := 0; i < b.N; i++ {
		solutions.Part1StringSearch(input)
	}
}

func BenchmarkD3P1RegexFindAllSubmatch(b *testing.B) {
	input, err := helper.ImportAsString("../input")
	if err != nil {
		fmt.Printf("Failed reading the file: %s", err.Error())
		os.Exit(1)
	}

	for i := 0; i < b.N; i++ {
		solutions.Part1RegexFindAllSubmatch(input)
	}
}

func BenchmarkD3P1RegexFindStringSubmatchIndex(b *testing.B) {
	input, err := helper.ImportAsString("../input")
	if err != nil {
		fmt.Printf("Failed reading the file: %s", err.Error())
		os.Exit(1)
	}

	for i := 0; i < b.N; i++ {
		solutions.Part1RegexFindStringSubmatchIndex(input)
	}
}
