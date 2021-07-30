package tasks

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type Coords struct {
	X, Y int
}

type Ball struct {
	Coords
	Value string
	//Links for balls in different ways
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

// Function using for reading all baorads from input file
// Returns : amount of bords , boards , if function have error it returns error else nil value
func ReadBoards(rd *bufio.Reader) (int, map[int]Board, error) {
	// Read amount of boards number
	line, err := rd.ReadString('\n')
	if err != nil {
		return 0, nil, err
	}
	countOfGamesStr := strings.TrimSpace(string(line))

	countOfGames, err := strconv.Atoi(countOfGamesStr)
	if err != nil {
		return 0, nil, err
	}

	line, err = rd.ReadString('\n')
	if err != nil {
		return 0, nil, err
	}
	m := make(map[int]Board, countOfGames)
	// Read boards
	for i := 0; i < countOfGames; i++ {
		board := make(Board, HEIGHT_OF_GAME_PLACE)
		for j := 0; j < HEIGHT_OF_GAME_PLACE; j++ {
			line, err = rd.ReadString('\n')
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
		line, err = rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return countOfGames, m, nil
			}
			return 0, nil, err
		}
	}

	return countOfGames, m, nil
}

// Function using for calculate points whitch incremets final scope
// Returns : calcualted points
func AddPoints(removed int) int {
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

// Function using for set correct refference between each ball in board .
// Each ball have refference to up , down , left and right ball
func BindBoard(m map[int]Board) {
	for _, board := range m {
		for i := 0; i < HEIGHT_OF_GAME_PLACE; i++ {
			for j := 0; j < WIDTH_OF_GAME_PLACE; j++ {

				//Checking for extreme cases (where can't get back refference)
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

				// Normal case
				board[i][j].right = board[i][j+1]
				board[i][j+1].left = board[i][j]
				board[i+1][j].up = board[i][j]
				board[i][j].down = board[i+1][j]
			}
		}
	}
}

// Function using for searching cluster into board
// Working like BFS alghorithm , but for each ball in deirection up , down , right and left
// For check ball has been checked or not , using clusterCoords mask
func findClusterForCurrentBall(clusterCoords [][]byte, ball *Ball, size *int) {
	//Creating queue
	queue := make([]*Ball, 0)
	queue = append(queue, ball)

	for len(queue) != 0 {
		currentBall := queue[0]
		queue = queue[1:]

		//Left ball
		if currentBall.left != nil && currentBall.left.Value == currentBall.Value && clusterCoords[currentBall.left.Y][currentBall.left.X] == 0 {
			*(size)++
			queue = append(queue, currentBall.left)
			clusterCoords[currentBall.left.Y][currentBall.left.X] = 1
		}

		// Right ball
		if currentBall.right != nil && currentBall.right.Value == currentBall.Value && clusterCoords[currentBall.right.Y][currentBall.right.X] == 0 {
			*(size)++
			queue = append(queue, currentBall.right)
			clusterCoords[currentBall.right.Y][currentBall.right.X] = 1
		}

		//Down ball
		if currentBall.down != nil && currentBall.down.Value == currentBall.Value && clusterCoords[currentBall.down.Y][currentBall.down.X] == 0 {
			*(size)++
			queue = append(queue, currentBall.down)
			clusterCoords[currentBall.down.Y][currentBall.down.X] = 1
		}

		//Up ball
		if currentBall.up != nil && currentBall.up.Value == currentBall.Value && clusterCoords[currentBall.up.Y][currentBall.up.X] == 0 {
			*(size)++
			queue = append(queue, currentBall.up)
			clusterCoords[currentBall.up.Y][currentBall.up.X] = 1
		}
	}

}

// Funciont using for find leftmost ball in cluster
// Returns : index of leftmost ball in cluster , send error if have no leftmost ball
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

// Funciont using for find bottomost ball in cluster
// Returns : index of bottommost ball in cluster
func findBottomMostBall(ball *Ball) *Ball {
	currentBall := ball
	for currentBall.down != nil && currentBall.down.Value == currentBall.Value {
		currentBall = currentBall.down
	}
	return currentBall
}

// Function finding largest cluster in board
// Returns : size of cluster , leftmost and bottommost ball if have error return error , else nil
func FindLargestCluster(board Board) (int, *Ball, [][]byte, error) {

	sizeOfMaxCluester := 0
	var maxClusterCoords [][]byte
	var maxBall *Ball

	for i := 0; i < HEIGHT_OF_GAME_PLACE; i++ {
		for j := 0; j < WIDTH_OF_GAME_PLACE; j++ {
			// Check for nil value in ball , and ball whitch using in current max cluster
			if board[i][j] == nil || (maxClusterCoords != nil && maxClusterCoords[board[i][j].Y][board[i][j].X] == 1) {
				continue
			}

			currentSize := 1
			currentClusterCoords := initClusterCoords()
			currentClusterCoords[board[i][j].Y][board[i][j].X] = 1
			findClusterForCurrentBall(currentClusterCoords, board[i][j], &currentSize)

			// Conditions from task
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

// Function using for remove cluster form board
// Using cluster mask for find indexes
func RemoveCluster(coordsCluster [][]byte, board Board) {
	for i := 0; i < HEIGHT_OF_GAME_PLACE; i++ {
		for j := 0; j < WIDTH_OF_GAME_PLACE; j++ {
			if coordsCluster[i][j] == 1 {
				board[i][j] = nil
			}
		}
	}
}

//Function using for make shif down balls as mach as possible
func shiftDown(board Board) {
	for i := 0; i < WIDTH_OF_GAME_PLACE; i++ {
		//Checking for usless col
		if checkIsAllNilColumn(i, board) {
			continue
		}
		// Staring from bottom and getting ball in that
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
				// After changes we need to rewrite refferences to ball in different directions

				// Down ball
				if j == HEIGHT_OF_GAME_PLACE-1 {
					board[j][i].down = nil
				} else {
					board[j][i].down = board[j+1][i]
					if board[j+1][i] != nil {
						board[j+1][i].up = board[j][i]
					}
				}

				// Up ball
				if j == 0 {
					board[j][i].up = nil
				} else {
					board[j][i].up = board[j-1][i]
					if board[j-1][i] != nil {
						board[j-1][i].down = board[j][i]
					}
				}

				// Left ball
				if i == 0 {
					board[j][i].left = nil
				} else {
					board[j][i].left = board[j][i-1]
					if board[j][i-1] != nil {
						board[j][i-1].right = board[j][i]
					}
				}

				// Right ball
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

//Function using for find empty cols in current board
// Returns index where cluster starts , and where ends
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

// Function find empty columns after removing cluster
// Returns : empty columns
func findEmptyColumns(mask [][]byte, board Board) map[int]int {
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

// Function using for swap columns in board
// That using same idea as a shiftDown() fundtion
func swapCols(board Board, from, to int) {
	for i := 0; i < HEIGHT_OF_GAME_PLACE; i++ {

		temp := board[i][from]
		board[i][from] = board[i][to]
		board[i][to] = temp

		if board[i][to] == nil {
			continue
		}

		board[i][to].X = to
		board[i][to].Y = i

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
	}
}

// Function using for fill empty colums in board
func shitfLeft(board Board, mask [][]byte) {
	if cols := findEmptyColumns(mask, board); cols != nil {
		for i := 0; i < len(cols); i++ {
			for j := cols[i] - i; j < WIDTH_OF_GAME_PLACE-i-1; j++ {
				swapCols(board, j+1, j)
			}
		}
	}
}

// Function using for compress board after deleting cluster
func CompressBoard(board Board, mask [][]byte) {
	shiftDown(board)
	shitfLeft(board, mask)
}

func CheckIsAllBallsRemoved(board Board) bool {
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

// Function using for checking is every cluster has only one ball
func CheckIsAllClusterHaveOne(board Board) (bool, int) {
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

// Main game cycle
// s - final scope  , a - amount of moves , c - count of alone clusters
func PlayGame(board Board) (s int, a int, c int, err error) {
	scope := 0
	amountOfMoves := 0
	//Clusters with one ball
	countOfAloneClusters := 0

	isAllBallsRemoved := false
	// Using for indicate : when every cluster has only one ball
	isEveryClusteHasOne := false
	isGiveBonus := false

	for !isAllBallsRemoved && !isEveryClusteHasOne {

		if ok, amount := CheckIsAllClusterHaveOne(board); ok {
			countOfAloneClusters = amount
			isEveryClusteHasOne = true
			continue
		}

		size, ball, mask, err := FindLargestCluster(board)
		if err != nil {
			return 0, 0, 0, err
		}
		RemoveCluster(mask, board)
		amountOfMoves++
		points := AddPoints(size)
		scope += points
		fmt.Print(moveMessage(ball, size, amountOfMoves, points))

		if CheckIsAllBallsRemoved(board) {
			isGiveBonus = true
			isAllBallsRemoved = true
			continue
		}

		CompressBoard(board, mask)
	}
	if isGiveBonus {
		scope += 1000
	}

	fmt.Printf("Final scope: %d , with %d balls remaining.\n", scope, countOfAloneClusters)
	fmt.Println()
	return scope, amountOfMoves, countOfAloneClusters, nil
}

func StartTask2(input *os.File) error {
	rd := bufio.NewReader(input)
	err := Task2(rd)
	if err != nil {
		return err
	}
	return nil
}

func Task2(rd *bufio.Reader) error {
	countOfGames, boards, err := ReadBoards(rd)
	if err != nil {
		return err
	}
	BindBoard(boards)

	for i := 0; i < countOfGames; i++ {
		fmt.Printf("Game : %d\n", i+1)
		_, _, _, err = PlayGame(boards[i])
		if err != nil {
			return err
		}
	}
	return nil
}
