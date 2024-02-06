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

		got := [3]float32{}

		if len(exercise.Sets) > 0 {
			got = exercise.Sets[0].Get()
		}

		assertArrays(t, got, want)
	})

}

func TestDeleteSet(t *testing.T) {
	t.Run("Delete Set from Sets slice with length = 1", func(t *testing.T) {
		exercise := makeExercise(t, 1)

		setNum := 1
		DelErr := exercise.DeleteSet(setNum)

		want := [3]float32{}
		got, GetErr := exercise.GetSetArray(setNum)

		assertError(t, DelErr, nil)
		assertError(t, GetErr, ErrNoSets)
		assertArrays(t, got, want)
	})

}

func TestGetSetArray(t *testing.T) {
	exercise := makeExercise(t, 1)

	want := [3]float32{1, 10, 8}
	got, err := exercise.GetSetArray(1)

	assertError(t, err, nil)
	assertArrays(t, got, want)
}

func assertArrays(t testing.TB, got, want [3]float32) {
	t.Helper()

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func makeExercise(t testing.TB, setNums int) Exercise {
	sets := []set.Set{}

	for i := 0; i < setNums; i++ {
		testSet := set.Set{SetNumber: i + 1, Weight: 10.0, Reps: 8}
		sets = append(sets, testSet)
	}
	return Exercise{Name: "Exercise test-name", Sets: sets}
}
