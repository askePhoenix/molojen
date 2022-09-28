package fn

import "fmt"

// FuncMultiple2 i 곱하기 2를 출력
func FuncMultiple2(i, _ int) int {
	return i * 2
}

func FuncPrintInt(i, _ int) {
	fmt.Println(i)
}

func FuncPrintString(i string, _ int) {
	fmt.Println(i)
}
