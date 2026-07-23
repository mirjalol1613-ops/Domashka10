package main

type Completable interface {
	Complete() error
	Status() string
}

type Prioritized interface {
	Priority() int
	PriorityName() string
}

type WorkItem interface {
	Completable
	Prioritized
	Title() string
	ID() string
}

func NextPriority(items []WorkItem) WorkItem {
	var best WorkItem

	for _, item := range items {
		if item.Status() == "done" {
			continue
		}
		if best == nil {
			best = item
			continue
		}
		if item.Priority() < best.Priority() {
			best = item
		}
	}

	return best
}

func CompleteAll(items []Completable) []error {
	var errs []error

	for _, item := range items {
		err := item.Complete()
		if err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}
