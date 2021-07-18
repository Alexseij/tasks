package tasks

import (
	"bufio"
	"log"
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

func updateCheckMatrix(matrix [][]bool, length, width int) {
	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			matrix[i][j] = false
		}
	}
}

func check(checkMatrix [][]bool, matrixA, matrixB [][]string, fromIndexLength, toIndexLength, fromIndexWidth, toIndexWidth int) {
	indexLengthA, indexWidthA := 0, 0
	//Invariant teels everithing is ok :-)
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

func readLengthAndWidth(rd *bufio.Reader) (l int, w int) {
	line, _ := rd.ReadString('\n')

	splitLine := strings.Split(strings.TrimSpace(line), " ")

	length, err := strconv.Atoi(splitLine[0])
	if err != nil {
		log.Fatal(err)
	}
	width, err := strconv.Atoi(splitLine[1])
	if err != nil {
		log.Fatal(err)
	}

	return length, width
}

func Task1(matrixA, matrixB [][]string, lengthA, widthA, lengthB, widthB int) {

	if lengthA == 0 || widthA == 0 {
		return
	}

	checkMatrix := make([][]bool, lengthA)

	for i := 0; i < lengthA; i++ {
		row := make([]bool, widthA)
		for j := 0; j < widthA; j++ {
			row[j] = false
		}
		checkMatrix[i] = row
	}

	for i := 0; i < lengthB-(lengthB%lengthA); i++ {
		for j := 0; j < widthB-(widthB%widthA); j++ {
			if matrixA[0][0] == matrixB[i][j] {
				check(checkMatrix, matrixA, matrixB, i, i+lengthA-1, j, j+widthA-1)
				flag := allIsTrue(checkMatrix, lengthA, widthA)
				updateCheckMatrix(checkMatrix, lengthA, widthA)
				if flag {
					changeMatrix(matrixB, i, i+lengthA, j, j+widthA)
				}
			}
		}
	}
}

func writeDataForTask1(matrixB [][]string, output *os.File, lengthB, widthB int) {
	for i := 0; i < lengthB; i++ {
		for j := 0; j < widthB; j++ {
			output.Write([]byte(matrixB[i][j]))
			output.Write([]byte(" "))
		}
		output.Write([]byte("\n"))
	}

}

func StartTask1(input, output *os.File) {

	rd := bufio.NewReader(input)

	lengthA, widthA := readLengthAndWidth(rd)
	matrixA, err := createMatrix(lengthA, widthA, rd)
	if err != nil {
		log.Fatal(err)
	}

	lengthB, widthB := readLengthAndWidth(rd)
	matrixB, err := createMatrix(lengthB, widthB, rd)
	if err != nil {
		log.Fatal(err)
	}

	Task1(matrixA, matrixB, lengthA, widthA, lengthB, widthB)

	writeDataForTask1(matrixB, output, lengthB, widthB)
}
