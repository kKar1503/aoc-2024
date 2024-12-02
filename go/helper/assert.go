package helper

import (
	"strconv"

	"golang.org/x/exp/constraints"
)

func UnsignedStr[T constraints.Unsigned](n T) string {
	return strconv.FormatUint(uint64(n), 10)
}

func SignedStr[T constraints.Signed](n T) string {
	return strconv.FormatInt(int64(n), 10)
}

func FloatStr[T constraints.Float](n T) string {
	return strconv.FormatFloat(float64(n), 'f', -1, 64)
}

func MustInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}

func MustUint(s string) uint {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		panic(err)
	}

	if i > 0xFFFFFFFF {
		panic("overflow")
	}

	return uint(i)
}

func MustInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}

	return i
}

func MustUint64(s string) uint64 {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		panic(err)
	}

	return i
}

func MustFloat64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}

	return f
}

func MustInts(ss []string) []int {
	is := make([]int, len(ss))
	for i, s := range ss {
		is[i] = MustInt(s)
	}

	return is
}

func MustUints(ss []string) []uint {
	is := make([]uint, len(ss))
	for i, s := range ss {
		is[i] = MustUint(s)
	}

	return is
}

func MustInt64s(ss []string) []int64 {
	is := make([]int64, len(ss))
	for i, s := range ss {
		is[i] = MustInt64(s)
	}

	return is
}

func MustUint64s(ss []string) []uint64 {
	is := make([]uint64, len(ss))
	for i, s := range ss {
		is[i] = MustUint64(s)
	}

	return is
}

func MustFloat64s(ss []string) []float64 {
	fs := make([]float64, len(ss))
	for i, s := range ss {
		fs[i] = MustFloat64(s)
	}

	return fs
}
