package main

import "fmt"

func main() {

	value := [][]byte{
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

	isValidSudoku(value)

}

func isValidSudoku(board [][]byte) bool {
	// use a map to compare the rows
	// use a map to compare the columns
	// step through each row in the array and add to the map
	// If a duplicate is found, return false
	// do the same for columns

	stephensMap := make(map[byte][][]int)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			locs := stephensMap[board[i][j]]

			for k := range locs {
				// check 3 conditions
				if locs[k][0] == i {
					fmt.Println("false")
					return false
				}
				if locs[k][1] == j {
					fmt.Println("false")
					return false
				}
				if locs[k][0]/3 == i/3 && locs[k][1]/3 == j/3 {
					fmt.Println("false")
					return false
				}
			}
			stephensMap[board[i][j]] = append(stephensMap[board[i][j]], []int{i, j})
		}
	}

	fmt.Println("true")
	return true
}
