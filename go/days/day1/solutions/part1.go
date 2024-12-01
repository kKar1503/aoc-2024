package solutions

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func Part1(input []string) []byte {
	sum := 0
	left := make([]int, len(input))
	right := make([]int, len(input))
	for i, line := range input {
		splits := strings.Split(line, "   ")
		right[i], _ = strconv.Atoi(splits[1])
		left[i], _ = strconv.Atoi(splits[0])
	}
	sort.Ints(left)
	sort.Ints(right)
	for i := 0; i < len(left); i++ {
		dist := int(math.Abs(float64(left[i]) - float64(right[i])))
		fmt.Println(left[i], right[i])
		sum += dist
	}

	return []byte(strconv.Itoa(sum))
}
