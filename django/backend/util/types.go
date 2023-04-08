package util

import "strconv"

func ParseInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		i = -1
	}
	return i
}
