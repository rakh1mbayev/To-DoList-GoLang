package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	ID       int
	Name     string
	Complete bool
}

type Tasks struct {
	tasks []Task
}

var tasks Tasks
var lastID int

func main() {
	fmt.Println("Welcome to To-Do List App!")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter command (add/list/update/delete/exit): ")
		scanner.Scan()
		command := strings.ToLower(scanner.Text())
		switch command {
		case "add":
			fmt.Println("Enter task name: ")
			scanner.Scan()
			taskName := scanner.Text()
			tasks.AddTask(taskName)
		case "list":
			tasks.ListTasks()
		case "update":
			fmt.Println("Enter task ID to update: ")
			scanner.Scan()
			taskID := 0
			fmt.Scan(scanner.Text(), &taskID)
			tasks.UpdateTask(taskID)
		case "delete":
			fmt.Println("Enter task ID to delete: ")
			scanner.Scan()
			taskID := 0
			fmt.Scan(scanner.Text(), &taskID)
			tasks.DeleteTask(taskID)
		case "exit":
			fmt.Println("Exiting...")
			break
		default:
			fmt.Println("Invalid command. Please try again.")
		}

	}
}
