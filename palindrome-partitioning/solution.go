package palindromepartitioning

func IsPalindrome(s string) bool {
	i := 0
	j := len(s) - 1

	for i < j {
		if s[i] != s[j] {
			return false
		}

		i++
		j--
	}

	return true
}

func Backtrack(s string, path []string, result *[][]string) {
	if len(s) == 0 {
		*result = append(*result, append([]string{}, path...))
		return
	}

	for i := 1; i <= len(s); i++ {
		if IsPalindrome(s[:i]) {
			Backtrack(s[i:], append(path, s[:i]), result)
		}
	}
}

func Partition(s string) [][]string {
	result := [][]string{}

	Backtrack(s, []string{}, &result)

	return result
}
