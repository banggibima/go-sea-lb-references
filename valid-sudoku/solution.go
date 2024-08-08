package validsudoku

func IsValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	boxes := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			cell := board[i][j]
			if cell == '.' {
				continue
			}

			if rows[i][cell] {
				return false
			}
			rows[i][cell] = true

			if cols[j][cell] {
				return false
			}
			cols[j][cell] = true

			boxIndex := (i/3)*3 + j/3
			if boxes[boxIndex][cell] {
				return false
			}
			boxes[boxIndex][cell] = true
		}
	}

	return true
}
