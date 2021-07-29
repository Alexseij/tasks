package main

import (
	"flag"
	"log"
	"os"

	"github.com/Alexseij/tasks"
)

var inputFile = flag.String("input", "input.txt", "file using for input data.")
var outputFile = flag.String("output", "output.txt", "file using for output data.")
var task = flag.String("task", "", "flag using for choose current task execution")

func main() {

	flag.Parse()

	if *task == "" {
		panic("No task")
	}

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

	switch *task {
	case "task1":
		err := tasks.StartTask1(file, output)
		if err != nil {
			log.Fatal(err)
		}
		break
	case "task2":
		err := tasks.StartTask2(file)
		if err != nil {
			log.Fatal(err)
		}
		break
	}
}
