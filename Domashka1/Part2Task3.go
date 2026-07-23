package main

import (
	"errors"
	"time"
)

type Task1 struct {
	id          string
	title       string
	description string
	priority    int
	status      string
	tags        []string
	assignee    string
	deadline    string
}
type TaskManager struct {
	projects map[string]*Project
	users    map[string]*User
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		projects: make(map[string]*Project),
		users:    make(map[string]*User),
	}
}
func (tm *TaskManager) CreateProject(name string) *Project {
	p := NewProject(name)
	tm.projects[name] = p
	return p
}
func (tm *TaskManager) FindTask(id string) (*Task, error) {
	for _, p := range tm.projects {
		if t, ok := p.tasks[id]; ok {
			return t, nil
		}
	}
	return nil, errors.New("task not found")
}
func (tm *TaskManager) UserTasks(username string) []*Task {
	var result []*Task
	for _, p := range tm.projects {
		for _, t := range p.tasks {
			if t.assignee == username {
				result = append(result, t)
			}
		}
	}
	return result
}
func (tm *TaskManager) Overdue(currentDate string) []*Task {
	current, err := time.Parse("13.08.2004", currentDate)
	if err != nil {
		return nil
	}

	var result []*Task
	for _, p := range tm.projects {
		for _, t := range p.tasks {
			if t.status != "in-progress" {
				continue
			}
			deadline, err := time.Parse("13.08.2004", t.deadline)
			if err != nil {
				continue
			}
			if deadline.Before(current) {
				result = append(result, t)
			}
		}
	}
	return result
}
