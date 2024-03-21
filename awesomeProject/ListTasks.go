package main

import "fmt"

func (x *Tasks) ListTasks() {
	if len(x.tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	fmt.Println("Tasks:")
	for _, task := range x.tasks {
		status := "Complete"
		if !task.Complete {
			status = "Incomplete"
		}
		fmt.Println(task.ID, task.Name, status)
	}
}
