package countandsay

import (
	"strconv"
)

func CountAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	prev := CountAndSay(n - 1)
	result := ""
	count := 1

	for i := 0; i < len(prev); i++ {
		if i == len(prev)-1 || prev[i] != prev[i+1] {
			result += strconv.Itoa(count) + string(prev[i])
			count = 1
		} else {
			count++
		}
	}

	return result
}
