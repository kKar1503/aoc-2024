package solutions

import (
	"strconv"
)

func isMasOrReverse(str ...byte) bool {
	if len(str) != 3 {
		return false
	}
	return (str[0] == 'M' && str[1] == 'A' && str[2] == 'S') || (str[0] == 'S' && str[1] == 'A' && str[2] == 'M')
}

func countMasInX(input []string, x, y int) int {
	if x-1 < 0 || x+1 >= len(input[0]) || y-1 < 0 || y+1 >= len(input) {
		return 0
	}

	// check top left to bottom right
	if isMasOrReverse(input[y-1][x-1], input[y][x], input[y+1][x+1]) &&
		// check top right to bottom left
		isMasOrReverse(input[y-1][x+1], input[y][x], input[y+1][x-1]) {
		return 1
	}

	return 0
}

func Part2(input []string) []byte {
	count := 0
	for y, line := range input {
		for x := 0; x < len(line); x++ {
			if line[x] == 'A' {
				count += countMasInX(input, x, y)
			}
		}
	}

	return []byte(strconv.Itoa(count))
}
