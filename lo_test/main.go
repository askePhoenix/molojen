package main

import (
	stream "github.com/askePhoenix/molojen/jen_test/test"
	"github.com/askePhoenix/molojen/lo_test/fn"
)

func main() {
	test := stream.IntArray{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	test.
		ToStringArray().
		Map(fn.Concat).
		ForEach(fn.FuncPrintString)
}
