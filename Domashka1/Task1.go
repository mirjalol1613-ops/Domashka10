package main

import (
	"errors"
	"fmt"
)

type Task struct {
	id          string
	title       string
	description string
	priority    int
	status      string
	tags        []string
	assignee    string
	deadline    string
}

func getTask() {
	var Music = Task{
		id:          "qlcd",
		title:       "Goal",
		description: "Good",
		priority:    1,
		status:      "in-progress",
		tags:        []string{"classic", "pop", "phonk"},
	}
	fmt.Println(Music.ID())
	Music.AddTag("funk")
	fmt.Println(Music.HasTag("funk"))
}

func (t Task) ID() string {
	return t.id
}

func (t Task) Title() string {
	return t.title
}

func (t Task) Priority() int {
	return t.priority
}

func (t Task) Status() string {
	return t.status
}

func (t Task) PriorityName() string {
	if t.priority == 1 {
		return "High"
	} else if t.priority == 2 {
		return "Medium"
	} else if t.priority == 3 {
		return "Low"
	}
	return "unknown"
}

func (t *Task) AddTag(tag string) {
	t.tags = append(t.tags, tag)

}

func (t Task) HasTag(tag string) bool {
	for _, Curtag := range t.tags {
		if Curtag == tag {
			return true
		}
	}
	return false
}

func (t *Task) Complete() error {
	if t.status == "done" {
		return errors.New("Music already playing")
	}
	t.status = "done"

	return nil
}

func (t *Task) Start() error {
	if t.status == "done" || t.status == "in-progress" {
		return errors.New("Music already playing")
	}
	t.status = "in-progress"
	return nil
}
