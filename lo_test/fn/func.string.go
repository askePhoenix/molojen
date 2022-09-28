package fn

import "strconv"

func IntToString(i int, _ int) string {
	return strconv.Itoa(i)
}

func Concat(i string, _ int) string {
	return i + " - test"
}
