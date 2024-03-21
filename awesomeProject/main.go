package main

import (
	"awesomeProject/Functionality"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var tasks Functionality.Tasks

func main() {
	tasks.Tasks = make([]Functionality.Task, 0)
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
