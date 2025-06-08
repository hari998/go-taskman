package root

import (
	"fmt"
	"strings"
	"taskman/internal/task"

	"github.com/spf13/cobra"
)


var addCmd = &cobra.Command{
    Use:   "add [task name]",
    Short: "Add a new task",
    Args:  cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        name := strings.Join(args, " ")
        t := task.Add(name)
        fmt.Printf("Task added: %s (ID: %d)\n", t.Name, t.ID)
    },
}

func init() {
	rootCmd.AddCommand(addCmd)
}

