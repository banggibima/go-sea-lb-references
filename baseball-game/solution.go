package baseballgame

import (
	"strconv"
)

func CalPoints(operations []string) int {
	result := []int{}

	for _, op := range operations {
		switch op {
		case "+":
			result = append(result, result[len(result)-1]+result[len(result)-2])
		case "D":
			result = append(result, result[len(result)-1]*2)
		case "C":
			result = result[:len(result)-1]
		default:
			score, _ := strconv.Atoi(op)
			result = append(result, score)
		}
	}

	sum := 0
	for _, r := range result {
		sum += r
	}

	return sum
}
