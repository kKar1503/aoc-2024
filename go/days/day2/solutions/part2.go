package solutions

import (
	"strings"

	"github.com/kKar1503/aoc-2024/helper"
)

func Part2(input []string) []byte {
	safe := 0
	for _, line := range input {
		levels := strings.Split(line, " ")
		levelsInNumber := helper.MustInts(levels)

		if isSafe(levelsInNumber) {
			safe++
			continue
		}

		for i := 0; i < len(levelsInNumber); i++ {
			if isSafe(helper.RemoveByIndex(levelsInNumber, i)) {
				safe++
				break
			}
		}
	}

	return []byte(helper.SignedStr(safe))
}

// func Part2(input []string) []byte {
// 	var result int
//
// 	for _, line := range input {
// 		levels := strings.Split(line, " ")
// 		levelsInNumber := helper.MustInts(levels)
//
// 		if isSafe(levelsInNumber) {
// 			result++
// 			continue
// 		}
//
// 		for i := 0; i < len(levelsInNumber); i++ {
// 			clone := make([]int, len(levelsInNumber))
// 			copy(clone, levelsInNumber)
//
// 			if isSafe(append(clone[:i], clone[i+1:]...)) {
// 				result++
// 				break
// 			}
// 		}
// 	}
//
// 	return []byte(helper.SignedStr(result))
// }
