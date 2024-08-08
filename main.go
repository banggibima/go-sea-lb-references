package main

import (
	"fmt"

	baseballgame "github.com/go-sea-lb-references/baseball-game"
	countandsay "github.com/go-sea-lb-references/count-and-say"
	differentwaystoaddparentheses "github.com/go-sea-lb-references/different-ways-to-add-parentheses"
	fizzbuzz "github.com/go-sea-lb-references/fizz-buzz"
	longestvalidparentheses "github.com/go-sea-lb-references/longest-valid-parentheses"
	mergesortedarray "github.com/go-sea-lb-references/merge-sorted-array"
	palindromepartitioning "github.com/go-sea-lb-references/palindrome-partitioning"
	replacewords "github.com/go-sea-lb-references/replace-words"
	smalleststringwithagivennumericvalue "github.com/go-sea-lb-references/smallest-string-with-a-given-numeric-value"
	validsudoku "github.com/go-sea-lb-references/valid-sudoku"
)

func main() {
	fmt.Println(longestvalidparentheses.LongestValidParentheses("(()"))

	fmt.Println(countandsay.CountAndSay(4))

	fmt.Println(validsudoku.IsValidSudoku([][]byte{
		[]byte("53..7...."),
		[]byte("6..195..."),
		[]byte(".98....6."),
		[]byte("8...6...3"),
		[]byte("4..8.3..1"),
		[]byte("7...2...6"),
		[]byte(".6....28."),
		[]byte("...419..5"),
		[]byte("....8..79"),
	}))

	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}

	mergesortedarray.Merge(nums1, 3, nums2, 3)

	fmt.Println(nums1)

	fmt.Println(palindromepartitioning.Partition("aab"))

	fmt.Println(differentwaystoaddparentheses.DiffWaysToCompute("2-1-1"))

	fmt.Println(fizzbuzz.FizzBuzz(15))

	fmt.Println(replacewords.ReplaceWords([]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery"))

	fmt.Println(baseballgame.CalPoints([]string{"5", "2", "C", "D", "+"}))

	fmt.Println(smalleststringwithagivennumericvalue.GetSmallestString(3, 27))
}
