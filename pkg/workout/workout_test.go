package workout

import (
	"go-lift/pkg/exercise"
	"reflect"
	"testing"
	"time"
)

const exerciseName = "Test Exercise"

func TestAddExercise(t *testing.T) {
	workout := Workout{Date: time.Now()}

	workout.AddExercise(exerciseName)
	compareExercise := &exercise.Exercise{Name: exerciseName}

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

func TestAddSet(t *testing.T) {
	workout := Workout{Date: time.Now()}
	workout.AddExercise(exerciseName)

	var weight float32 = 10.0
	reps := 8

	workout.AddSet(exerciseName, weight, reps)
	// fmt.Println(workout.Exercises[exerciseName].Sets)

	setsLength := len(workout.Exercises[exerciseName].Sets)
	// fmt.Println(setsLength)

	// Gets last element in Sets slice
	exercise := workout.Exercises[exerciseName]
	got, _ := exercise.GetSetArray(setsLength)

	want := [3]float32{float32(setsLength), weight, float32(reps)}

	if got != want {
		t.Errorf("Got %v want %v", got, want)
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
