package exercise

import "go-lift/pkg/set"

type Exercise struct {
	// Takes name and slice of 'Set' objects
	Name string
	Sets []set.Set
}
