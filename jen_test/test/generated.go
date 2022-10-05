package stream

import (
	fn "github.com/askePhoenix/molojen/lo_test/fn"
	lo "github.com/samber/lo"
)

type IntArray []int
type FuncIntArray func(int, int) int
type FuncIntArrayConsumer func(int, int)

func (i IntArray) Map(f FuncIntArray) IntArray {
	return lo.Map[int, int](i, f)
}
func (i IntArray) ForEach(f FuncIntArrayConsumer) {
	lo.ForEach[int](i, f)
}
func (i IntArray) ToStringArray() StringArray {
	return lo.Map[int, string](i, fn.IntToString)
}

type StringArray []string
type FuncStringArray func(string, int) string
type FuncStringArrayConsumer func(string, int)

func (s StringArray) Map(f FuncStringArray) StringArray {
	return lo.Map[string, string](s, f)
}
func (s StringArray) ForEach(f FuncStringArrayConsumer) {
	lo.ForEach[int](s, f)
}
func (s StringArray) ToStringArray() IntArray {
	return lo.Map[string, int](s, fn.StringToInt)
}
