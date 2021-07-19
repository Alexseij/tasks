package tasks_test

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/Alexseij/tasks"
)

var (
	flagAmountOfIterations = flag.Int("iterations", 100, "This flag using for amount of iteration in testing cycle for each generating matrix.")
	flagMaxLength          = flag.Int("length", 10, "This flag using for max length of matrix , when it prepare for generating.")
	flagMaxWidth           = flag.Int("width", 10, "This flag using for max width of matrix , when it prepare for generating.")
)

var (
	patternMatrix = [][]string{
		{"1", "1", "0", "0", "0"},
		{"0", "1", "1", "0", "0"},
		{"1", "0", "0", "1", "0"},
		{"1", "1", "1", "1", "0"},
		{"0", "0", "1", "1", "1"},
	}
)

// Function using for comparing two matrices
// Returns : true - if they equals , else false
func compareMatrix(matrixA, matrixB [][]string) bool {
	for i, row := range matrixA {
		for j := range row {
			if matrixA[i][j] != matrixB[i][j] {
				return false
			}
		}
	}
	return true
}

// Function using for generating random matrix in different tests
// Returns : new generated matrix , length and width of that matrix
func generateMatrix(minLength, minWidth, maxLenght, maxWidth int) (l int, w int, matrix [][]string) {

	if maxWidth < 0 || maxLenght < 0 {
		panic("Uncorrect falgs data , they must be positive.")
	}

	//Seed uses the provided seed value to initialize the default Source to a deterministic state
	//https://pkg.go.dev/math/rand#Rand.Seed
	rand.Seed(time.Now().UnixNano())
	//Generate random length and width
	//https://stackoverflow.com/questions/1202687/how-do-i-get-a-specific-range-of-numbers-from-rand
	length := rand.Int()%(maxLenght+1-minLength) + minLength
	width := rand.Int()%(maxWidth+1-minWidth) + minWidth

	//Can you imagine matrix 0 x 100 or 100 x 0 of 0 x 0 ? I can't.
	if length == 0 || width == 0 {
		return length, width, nil
	}
	result := make([][]string, length)

	for i := 0; i < length; i++ {
		row := make([]string, width)
		for j := 0; j < width; j++ {
			//using rand.Intn(2) cuz matrix can have only "1" or "0" values n'
			// rand.Intn(2) generates value in [0 , 2)
			row[j] = strconv.Itoa(rand.Intn(2))
		}
		result[i] = row
	}

	return length, width, result
}

//Function using for copying sample matrix into testMatrix
func bindMatrix(sample [][]string, length, width int) [][]string {
	testMatrix := make([][]string, length)
	for i := range sample {
		testMatrix[i] = make([]string, width)
		copy(testMatrix[i], sample[i])
	}
	return testMatrix
}

//Function using for make error messages if test didn't pass
//Returns : string message whitch using in console output
func makeMsgWithMatrix(matrix [][]string, header string) string {
	msg := fmt.Sprintf("%s \n", header)
	for i, v := range matrix {
		for j := range v {
			msg = fmt.Sprintf("%s%s ", msg, matrix[i][j])
		}
		msg = fmt.Sprintf("%s\n", msg)
	}
	msg = fmt.Sprintf("\n%s", msg)
	return msg
}

//Function prints fatal msg if test didn't pass
func printFatalMsg(t *testing.T, testMatrix, want, matrixA [][]string) {
	msgTestMatrix := makeMsgWithMatrix(testMatrix, "Result matrix")
	msgWantMatrix := makeMsgWithMatrix(want, "Want matrix")
	msgMatrixA := makeMsgWithMatrix(matrixA, "Matrix A")
	t.Fatalf("%s%s%s", msgTestMatrix, msgWantMatrix, msgMatrixA)
}

//Test checks empty matrix situation in input file
func TestTask1Empty(t *testing.T) {
	matrixA := [][]string{}
	lengthA := 0
	widthA := 0

	testMatrix := [][]string{
		{"1", "1", "0", "0", "0"},
		{"1", "1", "1", "0", "0"},
		{"1", "0", "0", "1", "0"},
		{"1", "1", "1", "0", "0"},
		{"0", "0", "1", "1", "1"},
	}

	want := [][]string{
		{"1", "1", "0", "0", "0"},
		{"1", "1", "1", "0", "0"},
		{"1", "0", "0", "1", "0"},
		{"1", "1", "1", "0", "0"},
		{"0", "0", "1", "1", "1"},
	}

	tasks.Task1(matrixA, testMatrix, lengthA, widthA, 5, 5)

	if !compareMatrix(want, testMatrix) {
		printFatalMsg(t, testMatrix, want, matrixA)
	}

}

// This test checks situations where matrix A have size bigger then matrix B on width and length
func TestTask1BigEnoguhtSize(t *testing.T) {
	amountOfIteration := *flagAmountOfIterations
	for amountOfIteration != 0 {
		//Test matrix it's B matrix in task
		lengthTest, widthTest, testMatrix := generateMatrix(0, 0, *flagMaxLength, *flagMaxWidth)

		lengthA, widthA, matrixA := generateMatrix(lengthTest+1, widthTest+1, *flagMaxLength+1, *flagMaxWidth+1)

		want := bindMatrix(testMatrix, lengthTest, widthTest)

		tasks.Task1(matrixA, testMatrix, lengthA, widthA, lengthTest, widthTest)

		if !compareMatrix(want, testMatrix) {
			printFatalMsg(t, testMatrix, want, matrixA)
		}
		amountOfIteration--
	}

}

//This test checks situations where matrix A have size bigger then matrix B only on width
func TestTask1BigEnoguhtSizeWidth(t *testing.T) {
	amountOfIteration := *flagAmountOfIterations

	for amountOfIteration != 0 {
		//Test matrix it's B matrix in task
		lengthTest, widthTest, testMatrix := generateMatrix(0, 0, *flagMaxLength, *flagMaxWidth)

		lengthA, widthA, matrixA := generateMatrix(0, widthTest+1, *flagMaxLength, *flagMaxWidth+1)

		want := bindMatrix(testMatrix, lengthTest, widthTest)

		tasks.Task1(matrixA, testMatrix, lengthA, widthA, lengthTest, widthTest)

		if !compareMatrix(want, testMatrix) {
			printFatalMsg(t, testMatrix, want, matrixA)
		}
		amountOfIteration--
	}
}

//This test checks situations where matrix A have size bigger then matrix B only on length
func TestTask1BigEnoguhtSizeLength(t *testing.T) {

	amountOfIteration := *flagAmountOfIterations

	for amountOfIteration != 0 {
		//Test matrix it's B matrix in task
		lengthTest, widthTest, testMatrix := generateMatrix(0, 0, *flagMaxLength, *flagMaxWidth)

		lengthA, widthA, matrixA := generateMatrix(lengthTest+1, 0, *flagMaxLength+1, *flagMaxWidth)

		want := bindMatrix(testMatrix, lengthTest, widthTest)

		tasks.Task1(matrixA, testMatrix, lengthA, widthA, lengthTest, widthTest)

		if !compareMatrix(want, testMatrix) {
			printFatalMsg(t, testMatrix, want, matrixA)
		}
		amountOfIteration--
	}
}

// <-- This tests checks input data from task -- > //

func TestTask1OnMatrixFromSampleInputFirst(t *testing.T) {
	matrixA := [][]string{
		{"1", "0"},
		{"1", "1"},
	}

	testMatrix := bindMatrix(patternMatrix, 5, 5)

	want := [][]string{
		{"1", "2", "*", "0", "0"},
		{"0", "2", "2", "0", "0"},
		{"2", "*", "0", "1", "0"},
		{"2", "2", "1", "2", "*"},
		{"0", "0", "1", "2", "2"},
	}

	tasks.Task1(matrixA, testMatrix, 2, 2, 5, 5)

	if !compareMatrix(want, testMatrix) {
		printFatalMsg(t, testMatrix, want, matrixA)
	}
}

func TestTask1OnMatrixFromSampleInputSecond(t *testing.T) {
	matrixA := [][]string{
		{"1"},
	}

	testMatrix := bindMatrix(patternMatrix, 5, 5)

	want := [][]string{
		{"2", "2", "0", "0", "0"},
		{"0", "2", "2", "0", "0"},
		{"2", "0", "0", "2", "0"},
		{"2", "2", "2", "2", "0"},
		{"0", "0", "2", "2", "2"},
	}

	tasks.Task1(matrixA, testMatrix, 1, 1, 5, 5)

	if !compareMatrix(want, testMatrix) {
		printFatalMsg(t, testMatrix, want, matrixA)
	}

}
func TestTask1OnMatrixFromSampleInputThird(t *testing.T) {
	matrixA := [][]string{
		{"0"},
	}

	testMatrix := bindMatrix(patternMatrix, 5, 5)

	want := [][]string{
		{"1", "1", "*", "*", "*"},
		{"*", "1", "1", "*", "*"},
		{"1", "*", "*", "1", "*"},
		{"1", "1", "1", "1", "*"},
		{"*", "*", "1", "1", "1"},
	}

	tasks.Task1(matrixA, testMatrix, 1, 1, 5, 5)

	if !compareMatrix(want, testMatrix) {
		printFatalMsg(t, testMatrix, want, matrixA)
	}

}
