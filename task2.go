package tasks

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func readBlocks(rd *bufio.Reader) (map[int]map[int][]int, error) {
	line, err := rd.ReadString('\n')

	if err != nil {
		return nil, err
	}

	countOfTrains := 0
	trains := make(map[int]map[int][]int)

	for strings.TrimSpace(line) != "0" {
		countOfCoaches, err := strconv.Atoi(strings.TrimSpace(line))

		if err != nil {
			return nil, err
		}

		trainLine, err := rd.ReadString('\n')
		if err != nil {
			return nil, err
		}

		countOfRoads := 0
		roads := make(map[int][]int)

		for strings.TrimSpace(trainLine) != "0" {
			coaches := make([]int, countOfCoaches)
			splitStr := strings.Split(strings.TrimSpace(trainLine), " ")

			for i := 0; i < countOfCoaches; i++ {
				coaches[i], err = strconv.Atoi(splitStr[i])
				if err != nil {
					return nil, err
				}
			}

			roads[countOfRoads] = coaches
			countOfRoads++
			trainLine, err = rd.ReadString('\n')
			if err != nil {
				return nil, err
			}
		}
		trains[countOfTrains] = roads
		countOfTrains++
		line, err = rd.ReadString('\n')
		if err != nil && err != io.EOF {
			return nil, err
		}
	}
	return trains, nil
}

func makeRoadForCompare(length int) []int {
	road := make([]int, length)
	for i := 0; i < length; i++ {
		road[i] = i + 1
	}
	return road
}

func compareCoaches(roadForCompare, road []int) bool {
	for i, coach := range road {
		if roadForCompare[i] == coach {
			return false
		}
	}
	return true
}

func findNeedPermutations(trains map[int]map[int][]int) map[int]map[int]string {
	result := make(map[int]map[int]string)

	for numTrain, roads := range trains {
		// We can get first road for pattern length
		roadForCompare := makeRoadForCompare(len(roads[0]))
		mapRoad := make(map[int]string)
		for numRoad, road := range roads {
			if compareCoaches(roadForCompare, road) {
				mapRoad[numRoad] = "YES"
			} else {
				mapRoad[numRoad] = "NO"
			}
		}
		result[numTrain] = mapRoad
	}

	return result
}

func writeDataForTask3(output *os.File, result map[int]map[int]string) (int, error) {
	amountOfBytes := 0

	for _, roads := range result {
		for _, answer := range roads {
			msg := fmt.Sprintf("%s\n", answer)

			writedBytes, err := output.Write([]byte(msg))
			amountOfBytes += writedBytes

			if err != nil {
				return amountOfBytes, err
			}

		}
		writedBytes, err := output.Write([]byte("\n"))
		amountOfBytes += writedBytes
		if err != nil {
			return amountOfBytes, err
		}
	}

	return amountOfBytes, nil
}

func StartTask3(input, output *os.File) {
	rd := bufio.NewReader(input)

	trains, err := readBlocks(rd)
	if err != nil {
		panic("read block")
	}

	result := findNeedPermutations(trains)

	_, err = writeDataForTask3(output, result)
	if err != nil {
		panic("write date for task")
	}
}
