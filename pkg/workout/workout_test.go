package workout

import (
	"go-lift/pkg/exercise"
	"reflect"
	"testing"
	"time"
)

func TestAddExercise(t *testing.T) {
	workout := Workout{Date: time.Now()}

	exerciseName := "Test Exercise"
	workout.AddExercise(exerciseName)
	compareExercise := exercise.Exercise{Name: exerciseName}

	want := compareExercise

	got, ok := workout.Exercises[exerciseName]

	if !ok {
		t.Errorf("No exercise with key: %v", exerciseName)
	}

	if ok {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v want %v", got, want)
		}
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
