package utils

import "strconv"

func StringToInt64(text string) (int, error) {
	return strconv.Atoi(text)
}
