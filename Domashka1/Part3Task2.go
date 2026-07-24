package main

import "fmt"

type Reporter interface {
	Report() string
}

func (p Project) Report() string {
	return fmt.Sprintf("Progress Report - %s: %.1f%% complete (%d tasks)", p.name, p.Progress(), len(p.tasks))
}
func (s Sprint) Report() string {
	return fmt.Sprintf("Sprint %d Report - %s: velocity =%d, %s", s.number, s.name, s.velocity, s.Project.Report())
}
func GenerateReports(reports []Reporter) {
	for _, r := range reports {
		fmt.Println(r.Report())
	}
}
