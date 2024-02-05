package exercise

import (
	"go-lift/pkg/set"
	"testing"
)

func TestAddSet(t *testing.T) {
	t.Run("Add set to empty set slice", func(t *testing.T) {
		var sets []set.Set
		exercise := Exercise{Name: "Exercise test-name", Sets: sets}
		var weight float32 = 10.0
		reps := 8
		exercise.AddSet(weight, reps)

		want := [3]float32{1, 10.0, 8}

		got := [3]float32{0, 0, 0}

		if len(exercise.Sets) > 0 {
			got = exercise.Sets[0].Get()
		}

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

}
