package main

func main() {

}

func isValidSudoku(board [][]byte) bool {

	var row, cols, box []map[byte]bool
	for i := 0; i < 9; i++ {
		row = append(row, make(map[byte]bool))
		cols = append(cols, make(map[byte]bool))
		box = append(box, make(map[byte]bool))
	}

	for key, value := range board {
		for bk, bv := range value {
			if bv == '.' {
				continue
			}
			if row[key][bv] || cols[bk][bv] || box[key/3*3+bk/3][bv] {
				return false
			}
			row[key][bv] = true
			cols[bk][bv] = true
			box[key/3*3+bk/3][bv] = true
		}
	}

	return true
}

// Input: board =
// [["8","3",".",".","7",".",".",".","."]
// ,["6",".",".","1","9","5",".",".","."]
// ,[".","9","8",".",".",".",".","6","."]
// ,["8",".",".",".","6",".",".",".","3"]
// ,["4",".",".","8",".","3",".",".","1"]
// ,["7",".",".",".","2",".",".",".","6"]
// ,[".","6",".",".",".",".","2","8","."]
// ,[".",".",".","4","1","9",".",".","5"]
// ,[".",".",".",".","8",".",".","7","9"]]
// Output: false
// Explanation: Same as Example 1, except with the 5 in the top left corner being modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.
