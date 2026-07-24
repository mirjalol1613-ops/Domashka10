package main

import "fmt"

func main() {
	getMusic()
	getUser()
	getTask()
	GenerateID()
	p := NewProject("Domashka")
	mirjalol := User{
		id:       1,
		username: "mirjalolSh",
		email:    "mijgog@gmail.com",
	}
	p.AddMember(&mirjalol)

	task1 := &Task{
		id:       "TASK-001",
		title:    "Write Project struct",
		priority: 1,
		status:   "in-progress",
	}
	p.AddTask(task1)

	err := p.CompleteTask("TASK-001")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(p.Summary())
	fmt.Println(p.Progress())

	task2 := &Task{
		id:       "TASK-002",
		title:    "Write tests",
		priority: 2,
		status:   "todo",
	}
	p.AddTask(task2)

	items := []WorkItem{task1, task2}
	next := NextPriority(items)
	if next == nil {
		fmt.Println("No incomplete tasks found")
	} else {
		fmt.Println(next.Title())
	}

	completables := []Completable{task1, task2}
	errs := CompleteAll(completables)
	fmt.Println(len(errs))
	s := Sprint{
		Project:   *NewProject("Sprint 1"),
		number:    1,
		startDate: "01.01.2020",
		endDate:   "31.01.2020",
	}

	reports := []Reporter{p, s}
	GenerateReports(reports)

	workTask1 := &Task{id: "T-1", title: "Fix login", priority: 1, status: "todo"}
	bug1 := &Bug{Task: Task{id: "B-1", title: "Crash on save", priority: 1, status: "todo"}, severity: "critical"}
	workTask2 := &Task{id: "T-2", title: "Write docs", priority: 3, status: "todo"}

	for _, item := range items {
		switch v := item.(type) {
		case *Task:
			fmt.Println("Task:", v.Title(), "priority:", v.PriorityName())
		case *Bug:
			fmt.Println("Bug:", v.Title(), "severity:", v.severity)
		default:
			fmt.Println("Unknown work item type")
		}
	}
	bugItems := []WorkItem{workTask1, bug1, workTask2}

	for _, item := range bugItems {
		switch v := item.(type) {
		case *Task:
			fmt.Println("Task:", v.Title(), "priority:", v.PriorityName())
		case *Bug:
			fmt.Println("Bug:", v.Title(), "severity:", v.severity)
		default:
			fmt.Println("Unknown work item type")
		}
	}
	grouped := Triage(bugItems)
	for priority, list := range grouped {
		fmt.Println(priority, "-", len(list), "items")
	}

	var items1 []Completable1 = []Completable1{
		Task12{Title: "Write docs", status: "todo"},
		Task12{Title: "Deploy service", status: "done"},
		Bug12{Title: "Fix login crash", status: "in_progress", Severity: "critical"},
		Bug12{Title: "Typo on homepage", status: "done", Severity: "low"},
	}

	notDone := EstimateCompletion(items1)
	fmt.Printf("Not done: %d out of %d\n", notDone, len(items))
}
