package tasks

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// NOTE : I deliberately miss errors checking in this code for improve my speed of coding :-)

//<-- This section of code has responsibility for First Task -->//

const (
	ONE_INDEX    = "1"
	SECOND_INDEX = "2"
	STAR_INDEX   = "*"
	ZERO_INDEX   = "0"
)

//Function using for read matrix data from input file
//Returns : matrix with size length x width , or error
func createMatrix(length, width int, rd *bufio.Reader) ([][]string, error) {

	if length == 0 || width == 0 {
		rd.ReadString('\n')
		return nil, nil
	}
	result := make([][]string, length)

	for i := 0; i < length; i++ {
		line, err := rd.ReadString('\n')
		if err != nil {
			return nil, err
		}

		rowStr := strings.Split(strings.TrimSpace(line), " ")
		result[i] = rowStr
	}

	return result, nil
}

//Function using for check all true values in param matrix
//Returns : if all values is true - true , else false
func allIsTrue(matrix [][]bool, length, width int) bool {
	isAllTrue := true
	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			if !matrix[i][j] {
				isAllTrue = false
				break
			}
		}
	}
	return isAllTrue
}

//Fuction using for clear data after using in matrix
func updateCheckMatrix(matrix [][]bool, length, width int) {
	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			matrix[i][j] = false
		}
	}
}

//Fucntion check equals between matrixA and matrixB from given indexes
//all results of comparing writes in matrix with size lengthA , widthA and type [][]bool
func check(checkMatrix [][]bool, matrixA, matrixB [][]string, fromIndexLength, toIndexLength, fromIndexWidth, toIndexWidth int) {
	indexLengthA, indexWidthA := 0, 0
	for i := fromIndexLength; i <= toIndexLength; i++ {
		for j := fromIndexWidth; j <= toIndexWidth; j++ {
			if matrixA[indexLengthA][indexWidthA] == matrixB[i][j] {
				checkMatrix[indexLengthA][indexWidthA] = true
			}
			indexWidthA++
		}
		indexWidthA = 0
		indexLengthA++
	}
}

//Fuction for changing matrix B values
func changeMatrix(matrix [][]string, fromIndexLength, toIndexLength, fromIndexWidth, toIndexWidth int) {
	for i := fromIndexLength; i < toIndexLength; i++ {
		for j := fromIndexWidth; j < toIndexWidth; j++ {
			if matrix[i][j] == ONE_INDEX {
				matrix[i][j] = SECOND_INDEX
			} else if matrix[i][j] == ZERO_INDEX {
				matrix[i][j] = STAR_INDEX
			}
		}
	}
}

//Function using for reading legnth and width form current matrix from input file
func readLengthAndWidth(rd *bufio.Reader) (l int, w int, err error) {
	line, err := rd.ReadString('\n')
	if err != nil && err != io.EOF {
		return 0, 0, err
	}
	splitLine := strings.Split(strings.TrimSpace(line), " ")

	length, err := strconv.Atoi(splitLine[0])
	if err != nil {
		return 0, 0, err
	}
	width, err := strconv.Atoi(splitLine[1])
	if err != nil {
		return 0, 0, err
	}

	return length, width, nil
}

// <-----------------Alghorithm------------------> //
// We have two matrices, this is matrix A,		   //
// the occurrences of which we will look		   //
// for in matrix B, for this we will simply		   //
// compare all values from the matrixA in          //
// matrix B and write the results of comparing     //
// into the matrix that was initialized above,     //
// namely the true - says that the values at the   //
// current index in matrix a coincided with the    //
// value in matrix B otherwise false               //

func Task1(matrixA, matrixB [][]string, lengthA, widthA, lengthB, widthB int) {
	// Pointless matrix
	if lengthA == 0 || widthA == 0 {
		return
	}
	// Matrix storing matches of matrices A in matrix B
	checkMatrix := make([][]bool, lengthA)

	for i := 0; i < lengthA; i++ {
		row := make([]bool, widthA)
		for j := 0; j < widthA; j++ {
			row[j] = false
		}
		checkMatrix[i] = row
	}
	// lengthB-(lengthB%lengthA) and widthB-(widthB%widthA) using for decreesse interations in cycle
	for i := 0; i < lengthB-(lengthB%lengthA); i++ {
		for j := 0; j < widthB-(widthB%widthA); j++ {
			if matrixA[0][0] == matrixB[i][j] {
				check(checkMatrix, matrixA, matrixB, i, i+lengthA-1, j, j+widthA-1)
				//If matrixA and matrixB equals in current index flag have value true
				flag := allIsTrue(checkMatrix, lengthA, widthA)
				updateCheckMatrix(checkMatrix, lengthA, widthA)
				if flag {
					changeMatrix(matrixB, i, i+lengthA, j, j+widthA)
				}
			}
		}
	}
}

//Function using for writting data into output file
//Returns : number of bytes written in file and error if it was
func writeDataForTask1(matrixB [][]string, output *os.File, lengthB, widthB int) (int, error) {
	amountOfBytes := 0

	for i := 0; i < lengthB; i++ {
		for j := 0; j < widthB; j++ {
			value := fmt.Sprintf("%s ", matrixB[i][j])
			bytes, err := output.Write([]byte(value))
			amountOfBytes += bytes
			if err != nil {
				return amountOfBytes, err
			}
		}
		bytes, err := output.Write([]byte("\n"))
		amountOfBytes += bytes
		if err != nil && err != io.EOF {
			return amountOfBytes, err
		}
	}
	return amountOfBytes, nil
}

//Function using for start solving task
//Returns : error if have trouble with I/O
func StartTask1(input, output *os.File) error {

	rd := bufio.NewReader(input)

	lengthA, widthA, err := readLengthAndWidth(rd)
	if err != nil {
		return err
	}

	matrixA, err := createMatrix(lengthA, widthA, rd)
	if err != nil {
		return err
	}

	lengthB, widthB, err := readLengthAndWidth(rd)
	if err != nil {
		return err
	}

	matrixB, err := createMatrix(lengthB, widthB, rd)
	if err != nil {
		return err
	}

	Task1(matrixA, matrixB, lengthA, widthA, lengthB, widthB)

	_, err = writeDataForTask1(matrixB, output, lengthB, widthB)
	if err != nil {
		return err
	}

	return nil
}
