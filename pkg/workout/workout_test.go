package workout

import (
	"go-lift/pkg/exercise"
	"reflect"
	"testing"
	"time"
)

func TestAddExercise(t *testing.T) {
	workout := Workout{Date: time.Now()}

	workout.AddExercise("Test Exercise")
	compareExercise := exercise.Exercise{Name: "Test Exercise"}

	want := compareExercise

	got := exercise.Exercise{}

	if len(workout.Exercises) > 0 {
		got = workout.Exercises[0]
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Got %v wanted %v", got, want)
	}

}

// func makeWorkout(t testing.TB, numWorkouts int) {
// 	// Creates test workout
// 	// workout :=
// }

// func makeExercise(t testing.TB, numSets int) exercise.Exercise {
// 	// Creates Test Exercise
// 	exercise := exercise.Exercise{Name: "Test Exercise"}

// 	for i := 0; i < numSets; i++ {
// 		weight := float32(10 + 2*i)
// 		exercise.AddSet(weight, 8)
// 	}

// 	return exercise
// }
