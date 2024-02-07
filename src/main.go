package main

import (
	"go-lift/pkg/workout"
	"time"
)

func main() {
	// Typical user journey
	w1 := workout.Workout{Date: time.Now()}

	w1.AddExercise("Incline DB Press")

	// w1.Exercises[]

}
