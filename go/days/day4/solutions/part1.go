package solutions

import (
	"strconv"
)

func isXmas(str ...byte) bool {
	if len(str) != 4 {
		return false
	}
	return str[0] == 'X' && str[1] == 'M' && str[2] == 'A' && str[3] == 'S'
}

func countCircularXmas(input []string, x, y int) int {
	hasSpaceLeft := x-3 >= 0
	hasSpaceRight := x+3 < len(input[0])
	hasSpaceUp := y-3 >= 0
	hasSpaceDown := y+3 < len(input)
	count := 0
	// check left to right
	if hasSpaceRight {
		if isXmas(input[y][x], input[y][x+1], input[y][x+2], input[y][x+3]) {
			count++
		}
	}

	// check right to left
	if hasSpaceLeft {
		if isXmas(input[y][x], input[y][x-1], input[y][x-2], input[y][x-3]) {
			count++
		}
	}

	// check up to down
	if hasSpaceDown {
		if isXmas(input[y][x], input[y+1][x], input[y+2][x], input[y+3][x]) {
			count++
		}
	}

	// check down to up
	if hasSpaceUp {
		if isXmas(input[y][x], input[y-1][x], input[y-2][x], input[y-3][x]) {
			count++
		}
	}

	// check diagonal top left to bottom right
	if hasSpaceRight && hasSpaceDown {
		if isXmas(input[y][x], input[y+1][x+1], input[y+2][x+2], input[y+3][x+3]) {
			count++
		}
	}

	// check diagonal bottom left to top right
	if hasSpaceRight && hasSpaceUp {
		if isXmas(input[y][x], input[y-1][x+1], input[y-2][x+2], input[y-3][x+3]) {
			count++
		}
	}

	// check diagonal top right to bottom left
	if hasSpaceLeft && hasSpaceDown {
		if isXmas(input[y][x], input[y+1][x-1], input[y+2][x-2], input[y+3][x-3]) {
			count++
		}
	}

	// check diagonal bottom right to top left
	if hasSpaceLeft && hasSpaceUp {
		if isXmas(input[y][x], input[y-1][x-1], input[y-2][x-2], input[y-3][x-3]) {
			count++
		}
	}
	return count
}

func Part1(input []string) []byte {
	count := 0
	for y, line := range input {
		for x := 0; x < len(line); x++ {
			if line[x] == 'X' {
				count += countCircularXmas(input, x, y)
			}
		}
	}

	return []byte(strconv.Itoa(count))
}
