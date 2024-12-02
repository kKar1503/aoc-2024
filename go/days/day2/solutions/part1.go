package solutions

import (
	"strconv"
	"strings"

	"github.com/kKar1503/aoc-2024/helper"
)

func isSafe(levels []int) bool {
	prev := levels[0]
	dir := levels[1] - prev
	if dir == 0 || helper.Abs(dir) > 3 {
		return false
	}
	prev = levels[1]
	for i := 2; i < len(levels); i++ {
		delta := levels[i] - prev
		if delta == 0 {
			return false
		}
		if helper.Sign(delta) != helper.Sign(dir) {
			return false
		}
		if helper.Abs(delta) > 3 {
			return false
		}
		prev = levels[i]
	}
	return true
}

func Part1(input []string) []byte {
	safe := 0
	for _, line := range input {
		levels := strings.Split(line, " ")
		levelsInNumber := helper.MustInts(levels)

		if isSafe(levelsInNumber) {
			safe++
		}
	}

	return []byte(strconv.Itoa(safe))
}
