package exercise

import "go-lift/pkg/set"

type Exercise struct {
	// Takes name and slice of 'Set' objects
	Name string
	Sets []set.Set
}

func (e *Exercise) AddSet(weight float32, reps int) {
	newSetNum := len(e.Sets) + 1
	newSet := set.Set{SetNumber: newSetNum, Weight: weight, Reps: reps}
	e.Sets = append(e.Sets, newSet)
}
