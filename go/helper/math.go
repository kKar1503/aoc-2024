package helper

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Float | constraints.Integer
}

func Abs[T Number](n T) T {
	if n < 0 {
		return -n
	}

	return n
}

func GCD[T constraints.Integer](a, b T) T {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}

func lcm[T constraints.Integer](a, b T) T {
	return a * b / GCD(a, b)
}

func LCM[T constraints.Integer](vs ...T) T {
	if len(vs) < 2 {
		panic("need at least 2 inputs")
	}

	result := vs[0]
	for i := 1; i < len(vs); i++ {
		result = lcm(result, vs[i])
	}

	return result
}

func Sign[T Number](v T) int {
	switch {
	case v < 0:
		return -1
	case v == 0:
		return 0
	default:
		return 1
	}
}
