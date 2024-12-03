package solutions

import (
	"fmt"
	"strconv"
)

func parseDo(input string, startingIdx int) (found bool, endingIdx int) {
	if startingIdx+4 >= len(input) {
		return false, -1
	}
	if input[startingIdx:startingIdx+4] != "do()" {
		return false, -1
	}
	return true, startingIdx + 4
}

func parseDont(input string, startingIdx int) (found bool, endingIdx int) {
	if startingIdx+7 >= len(input) {
		return false, -1
	}
	if input[startingIdx:startingIdx+7] != "don't()" {
		return false, -1
	}
	return true, startingIdx + 7
}

func Part2(input string) []byte {
	idx := 0
	sum := 0
	inDo := true
	inDont := false
	for idx < len(input) {
		fmt.Printf("idx = %d, input[idx]=%c | ", idx, input[idx])
		hasDo, endingIdx := parseDo(input, idx)
		if hasDo {
			fmt.Printf("do found, %s\n", input[idx:endingIdx])
			inDo = true
			inDont = false
			idx = endingIdx
			continue
		}
		hasDont, endingIdx := parseDont(input, idx)
		if hasDont {
			fmt.Printf("dont found, %s\n", input[idx:endingIdx])
			inDont = true
			inDo = false
			idx = endingIdx
			continue
		}

		num1, num2, endingIdx := parseMultiplication(input, idx)
		if endingIdx == -1 {
			idx++
			fmt.Printf("skipping\n")
			continue
		}
		fmt.Printf("num1 = %d, num2 = %d, parsed=%s, ", num1, num2, input[idx:endingIdx])
		if inDont {
			fmt.Printf("skipping\n")
			idx = endingIdx
			continue
		}
		if inDo {
			fmt.Printf("doing\n")
			sum += num1 * num2
			idx = endingIdx
			continue
		}
	}
	return []byte(strconv.Itoa(sum))
}
