package main

import "fmt"

type Priority string

const (
	Low    Priority = "LOW"
	Medium Priority = "MEDIUM"
	High   Priority = "HIGH"
)

func priorityWeight(p Priority) int {
	switch p {
	case High:
		return 3
	case Medium:
		return 2
	default:
		return 1
	}
}

type BaseItem struct {
	ID    string
	Title string
}

func (b BaseItem) Describe() string {
	return fmt.Sprintf("%s %q", b.ID, b.Title)
}

type Task123 struct {
	BaseItem
	Priority   Priority
	status     string
	assignedTo string
}

func NewTask1(id, title string, priority Priority) *Task {
	return &Task{
		BaseItem1: BaseItem{ID: id, Title: title},
		Priority1: priority,
		status:    "todo",
	}
}

func (t *Task) AssignTo(user string) {
	t.assignedTo = user
}

func (t *Task) Start() error {
	if t.status != "todo" {
		return fmt.Errorf("cannot start: task is %q", t.status)
	}
	t.status = "in-progress"
	return nil
}

func (t *Task) Complete() error {
	if t.status == "done" {
		return fmt.Errorf("already completed")
	}
	t.status = "done"
	return nil
}

func (t *Task) IsDone() bool {
	return t.status == "done"
}

func (t *Task) Status() string {
	return t.status
}

func (t Task) String() string {
	return fmt.Sprintf("%s [%s] %-30q %s", t.ID, t.Priority, t.Title, t.status)
}

type Reportable interface {
	Report() string
}

func PrintReports(items []Reportable) {
	for _, r := range items {
		fmt.Println(r.Report())
	}
}

type Project struct {
	Name  string
	Tasks []*Task
}

func NewProject(name string) *Project {
	return &Project{Name: name}
}

func (p *Project) AddTask(t *Task) {
	p.Tasks = append(p.Tasks, t)
}

func (p *Project) FindTask(id string) *Task {
	for _, t := range p.Tasks {
		if t.ID == id {
			return t
		}
	}
	return nil
}

func (p *Project) Progress() (done, total int, pct float64) {
	total = len(p.Tasks)
	for _, t := range p.Tasks {
		if t.IsDone() {
			done++
		}
	}
	if total > 0 {
		pct = float64(done) / float64(total) * 100
	}
	return
}

func (p *Project) NextPriority() *Task {
	var best *Task
	for _, t := range p.Tasks {
		if t.IsDone() {
			continue
		}
		if best == nil || priorityWeight(t.Priority) > priorityWeight(best.Priority) {
			best = t
		}
	}
	return best
}

func (p *Project) UserTasks(user string) []*Task {
	var result []*Task
	for _, t := range p.Tasks {
		if t.assignedTo == user {
			result = append(result, t)
		}
	}
	return result
}

func (p *Project) Report() string {
	done, total, pct := p.Progress()
	return fmt.Sprintf("Project Report: %q — %d/%d done (%.0f%%)", p.Name, done, total, pct)
}

type Sprint struct {
	Number  int
	Start   string
	End     string
	TaskIDs []string
	project *Project
}

func NewSprint(number int, start, end string, project *Project, taskIDs ...string) *Sprint {
	return &Sprint{Number: number, Start: start, End: end, TaskIDs: taskIDs, project: project}
}

func (s *Sprint) Velocity() int {
	done := 0
	for _, id := range s.TaskIDs {
		if t := s.project.FindTask(id); t != nil && t.IsDone() {
			done++
		}
	}
	return done
}

func (s *Sprint) Report() string {
	return fmt.Sprintf("Sprint #%d Report: velocity=%d, 5 days remaining", s.Number, s.Velocity())
}
.