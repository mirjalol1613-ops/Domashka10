package main

import (
	"errors"
	"time"
)

type Sprint struct {
	Project
	number    int
	startDate string
	endDate   string
	velocity  int
}

func (s *Sprint) AddTask(t *Task) error {
	if t.status != "todo" {
		return errors.New("only tasks with status 'todo' can be added to a sprint")
	}
	s.Project.AddTask(t)
	return nil
}
func (s Sprint) IsActive() bool {
	start, err1 := time.Parse("13.08.2004", s.startDate)
	end, err2 := time.Parse("13.08.2004", s.endDate)
	if err1 != nil || err2 != nil {
		return false
	}
	now := time.Now()
	return now.After(start) && now.Before(end)
}

func (s Sprint) DaysRemaining() int {
	end, err := time.Parse("13.08.2004", s.endDate)
	if err != nil {
		return 0
	}
	remaining := int(time.Until(end).Hours() / 24)
	if remaining < 0 {
		return 0
	}
	return remaining
}
