package tasks_test

import (
	"bufio"
	"log"
	"os"
	"testing"

	"github.com/Alexseij/tasks"
)

var TestBoards map[int]tasks.Board

func init() {
	TestBoards = ReadBoardsForTest()
	tasks.BindBoard(TestBoards)
}

func ReadBoardsForTest() map[int]tasks.Board {
	file, err := os.Open("test_files/task2_test.txt")
	if err != nil {
		log.Print("File is empty")
		os.Exit(-1)
	}
	defer file.Close()

	rd := bufio.NewReader(file)

	_, testBoards, err := tasks.ReadBoards(rd)
	if err != nil {
		log.Fatal(err)
	}

	return testBoards
}

func PlayTestGame(
	t *testing.T, board tasks.Board,
	wantMoves map[int]map[string]interface{},
	wantFinalScope int,
	wantBallsRemaining int,
) (err error) {
	scope := 0
	amountOfMoves := 0
	countOfAloneClusters := 0

	wantMovesLen := len(wantMoves)

	isAllBallsRemoved := false
	isEveryClusteHasOne := false
	isGiveBonus := false

	for !isAllBallsRemoved && !isEveryClusteHasOne {

		if ok, amount := tasks.CheckIsAllClusterHaveOne(board); ok {
			countOfAloneClusters = amount
			isEveryClusteHasOne = true
			continue
		}

		size, ball, mask, err := tasks.FindLargestCluster(board)
		if err != nil {
			return err
		}

		if amountOfMoves >= wantMovesLen {
			t.Fatalf("amountOfMoves = %d , want : %d", amountOfMoves+1, wantMovesLen)
		}

		points := tasks.AddPoints(size)

		wantPoints := wantMoves[amountOfMoves]["add"].(int)
		if wantPoints != points {
			t.Fatalf("Add point = %d , want : %d", points, wantPoints)
		}

		color := ball.Value
		wantColor := wantMoves[amountOfMoves]["color"].(string)
		if color != wantColor {
			t.Fatalf("Current color = %s , want : %s", color, wantColor)
		}

		wantSize := wantMoves[amountOfMoves]["removed"].(int)
		if size != wantSize {
			t.Fatalf("Removed balls from cluster = %d , want : %d", size, wantSize)
		}

		wantCoordX := wantMoves[amountOfMoves]["X"].(int)
		wantCoordY := wantMoves[amountOfMoves]["Y"].(int)
		if wantCoordX != ball.X+1 || wantCoordY != tasks.HEIGHT_OF_GAME_PLACE-ball.Y {
			t.Fatalf(
				"Current coords : (%d , %d) , want : (%d , %d)",
				wantCoordY,
				wantCoordX,
				tasks.HEIGHT_OF_GAME_PLACE-ball.Y,
				ball.X+1,
			)
		}
		amountOfMoves++
		tasks.RemoveCluster(mask, board)
		scope += points
		if tasks.CheckIsAllBallsRemoved(board) {
			isGiveBonus = true
			isAllBallsRemoved = true
			continue
		}

		tasks.CompressBoard(board, mask)
	}
	if isGiveBonus {
		scope += 1000
	}
	if wantFinalScope != scope {
		t.Fatalf("Final scope = %d , want : %d", scope, wantFinalScope)
	}

	if countOfAloneClusters != wantBallsRemaining {
		t.Fatalf("Balls remaining = %d , want : %d", countOfAloneClusters, wantBallsRemaining)
	}

	return nil
}

func TestTask2First(t *testing.T) {

	//First key-map value which move , second info abount move : color of ball , added scope and etc.
	wantMoves := make(map[int]map[string]interface{})

	wantMoves[0] = map[string]interface{}{
		"add":     900,
		"removed": 32,
		"color":   "B",
		"Y":       4,
		"X":       1,
	}

	wantMoves[1] = map[string]interface{}{
		"add":     1369,
		"removed": 39,
		"color":   "R",
		"Y":       2,
		"X":       1,
	}

	wantMoves[2] = map[string]interface{}{
		"add":     1225,
		"removed": 37,
		"color":   "G",
		"Y":       1,
		"X":       1,
	}

	wantMoves[3] = map[string]interface{}{
		"add":     81,
		"removed": 11,
		"color":   "B",
		"Y":       3,
		"X":       4,
	}

	wantMoves[4] = map[string]interface{}{
		"add":     36,
		"removed": 8,
		"color":   "R",
		"Y":       1,
		"X":       1,
	}

	wantMoves[5] = map[string]interface{}{
		"add":     16,
		"removed": 6,
		"color":   "G",
		"Y":       2,
		"X":       1,
	}

	wantMoves[6] = map[string]interface{}{
		"add":     16,
		"removed": 6,
		"color":   "B",
		"Y":       1,
		"X":       6,
	}

	wantMoves[7] = map[string]interface{}{
		"add":     9,
		"removed": 5,
		"color":   "R",
		"Y":       1,
		"X":       2,
	}

	wantMoves[8] = map[string]interface{}{
		"add":     9,
		"removed": 5,
		"color":   "G",
		"Y":       1,
		"X":       2,
	}

	wantFinalScope := 3661
	wantBallsRemaining := 1
	tasks.PrintBoard(TestBoards[0])
	PlayTestGame(t, TestBoards[0], wantMoves, wantFinalScope, wantBallsRemaining)
}

func TestTask2Second(t *testing.T) {
	//First key-map value which move , second info abount move : color of ball , added scope and etc.
	wantMoves := make(map[int]map[string]interface{})

	wantMoves[0] = map[string]interface{}{
		"add":     784,
		"removed": 30,
		"color":   "G",
		"Y":       1,
		"X":       1,
	}

	wantMoves[1] = map[string]interface{}{
		"add":     784,
		"removed": 30,
		"color":   "R",
		"Y":       1,
		"X":       1,
	}

	wantMoves[2] = map[string]interface{}{
		"add":     784,
		"removed": 30,
		"color":   "B",
		"Y":       1,
		"X":       1,
	}

	wantMoves[3] = map[string]interface{}{
		"add":     784,
		"removed": 30,
		"color":   "G",
		"Y":       1,
		"X":       1,
	}

	wantMoves[4] = map[string]interface{}{
		"add":     784,
		"removed": 30,
		"color":   "R",
		"Y":       1,
		"X":       1,
	}

	wantFinalScope := 4920
	wantBallsRemaining := 0
	tasks.PrintBoard(TestBoards[1])
	PlayTestGame(t, TestBoards[1], wantMoves, wantFinalScope, wantBallsRemaining)

}

func TestTask2Third(t *testing.T) {
	var wantMoves map[int]map[string]interface{}

	wantFinalScope := 0
	wantBallsRemaining := 150
	tasks.PrintBoard(TestBoards[2])
	PlayTestGame(t, TestBoards[2], wantMoves, wantFinalScope, wantBallsRemaining)
}

func TestTask2Fourth(t *testing.T) {
	wantMoves := make(map[int]map[string]interface{})

	wantFinalScope := 20736
	wantBallsRemaining := 4

	wantMoves[0] = map[string]interface{}{
		"add":     20736,
		"removed": 146,
		"color":   "G",
		"Y":       2,
		"X":       1,
	}
	tasks.PrintBoard(TestBoards[3])
	PlayTestGame(t, TestBoards[3], wantMoves, wantFinalScope, wantBallsRemaining)
}

func TestTask2Fifth(t *testing.T) {

	wantMoves := make(map[int]map[string]interface{})

	wantFinalScope := 22904
	wantBallsRemaining := 0

	wantMoves[0] = map[string]interface{}{
		"add":     21904,
		"removed": 150,
		"color":   "G",
		"Y":       1,
		"X":       1,
	}
	tasks.PrintBoard(TestBoards[4])
	PlayTestGame(t, TestBoards[4], wantMoves, wantFinalScope, wantBallsRemaining)

}
