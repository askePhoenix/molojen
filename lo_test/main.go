package main

import (
	"fmt"
	"github.com/samber/lo"
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

func main() {
	test := IntArray{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	test.
		Map(funcMultiple2).
		Map(funcMultiple2).
		ForEach(funcPrintInt)
}

func funcMultiple2(i, _ int) int {
	return i * 2
}

func funcPrintInt(i, _ int) {
	fmt.Println(i)
}
