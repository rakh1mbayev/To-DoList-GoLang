package Functionality

import "fmt"

type Task struct {
	ID       int
	Name     string
	Complete bool
}
type Tasks struct {
	tasks []Task
}

var lastID int

func (x *Tasks) AddTask(Name string) {
	lastID++
	task := Task{lastID, Name, false}
	x.tasks = append(x.tasks, task)
	fmt.Println("Task added successfully.")

}
