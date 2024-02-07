package main

import (
	"go-lift/pkg/workout"
	"time"
)

func main() {
	// Typical user journey

	// Initiate new workout at current date
	w1 := workout.Workout{Date: time.Now()}

	// Add new workout with name
	exerciseName := "Incline DB Press"
	w1.AddExercise(exerciseName)

	// Add set to exercise
	(w1.Exercises[exerciseName]).AddSet(10, 8)

}
