package Functionality

import (
	"fmt"
)

func (x *Tasks) DeleteTask(id int) {
	for i := range x.Tasks {
		if x.Tasks[i].ID == id {
			x.Tasks = append(x.Tasks[:i], x.Tasks[i+1:]...)
			fmt.Println("Task deleted successfully.")
			return
		}
	}
	fmt.Println("Task not found.")
}
