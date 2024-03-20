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

var tasks []Task
var lastID int

func AddTask(Name string) {
	lastID++
	task := Task{lastID, Name, false}
	tasks = append(tasks, task)
	fmt.Println("Task added successfully.")
}

func ListTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	fmt.Println("Tasks:")
	for _, task := range tasks {
		status := "Complete"
		if !task.Complete {
			status = "Incomplete"
		}
		fmt.Println(task.ID, task.Name, status)
	}
}

func UpdateTask(id int) {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Complete = !tasks[i].Complete
			fmt.Println("Task updated successfully.")
			return
		}
	}
}

func DeleteTask(id int) {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("Task deleted successfully.")
			return
		}
	}
	fmt.Println("Task not found.")
}

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
			AddTask(taskName)
		case "list":
			ListTasks()
		case "update":
			fmt.Println("Enter task ID to update: ")
			scanner.Scan()
			taskID := 0
			fmt.Scan(scanner.Text(), &taskID)
			UpdateTask(taskID)
		case "delete":
			fmt.Println("Enter task ID to delete: ")
			scanner.Scan()
			taskID := 0
			fmt.Scan(scanner.Text(), &taskID)
			DeleteTask(taskID)
		case "exit":
			fmt.Println("Exiting...")
			break
		default:
			fmt.Println("Invalid command. Please try again.")
		}

	}
}
