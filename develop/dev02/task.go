package main

import (
	"errors"
	"strconv"
	"strings"
)

var (
	InvalidInputError = errors.New("invalid string")
)

func UnpackStr(str string) (string, error) {
	var (
		resultStr strings.Builder
		count     int
		symbol    int32 = 0
	)
	for _, i := range str {
		if i > '0' && i <= '9' {
			if symbol == 0 {
				return "", InvalidInputError
			}
			count, _ = strconv.Atoi(string(i))
			for index := 0; index < count; index++ {
				resultStr.WriteRune(symbol)
			}
			symbol = 0
		} else {
			if symbol != 0 {
				resultStr.WriteRune(symbol)
			}
			symbol = i
		}
	}
	if symbol != 0 {
		resultStr.WriteRune(symbol)
	}
	return resultStr.String(), nil
}
