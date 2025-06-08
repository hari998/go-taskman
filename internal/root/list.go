package root

import (
	"fmt"
	"taskman/internal/task"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := task.GetAll()
		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}
		fmt.Println("Listing all tasks:")
		for _, t := range tasks {
			fmt.Printf("ID: %d Task: %s\n", t.ID, t.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
