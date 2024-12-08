package utils

import "strconv"

func StringToUint(s string) uint64 {
	u, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0
	}
	return u
}
