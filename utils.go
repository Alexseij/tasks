package tasks

import "fmt"

func PrintQueue(queue []*Ball) {
	for i := 0; i < len(queue); i++ {
		fmt.Print(queue[i].Value, " ")
	}
	fmt.Println()
}

func PrintMatrix(mask [][]byte) {
	for i := 0; i < HEIGHT_OF_GAME_PLACE; i++ {
		for j := 0; j < WIDTH_OF_GAME_PLACE; j++ {
			fmt.Print(mask[i][j], " ")
		}
		fmt.Println()
	}
}

func PrintBoard(board Board) {
	for j := 0; j < HEIGHT_OF_GAME_PLACE; j++ {
		for k := 0; k < WIDTH_OF_GAME_PLACE; k++ {
			if board[j][k] == nil {
				fmt.Print("n ")
				continue
			}
			fmt.Print(board[j][k].Value, " ")
		}
		fmt.Println()
	}
}

func CheckIsAllNilColumn(column int, board Board) bool {
	for i := 0; i < HEIGHT_OF_GAME_PLACE; i++ {
		if board[i][column] != nil {
			return false
		}
	}
	return true
}

func CheckIsAllNilRow(row int, board Board) bool {
	for i := 0; i < WIDTH_OF_GAME_PLACE; i++ {
		if board[row][i] != nil {
			return false
		}
	}
	return true
}
