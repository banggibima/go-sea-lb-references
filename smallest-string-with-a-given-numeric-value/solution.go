package smalleststringwithagivennumericvalue

func GetSmallestString(n int, k int) string {
	result := make([]byte, n)

	for i := n - 1; i >= 0; i-- {
		maxValue := k - i
		if maxValue > 26 {
			maxValue = 26
		}
		result[i] = byte('a' + maxValue - 1)
		k -= maxValue
	}

	return string(result)
}
