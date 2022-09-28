package fn

import "strconv"

func IntToString(i int, _ int) string {
	return strconv.Itoa(i)
}

func StringToInt(s string, _ int) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func Concat(i string, _ int) string {
	return i + " - test"
}
