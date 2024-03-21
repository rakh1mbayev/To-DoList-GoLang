package main

import "fmt"

func (x *Tasks) AddTask(Name string) {
	lastID++
	task := Task{lastID, Name, false}
	x.tasks = append(x.tasks, task)
	fmt.Println("Task added successfully.")
}
