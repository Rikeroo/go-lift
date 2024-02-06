package exercise

import (
	"go-lift/pkg/set"
)

const (
	ErrNoSets          = SetErr("No Sets in this Exercise")
	ErrSetDoesNotExist = SetErr("Set with that index Does not exist")
)

type SetErr string

func (e SetErr) Error() string {
	return string(e)
}

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

func (e *Exercise) DeleteSet(setNum int) error {

	// fmt.Printf("Length of Sets: %v", len(e.Sets))
	switch {
	case len(e.Sets) == 0:
		return ErrNoSets
	case len(e.Sets) == 1:
		e.Sets = []set.Set{}
	case len(e.Sets) > 1:
		e.Sets = append(e.Sets[:setNum-1], e.Sets[setNum])
	}

	return nil

}

func (e *Exercise) GetSetArray(setNum int) ([3]float32, error) {
	// Returns float32 Array of Set info, given setNum
	// Uses Set.Get() Internally

	switch {
	case len(e.Sets) < 1:
		return [3]float32{}, ErrNoSets
	case len(e.Sets) < setNum:
		return [3]float32{}, ErrSetDoesNotExist
	}

	return e.Sets[setNum-1].Get(), nil
}
