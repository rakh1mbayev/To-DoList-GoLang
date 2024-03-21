package Functionality

import (
	"fmt"
)

func (x *Tasks) DeleteTask(id int) {
	for i := range x.tasks {
		if x.tasks[i].ID == id {
			x.tasks = append(x.tasks[:i], x.tasks[i+1:]...)
			fmt.Println("Task deleted successfully.")
			return
		}
	}
	fmt.Println("Task not found.")
}
