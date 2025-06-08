package main

import "taskman/internal/root"

// "fmt"
// "os"
// "taskman/internal/cli"

func main() {
	// cli.Run(os.Args[1:])
	root.Execute()
	// fmt.Println("Task Manager has finished running.")
}
