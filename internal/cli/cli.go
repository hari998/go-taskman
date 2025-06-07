package cli

import (
	"fmt"
	"strings"
	"taskman/internal/task"
	// "taskman/pkg/logger"
)

func Run(args []string) {
	// fmt.Println("Task Manager is running with arguments:", args)
	// logger.Info("Logger >> " + strings.Join(args, " "))

	if len(args) < 1 {
		fmt.Println("No command provided.\nUsage: taskman [add|list]")
		return
	}

	command := args[0]

	switch command {
	case "add":
		if len(args) < 2 {
			fmt.Println("No task name provided.\nUsage: taskman add <task_name>")
			return
		}
		name := strings.Join(args[1:], " ")
		t := task.Add(name)

		fmt.Printf("Task added: %s (ID: %d)\n", t.Name, t.ID)

	case "list":
		tasks := task.GetAll()
		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}
		fmt.Println("Listing all tasks:")
		for _, t := range tasks {
			fmt.Println("ID: ", t.ID, "Name: ", t.Name)
		}

	default:
		fmt.Println("Unknown command:", command)
		fmt.Println("Available commands: add, list. Usage: taskman [add|list]")

	}

}
