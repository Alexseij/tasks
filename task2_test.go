package tasks_test

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Alexseij/tasks"
)

func deleteFile() {
	time.Sleep(time.Second * 10)
	os.Remove("test.txt")
}

func writeDataToTestFile(data [][]string) (*os.File, error) {
	file, err := os.OpenFile("test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	for i := 0; i < tasks.HEIGHT_OF_GAME_PLACE; i++ {
		for j := 0; j < tasks.WIDTH_OF_GAME_PLACE; j++ {
			var line string
			line = fmt.Sprintf("%s ", data[i][j])
			if j != tasks.WIDTH_OF_GAME_PLACE-1 {
				line = fmt.Sprint("%s\n", data[i][j])
			}
			file.Write([]byte(line))
		}
		file.Write([]byte("\n"))
	}
	go deleteFile()
	return file, nil
}

func TestTask2SampleInputFirst() {
	test := [][]string{
		{"R", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "B"},
		{"G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G"},
		{"G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G"},
		{"G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G"},
		{"G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G"},
		{"G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G"},
		{"G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G"},
		{"G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G"},
		{"G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G"},
		{"B", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "G", "R"},
	}

	want := [][]string{
		{"n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n"},
		{"n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n"},
		{"n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n"},
		{"n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n"},
		{"n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n"},
		{"n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n"},
		{"n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n"},
		{"n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n"},
		{"R", "B", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n"},
		{"B", "R", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n", "n"},
	}
	file, err := writeDataToTestFile(test)
	if err != nil {
		log.Fatal(err)
	}

	rd := bufio.NewReader(file)

}
