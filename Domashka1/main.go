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
}
