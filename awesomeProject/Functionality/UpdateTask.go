package Functionality

import (
	"fmt"
)

func (x *Tasks) UpdateTask(id int) {
	for i := range x.Tasks {
		if x.Tasks[i].ID == id {
			x.Tasks[i].Complete = !x.Tasks[i].Complete
			fmt.Println("Task updated successfully.")
			return
		}
	}
}
