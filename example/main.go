package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Alexseij/tasks"
)

func spinner() {
	for {
		for _, v := range `-\|/` {
			fmt.Printf("\r%c", v)
			time.Sleep(time.Millisecond * 200)
		}
	}
}

func main() {

	go spinner()

	file, err := os.Open("input.txt")
	if err != nil {
		panic("write file")
	}
	defer file.Close()

	output, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("output file")
	}

	tasks.StartTask3(file, output)

}
