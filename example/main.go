package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/Alexseij/tasks"
)

func spinner(delay time.Duration) {
	for {
		for _, v := range `-\|/` {
			fmt.Printf("\r%c", v)
			time.Sleep(delay)
		}
	}
}

var inputFile = flag.String("input", "input.txt", "file using for input data.")
var outputFile = flag.String("output", "output.txt", "file using for output data.")

func main() {

	go spinner(100 * time.Millisecond)

	file, err := os.Open(*inputFile)
	if err != nil {
		panic("write file")
	}
	defer file.Close()

	output, err := os.OpenFile(*outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("output file")
	}
	defer output.Close()

	tasks.StartTask1(file, output)

}
