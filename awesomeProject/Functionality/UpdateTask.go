package Functionality

import (
	"fmt"
)

func (x *Tasks) UpdateTask(id int) {
	for i := range x.tasks {
		if x.tasks[i].ID == id {
			x.tasks[i].Complete = !x.tasks[i].Complete
			fmt.Println("Task updated successfully.")
			return
		}
	}
}
