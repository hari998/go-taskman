package task

import (
	"encoding/json"
	"taskman/internal/notify"
	"fmt"
	"os"
	"path/filepath"
)

// import "fmt"

type Task struct {
	ID   int
	Name string
}

var tasks []Task
var nextID = 1
var dataFile string


func Add(name string) Task {
	t := Task{
		ID:   nextID,
		Name: name,
	}

	tasks = append(tasks, t)
	nextID++
	Save()
	notify.Show("Task Added", t.Name)
	return t
}

func GetAll() []Task {
	return tasks
}

func Clear() {
	tasks = nil
	nextID = 1
}


func init() {

	home, err := os.UserHomeDir()

	if err != nil {
		panic(fmt.Sprintf("Error getting home directory: %v", err))
	}

	// fmt.Println("Home Directory:", home)

	configDir := filepath.Join(home, ".taskman")
	os.MkdirAll(configDir, 0755)

	dataFile = filepath.Join(configDir, "tasks.json")
	// fmt.Println("Data File Path:", dataFile)

	Load()
}


func Save() {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(dataFile, data, 0644)
	if err != nil {
		panic(err)
	}
}

func Load() {
	data, err := os.ReadFile(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return // File does not exist, nothing to load
		}
		panic(err)
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		panic(err)
	}

	nextID = 1

	for _, t := range tasks {
		if t.ID >= nextID {
			nextID = t.ID + 1
		}
	}

}