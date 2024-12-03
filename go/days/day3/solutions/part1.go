package solutions

import (
	"regexp"
	"strconv"

	"github.com/kKar1503/aoc-2024/helper"
)

func parseNumber(input string, startingIdx int) (num, endingIdx int) {
	for i := startingIdx; i < len(input); i++ {
		if input[i] >= '0' && input[i] <= '9' {
			num = num*10 + int(input[i]-'0')
		} else {
			return num, i
		}
	}
	return -1, -1
}

func parseMultiplication(input string, startingIdx int) (num1, num2, endingIdx int) {
	if startingIdx+4 >= len(input) {
		return -1, -1, -1
	}
	if input[startingIdx:startingIdx+4] != "mul(" {
		return -1, -1, -1
	}
	num, endingIdx := parseNumber(input, startingIdx+4)
	if endingIdx == -1 {
		return -1, -1, -1
	}
	num1 = num
	if input[endingIdx] != ',' {
		return -1, -1, -1
	}
	num, endingIdx = parseNumber(input, endingIdx+1)
	if endingIdx == -1 {
		return -1, -1, -1
	}
	num2 = num
	if input[endingIdx] != ')' {
		return -1, -1, -1
	}
	return num1, num2, endingIdx + 1
}

func Part1StringSearch(input string) int {
	idx := 0
	sum := 0

	for idx < len(input) {
		if input[idx] != 'm' {
			idx++
			continue
		}

		num1, num2, endingIdx := parseMultiplication(input, idx)
		if endingIdx == -1 {
			idx++
			continue
		}
		sum += num1 * num2
		idx = endingIdx
	}

	return sum
}

func Part1RegexFindAllSubmatch(input string) int {
	reg := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	results := reg.FindAllSubmatch([]byte(input), -1)
	if results == nil {
		return 0
	}

	sum := 0
	for _, result := range results {
		sum += helper.MustInt(string(result[1])) * helper.MustInt(string(result[2]))
	}

	return sum
}

func Part1RegexFindStringSubmatchIndex(input string) int {
	reg := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	sum := 0

	for loc := reg.FindStringSubmatchIndex(input); loc != nil; loc = reg.FindStringSubmatchIndex(input) {
		sum += helper.MustInt(input[loc[2]:loc[3]]) * helper.MustInt(input[loc[4]:loc[5]])
		input = input[loc[5]:]
	}

	return sum
}

func Part1(input string) []byte {
	return []byte(strconv.Itoa(Part1RegexFindStringSubmatchIndex(input)))
}
