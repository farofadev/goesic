package utils

import (
	"fmt"
	"strconv"
	"time"
)

func GetAtoi(a interface{}) int {
	if fmt.Sprintf("%T", a) == "string" {
		value, _ := strconv.Atoi(a.(string))

		return value
	}

	return a.(int)
}

func FormatDateTimeString(t time.Time) string {

	return t.Format("06-01-02 15:04:05")


}