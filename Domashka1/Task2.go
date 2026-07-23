package main

import (
	"errors"
	"fmt"
)

type NewTask struct {
	title       string
	description string
	priority    int
	id          string
}

func getMusic() {
	music := NewTask{
		title:       "music",
		description: "good music",
		priority:    1,
		id:          "TASK-000",
	}
	fmt.Println(music)
	fmt.Println(music.Title())
	fmt.Println(music.Priority())
	fmt.Println(GenerateID())

}

func (n NewTask) Title() string {
	if n.title != "" {
		return n.title
	}
	return "title is empty"
}

func (n NewTask) Priority() (int, error) {
	if n.priority >= 1 && n.priority <= 3 {
		return n.priority, nil
	}
	return 0, errors.New("priority must be between 1 and 3")
}

var counter = 1

func GenerateID() string {
	id := fmt.Sprintf("TASK-%03d", counter)
	counter++
	return id
}
