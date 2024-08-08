package validsudoku

import (
	"testing"
)

func TestIsValidSudoku(t *testing.T) {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	got := IsValidSudoku(board)
	want := true
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}

	board = [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	got = IsValidSudoku(board)
	want = false
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}

	board = [][]byte{
		{'.', '.', '.', '.', '5', '.', '.', '1', '.'},
		{'.', '.', '.', '.', '8', '6', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'9', '3', '.', '.', '2', '.', '4', '.', '.'},
		{'.', '.', '7', '.', '.', '.', '3', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '3', '4', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '3', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '5', '2', '.', '.'},
	}
	got = IsValidSudoku(board)
	want = false
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}
}
