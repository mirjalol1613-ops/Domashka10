package main

import (
	"errors"
	"fmt"
)

type User struct {
	id            int
	username      string
	email         string
	assignedTasks []string
}

func getUser() {
	mirjalol := User{
		id:            1,
		username:      "mirjalolSh",
		email:         "mijgog@gmail.com",
		assignedTasks: []string{"easy", "mid", "hard"},
	}
	fmt.Println(mirjalol)
	fmt.Println(mirjalol.Username())
	fmt.Println(mirjalol.Email())
	mirjalol.AssignTask("Ultra")
	fmt.Println(mirjalol.TaskCount())

}

func (u User) Username() string {
	return u.username
}

func (u User) Email() string {
	return u.email
}

func (u *User) AssignTask(taskID string) {
	u.assignedTasks = append(u.assignedTasks, taskID)
}

func (u *User) UnassignTask(taskID string) error {
	for i, task := range u.assignedTasks {
		if task == taskID {
			u.assignedTasks = append(u.assignedTasks[:i], u.assignedTasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")

}

func (u User) TaskCount() int {
	return len(u.assignedTasks)
}
