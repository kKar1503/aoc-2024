package solutions

import (
	"fmt"
	"strconv"
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

func Part1(input string) []byte {
	idx := 0
	sum := 0
	for idx < len(input) {
		fmt.Printf("idx = %d, input[idx]=%c | ", idx, input[idx])
		num1, num2, endingIdx := parseMultiplication(input, idx)
		if endingIdx == -1 {
			idx++
			fmt.Println("skipping")
			continue
		}
		fmt.Printf("num1 = %d, num2 = %d, parsed=%s\n", num1, num2, input[idx:endingIdx])
		sum += num1 * num2
		idx = endingIdx
	}
	return []byte(strconv.Itoa(sum))
}
