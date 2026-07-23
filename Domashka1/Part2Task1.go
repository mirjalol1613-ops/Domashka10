package main

import (
	"errors"
	"fmt"
)

type Project struct {
	name    string
	tasks   map[string]*Task
	members []*User
}

func NewProject(name string) *Project {
	return &Project{
		name:  name,
		tasks: make(map[string]*Task),
	}
}

func (p *Project) AddTask(t *Task) {
	p.tasks[t.ID()] = t
}

func (p *Project) AddMember(u *User) {
	p.members = append(p.members, u)
}

func (p *Project) AssignTask(taskID string, user *User) error {
	t, ok := p.tasks[taskID]
	if !ok {
		return errors.New("task not found")
	}
	user.AssignTask(taskID)
	_ = t
	return nil
}

func (p *Project) CompleteTask(taskID string) error {
	t, ok := p.tasks[taskID]
	if !ok {
		return errors.New("task not found")
	}
	return t.Complete()
}

func (p Project) Progress() float64 {
	if len(p.tasks) == 0 {
		return 0
	}
	done := 0
	for _, t := range p.tasks {
		if t.Status() == "done" {
			done++
		}
	}
	return float64(done) / float64(len(p.tasks)) * 100
}

func (p Project) Summary() string {
	return fmt.Sprintf("Project %s: %d tasks, %d members, %.1f%% done",
		p.name, len(p.tasks), len(p.members), p.Progress())
}
