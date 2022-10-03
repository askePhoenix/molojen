package main

import lo "github.com/samber/lo"

type IntArray []int
type FuncIntArray func(int, int) int
type FuncIntArrayConsumer func(int, int)

func (i IntArray) Map(f FuncIntArray) IntArray {
	return lo.Map[int, int](i, f)
}
func (i IntArray) ForEach(f FuncIntArrayConsumer) {
	lo.Map[int, int](i, f)
}
