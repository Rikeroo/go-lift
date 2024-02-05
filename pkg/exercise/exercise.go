package exercise

import "go-lift/pkg/set"

type Exercise struct {
	// Takes name and slice of 'Set' objects
	Name string
	Sets []set.Set
}

func (e *Exercise) AddSet(weight float32, reps int) {
	// return [3]float32{0, 0, 0}
}
