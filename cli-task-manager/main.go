package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"encoding/json"
)

const tasksFileName = "tasks.json"

type Task struct {
	ID         int
	Name       string
	Completed  bool
	Created_at string
	Categories []string
}

func loadTasks() ([]Task, error) {
	var tasks = []Task{}
	file, err := os.ReadFile(tasksFileName)
	if err != nil {
		if os.IsNotExist(err) {
			return tasks, nil
		}
		return nil, err
	}

	json.Unmarshal(file, &tasks)

	return tasks, err
}

func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "\t")
	if err != nil {
		return err
	}

	return os.WriteFile(tasksFileName, data, 0644)
}

func createTask(name string, categories []string) error {
	tasks, err := loadTasks()
	fmt.Println(tasks)
	if err != nil {
		return err
	}

	categoryMap := make(map[string]bool)

	var uniqueCategories []string
	for _, cat := range categories {
		lowerCat := strings.ToLower(cat)
		if _, ok := categoryMap[lowerCat]; !ok {
			categoryMap[lowerCat] = true
			uniqueCategories = append(uniqueCategories, lowerCat)
		}
	}

	id := len(tasks) + 1
	task := Task{
		ID:         id,
		Name:       name,
		Completed:  false,
		Created_at: time.Now().Format("2006/1/2 15:04"),
		Categories: categories,
	}

	tasks = append(tasks, task)

	return saveTasks(tasks)
}

func listTasks(category string) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	category = strings.ToLower(category)
	filteredTasks := []Task{}

	if category == "" {
		filteredTasks = tasks
	} else {
		for _, task := range tasks {
			for _, cat := range task.Categories {
				if cat == category {
					filteredTasks = append(filteredTasks, task)
					break
				}
			}
		}
	}

	for _, task := range filteredTasks {

		status := " "
		if task.Completed {
			status = "X"
		}
		fmt.Printf("%d. [%s] %s\n", task.ID, status, task.Name)
	}

	return nil
}

func completeTask(id int) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Completed = true
			break
		}
	}

	return saveTasks(tasks)
}

func deleteTask(id int) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}

	return saveTasks(tasks)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [ create | list | delete | complete ] [task_id/task_name]")
		return
	}

	command := os.Args[1]

	switch command {
	case "create":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go create [task_name]")
			return
		}
		name := os.Args[2]
		categories := os.Args[3:]
		if err := createTask(name, categories); err != nil {
			fmt.Println("Error creating task:", err)
			return
		}
	case "list":
		if len(os.Args) > 3 {
			fmt.Println("Usage: go run main.go list [category]")
			return
		}
		var category string
		if len(os.Args) > 2 {
			category = os.Args[2]
		}
		if err := listTasks(category); err != nil {
			fmt.Println("Error listing tasks:", err)
			return
		}
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go delete [task_id]")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task id")
			return
		}
		if err := deleteTask(id); err != nil {
			fmt.Println("Error deleting task:", err)
			return
		}
	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go complete [task_id]")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task id")
			return
		}
		if err := completeTask(id); err != nil {
			fmt.Println("Error completing task:", err)
			return
		}
	default:
		fmt.Println("Invalid command")
	}
}
