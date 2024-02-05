package exercise

import (
	"go-lift/pkg/set"
	"testing"
)

func TestAddSet(t *testing.T) {
	var sets []set.Set
	exercise := Exercise{Name: "Exercise test-name", Sets: sets}
	weight := 10.0
	reps := 8
	exercise.AddSet(weight, reps)

	want := [3]float32{1, 10.0, 8}
	got := exercise.Sets[0].Get()
}
