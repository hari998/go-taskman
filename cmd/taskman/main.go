package main

import (
	// "fmt"
	"os"
	"taskman/internal/cli"
)

func main() {
	// fmt.Println("Task Manager is running...", os.Args)
	cli.Run(os.Args[1:])
	// fmt.Println("Task Manager has finished running.")
}