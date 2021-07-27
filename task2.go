package tasks

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
)

type Coords struct {
	X, Y int
}

type Ball struct {
	Coords
	Value string
	up    *Ball
	down  *Ball
	right *Ball
	left  *Ball
}

type Board [][]*Ball

const (
	WIDTH_OF_GAME_PLACE  = 15
	HEIGHT_OF_GAME_PLACE = 10
)

func readGames(rd *bufio.Reader) (int, map[int]Board, error) {

	line, _, err := rd.ReadLine()
	if err != nil {
		return 0, nil, err
	}
	countOfGamesStr := strings.TrimSpace(string(line))

	countOfGames, err := strconv.Atoi(countOfGamesStr)
	if err != nil {
		return 0, nil, err
	}

	line, _, err = rd.ReadLine()
	if err != nil {
		return 0, nil, err
	}
	m := make(map[int]Board, countOfGames)

	for i := 0; i < countOfGames; i++ {
		board := make(Board, HEIGHT_OF_GAME_PLACE)
		for j := 0; j < HEIGHT_OF_GAME_PLACE; j++ {
			line, _, err = rd.ReadLine()
			if err != nil {
				return 0, nil, err
			}
			row := make([]*Ball, WIDTH_OF_GAME_PLACE)
			for k := 0; k < WIDTH_OF_GAME_PLACE; k++ {
				row[k] = &Ball{
					Coords: Coords{
						X: k,
						Y: j,
					},
					Value: string(line[k]),
					up:    nil,
					down:  nil,
					right: nil,
					left:  nil,
				}
			}
			board[j] = row
		}
		m[i] = board
		line, _, err = rd.ReadLine()
		if err != nil && err != io.EOF {
			return 0, nil, err
		}
	}

	return countOfGames, m, nil
}

func addPoints(removed int) int {
	return int(math.Pow(float64(removed-2), 2))
}

func initClusterCoords() [][]byte {
	clusterCoords := make([][]byte, HEIGHT_OF_GAME_PLACE)
	for i := 0; i < HEIGHT_OF_GAME_PLACE; i++ {
		row := make([]byte, WIDTH_OF_GAME_PLACE)
		clusterCoords[i] = row
	}
	return clusterCoords
}

func BindBoard(m map[int]Board) {
	for _, board := range m {
		for i := 0; i < HEIGHT_OF_GAME_PLACE; i++ {
			for j := 0; j < WIDTH_OF_GAME_PLACE; j++ {
				if j == WIDTH_OF_GAME_PLACE-1 && i == HEIGHT_OF_GAME_PLACE-1 {
					break
				} else if j == WIDTH_OF_GAME_PLACE-1 {
					board[i][j].down = board[i+1][j]
					board[i+1][j].up = board[i][j]
					continue
				} else if i == HEIGHT_OF_GAME_PLACE-1 {
					board[i][j].right = board[i][j+1]
					board[i][j+1].left = board[i][j]
					continue
				}
				board[i][j].right = board[i][j+1]
				board[i][j+1].left = board[i][j]
				board[i+1][j].up = board[i][j]
				board[i][j].down = board[i+1][j]
			}
		}
	}
}

func findClusterForCurrentBall(clusterCoords [][]byte, ball *Ball, size *int) {
	queue := make([]*Ball, 0)
	queue = append(queue, ball)

	for len(queue) != 0 {
		currentBall := queue[0]
		queue = queue[1:]
		if currentBall.left != nil && currentBall.left.Value == currentBall.Value && clusterCoords[currentBall.left.Y][currentBall.left.X] == 0 {
			*(size)++
			queue = append(queue, currentBall.left)
			clusterCoords[currentBall.left.Y][currentBall.left.X] = 1
		}
		if currentBall.right != nil && currentBall.right.Value == currentBall.Value && clusterCoords[currentBall.right.Y][currentBall.right.X] == 0 {
			*(size)++
			queue = append(queue, currentBall.right)
			clusterCoords[currentBall.right.Y][currentBall.right.X] = 1
		}
		if currentBall.down != nil && currentBall.down.Value == currentBall.Value && clusterCoords[currentBall.down.Y][currentBall.down.X] == 0 {
			*(size)++
			queue = append(queue, currentBall.down)
			clusterCoords[currentBall.down.Y][currentBall.down.X] = 1
		}
		if currentBall.up != nil && currentBall.up.Value == currentBall.Value && clusterCoords[currentBall.up.Y][currentBall.up.X] == 0 {
			*(size)++
			queue = append(queue, currentBall.up)
			clusterCoords[currentBall.up.Y][currentBall.up.X] = 1
		}
	}

}

func findLeftMostBall(coords [][]byte, board Board) (*Ball, error) {
	for i := 0; i < WIDTH_OF_GAME_PLACE; i++ {
		for j := 0; j < HEIGHT_OF_GAME_PLACE; j++ {
			if coords[j][i] == 1 {
				return board[j][i], nil
			}
		}
	}
	return nil, fmt.Errorf("Have no balls in cluster")
}

func findBottomMostBall(ball *Ball) *Ball {
	currentBall := ball
	for currentBall.down != nil && currentBall.down.Value == currentBall.Value {
		currentBall = currentBall.down
	}
	return currentBall
}

func findLargestCluster(board Board) (int, *Ball, [][]byte, error) {

	sizeOfMaxCluester := 0
	var maxClusterCoords [][]byte
	var maxBall *Ball

	for i := 0; i < HEIGHT_OF_GAME_PLACE; i++ {
		for j := 0; j < WIDTH_OF_GAME_PLACE; j++ {
			if board[i][j] == nil || (maxClusterCoords != nil && maxClusterCoords[board[i][j].Y][board[i][j].X] == 1) {
				continue
			}

			currentSize := 1
			currentClusterCoords := initClusterCoords()
			currentClusterCoords[board[i][j].Y][board[i][j].X] = 1
			findClusterForCurrentBall(currentClusterCoords, board[i][j], &currentSize)

			if currentSize > sizeOfMaxCluester {
				sizeOfMaxCluester = currentSize
				maxClusterCoords = currentClusterCoords
				ball, err := findLeftMostBall(currentClusterCoords, board)
				if err != nil {
					return 0, nil, nil, err
				}
				maxBall = findBottomMostBall(ball)

			} else if currentSize == sizeOfMaxCluester {
				currentClusterBall, err := findLeftMostBall(currentClusterCoords, board)
				if err != nil {
					return 0, nil, nil, err
				}
				maxClusterBall, err := findLeftMostBall(maxClusterCoords, board)
				if err != nil {
					return 0, nil, nil, err
				}

				if (currentClusterBall.X < maxClusterBall.X) || (currentClusterBall.X == maxClusterBall.X && currentClusterBall.Y >= maxClusterBall.Y) {
					maxClusterCoords = currentClusterCoords
					maxBall = findBottomMostBall(currentClusterBall)
				}
			}

		}
	}
	return sizeOfMaxCluester, maxBall, maxClusterCoords, nil

}

func removeCluster(coordsCluster [][]byte, board Board) {
	for i := 0; i < HEIGHT_OF_GAME_PLACE; i++ {
		for j := 0; j < WIDTH_OF_GAME_PLACE; j++ {
			if coordsCluster[i][j] == 1 {
				board[i][j] = nil
			}
		}
	}
}

func checkIsAllNilColumn(column int, board Board) bool {
	for i := 0; i < HEIGHT_OF_GAME_PLACE; i++ {
		if board[i][column] != nil {
			return false
		}
	}
	return true
}

func checkIsAllNilRow(row int, board Board) bool {
	for i := 0; i < WIDTH_OF_GAME_PLACE; i++ {
		if board[row][i] != nil {
			return false
		}
	}
	return true
}

func shiftDown(board Board) {
	for i := 0; i < WIDTH_OF_GAME_PLACE; i++ {
		if checkIsAllNilColumn(i, board) {
			continue
		}
		for j := HEIGHT_OF_GAME_PLACE - 1; j > 0; j-- {
			k := j
			for board[k][i] == nil && k > 0 {
				k--
			}
			if board[k][i] != nil {
				temp := board[k][i]
				board[k][i] = nil
				board[j][i] = temp
				board[j][i].Y = j
				board[j][i].X = i

				if j == HEIGHT_OF_GAME_PLACE-1 {
					board[j][i].down = nil
				} else {
					board[j][i].down = board[j+1][i]
					if board[j+1][i] != nil {
						board[j+1][i].up = board[j][i]
					}
				}

				if j == 0 {
					board[j][i].up = nil
				} else {
					board[j][i].up = board[j-1][i]
					if board[j-1][i] != nil {
						board[j-1][i].down = board[j][i]
					}
				}

				if i == 0 {
					board[j][i].left = nil
				} else {
					board[j][i].left = board[j][i-1]
					if board[j][i-1] != nil {
						board[j][i-1].right = board[j][i]
					}
				}

				if i == WIDTH_OF_GAME_PLACE-1 {
					board[j][i].right = nil
				} else {
					board[j][i].right = board[j][i+1]
					if board[j][i+1] != nil {
						board[j][i+1].left = board[j][i]
					}
				}

			}
		}
	}
}

func calculateWidthOfMask(mask [][]byte) (int, int) {

	fromIndex := 0
	toIndex := 0
loopFromIndex:
	for i := 0; i < WIDTH_OF_GAME_PLACE; i++ {
		for j := 0; j < HEIGHT_OF_GAME_PLACE; j++ {
			if mask[j][i] != 0 {
				fromIndex = i
				break loopFromIndex
			}
		}
	}

	for i := fromIndex; i < WIDTH_OF_GAME_PLACE; i++ {
		isAllZero := true
		for j := 0; j < HEIGHT_OF_GAME_PLACE; j++ {
			if mask[j][i] != 0 {
				isAllZero = false
				break
			}
		}
		if isAllZero {
			toIndex = i
			return fromIndex, toIndex
		}
	}
	return fromIndex, WIDTH_OF_GAME_PLACE
}

func isEmptyColumn(mask [][]byte, board Board) map[int]int {
	fromIndex, toIndex := calculateWidthOfMask(mask)

	cols := make(map[int]int)
	countOfCols := 0
	for i := fromIndex; i < toIndex; i++ {
		isEmptyColumn := true
		for j := 0; j < HEIGHT_OF_GAME_PLACE; j++ {
			if board[j][i] != nil {
				isEmptyColumn = false
				break
			}
		}
		if isEmptyColumn {
			cols[countOfCols] = i
			countOfCols++
		}
	}
	return cols
}

func findFirstColumn(board Board, fromIndex int) (int, error) {
	for i := fromIndex; i < WIDTH_OF_GAME_PLACE; i++ {
		if !checkIsAllNilColumn(i, board) {
			return i, nil
		}
	}
	return 0, fmt.Errorf("have no cols")
}

func swapCols(board Board, from, to int) {
	for i := 0; i < HEIGHT_OF_GAME_PLACE; i++ {
		temp := board[i][from]
		board[i][from] = board[i][to]
		board[i][to] = temp
		if board[i][to] == nil {
			continue
		}
		if to == 0 {
			board[i][to].left = nil
		} else {
			board[i][to].left = board[i][to-1]
			if board[i][to-1] != nil {
				board[i][to-1].right = board[i][to]
			}
		}

		if to == WIDTH_OF_GAME_PLACE-1 {
			board[i][to].right = nil
		} else {
			board[i][to].right = board[i][to+1]
			if board[i][to+1] != nil {
				board[i][to+1].left = board[i][to]
			}

		}

		if i == 0 {
			board[i][to].up = nil
		} else {
			board[i][to].up = board[i-1][to]
			if board[i-1][to] != nil {
				board[i-1][to].down = board[i][to]
			}
		}

		if i == HEIGHT_OF_GAME_PLACE-1 {
			board[i][to].down = nil
		} else {
			board[i][to].down = board[i+1][to]
			if board[i+1][to] != nil {
				board[i+1][to].up = board[i][to]
			}
		}
		board[i][to].X = to
		board[i][to].Y = i
	}
}

func shitfLeft(board Board, mask [][]byte) {
	if cols := isEmptyColumn(mask, board); cols != nil {
		for i := 0; i < len(cols); i++ {
			for j := cols[i] - i; j < WIDTH_OF_GAME_PLACE-i-1; j++ {
				swapCols(board, j+1, j)
			}
		}
	}
}

func compressBoard(board Board, mask [][]byte) {
	shiftDown(board)
	shitfLeft(board, mask)
}

func checkIsAllBallsRemoved(board Board) bool {
	for i := 0; i < HEIGHT_OF_GAME_PLACE; i++ {
		for j := 0; j < WIDTH_OF_GAME_PLACE; j++ {
			if board[i][j] != nil {
				return false
			}
		}
	}
	return true
}

func moveMessage(ball *Ball, removed, numOfMoves, addToScope int) string {
	return fmt.Sprintf(
		"Move %d at (%d , %d) : removed %d balls of color %s , got %d points.\n",
		numOfMoves,
		HEIGHT_OF_GAME_PLACE-ball.Y,
		ball.X+1,
		removed,
		ball.Value,
		addToScope,
	)
}

func checkIsAllClusterHaveOne(board Board) (bool, int) {
	countCountOfClusters := 0

	for i := 0; i < HEIGHT_OF_GAME_PLACE; i++ {
		for j := 0; j < WIDTH_OF_GAME_PLACE; j++ {
			if board[i][j] == nil {
				continue
			}

			value := board[i][j].Value

			left := board[i][j].left
			right := board[i][j].right
			down := board[i][j].down
			up := board[i][j].up

			if (left != nil && left.Value == value) ||
				(right != nil && right.Value == value) ||
				(up != nil && up.Value == value) ||
				(down != nil && down.Value == value) {
				return false, 0
			}

			countCountOfClusters++
		}
	}
	return true, countCountOfClusters
}

func playGame(board Board) error {

	scope := 0
	amountOfMoves := 0
	countOfAloneClusters := 0

	isAllBallsRemoved := false
	isEveryClusteHasOne := false
	isGiveBonus := false

	for !isAllBallsRemoved && !isEveryClusteHasOne {

		if ok, amount := checkIsAllClusterHaveOne(board); ok {
			countOfAloneClusters = amount
			isEveryClusteHasOne = true
			continue
		}

		size, ball, mask, err := findLargestCluster(board)
		if err != nil {
			return err
		}
		removeCluster(mask, board)
		amountOfMoves++
		points := addPoints(size)
		scope += points
		fmt.Print(moveMessage(ball, size, amountOfMoves, points))

		if checkIsAllBallsRemoved(board) {
			isGiveBonus = true
			isAllBallsRemoved = true
			continue
		}

		compressBoard(board, mask)
	}
	if isGiveBonus {
		scope += 1000
	}

	fmt.Printf("Final scope: %d , with %d balls remaining.\n", scope, countOfAloneClusters)
	fmt.Println()
	return nil
}

func Task2(rd *bufio.Reader) error {
	countOfGames, boards, err := readGames(rd)
	if err != nil {
		return err
	}
	BindBoard(boards)

	for i := 0; i < countOfGames; i++ {
		fmt.Printf("Game : %d\n", i+1)
		err = playGame(boards[i])
		if err != nil {
			return err
		}
	}
	return nil
}
