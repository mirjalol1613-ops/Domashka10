package main

type Task12 struct {
	Title  string
	status string
}

func (t Task12) Status1() string {
	return t.status
}

type Bug12 struct {
	Title    string
	Severity string
	status   string
}

func (b Bug12) Status1() string {
	return b.status
}

type Completable1 interface {
	Status1() string
}

func EstimateCompletion(items []Completable1) int {
	count := 0
	for _, item := range items {
		if item.Status1() != "done" {
			count++
		}
	}
	return count
}
