package differentwaystoaddparentheses

import (
	"strconv"
)

func DiffWaysToCompute(expression string) []int {
	result := []int{}

	for i, c := range expression {
		if c == '+' || c == '-' || c == '*' {
			left := DiffWaysToCompute(expression[:i])
			right := DiffWaysToCompute(expression[i+1:])

			for _, l := range left {
				for _, r := range right {
					switch c {
					case '+':
						result = append(result, l+r)
					case '-':
						result = append(result, l-r)
					case '*':
						result = append(result, l*r)
					}
				}
			}
		}
	}

	if len(result) == 0 {
		if num, err := strconv.Atoi(expression); err == nil {
			result = append(result, num)
		}
	}

	return result
}
