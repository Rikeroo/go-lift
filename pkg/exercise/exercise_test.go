package exercise

import (
	"go-lift/pkg/set"
	"testing"
)

func TestAddSet(t *testing.T) {
	var sets []set.Set
	exercise := Exercise{Name: "Exercise test-name", Sets: sets}
	weight := 10
	reps := 8
	exercise.AddSet(weight, reps)

	//
	//want :=
	// got :=
}
