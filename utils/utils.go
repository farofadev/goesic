package utils

import (
	"fmt"
	"strconv"
)

func GetAtoi(a interface{}) (int) {
	if fmt.Sprintf("%T", a) == "string" {
		value, _ := strconv.Atoi(a.(string))

		return value
	}

	return a.(int)
}
