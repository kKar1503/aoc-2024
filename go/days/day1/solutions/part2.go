package solutions

import (
	"strconv"
	"strings"
)

func Part2(input []string) []byte {
	left := make([]int, len(input))
	rightCount := make(map[int]int)
	for i, line := range input {
		splits := strings.Split(line, "   ")
		right, _ := strconv.Atoi(splits[1])
		if _, ok := rightCount[right]; !ok {
			rightCount[right] = 1
		} else {
			rightCount[right]++
		}
		left[i], _ = strconv.Atoi(splits[0])
	}
	similarityScore := 0
	for i := 0; i < len(left); i++ {
		count, ok := rightCount[left[i]]
		if ok {
			similarityScore += count * left[i]
		}
	}
	return []byte(strconv.Itoa(similarityScore))
}
