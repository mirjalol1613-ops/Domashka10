package main

type Bug struct {
	Task
	severity string
}

func Triage(items []WorkItem) map[string][]WorkItem {
	result := make(map[string][]WorkItem)
	for _, item := range items {
		name := item.PriorityName()
		result[name] = append(result[name], item)
	}
	return result
}
