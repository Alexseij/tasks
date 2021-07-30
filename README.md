# How does it use ?

Use command : `go run main.go -task="TASK_NAME"`.

That command has three flags :

`-input` - path to input file (by default example/input.txt).
`-output` - path to output file (by default example/output.txt).
`-task` - which task is using.

# Values for -task

1) task1 - DigitalLab
2) task2 - RGBGame

# How test tasks

Use command : `go test -v`

Note : If you want to change `task2_test.txt` file after every board should written `\n` byte.
