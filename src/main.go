package main

import (
	"fmt"
	"go-lift/pkg/workout"
	"html/template"
	"log"
	"net/http"
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
	w1.Exercises[exerciseName].AddSet(10, 8)
	w1.Exercises[exerciseName].AddSet(12, 8)
	w1.Exercises[exerciseName].AddSet(14, 6)

	// Adding another exercise
	exerciseName = "Squat"
	w1.AddExercise(exerciseName)

	// Add Sets to exercise
	w1.Exercises[exerciseName].AddSet(30, 8)
	w1.Exercises[exerciseName].AddSet(50, 8)
	w1.Exercises[exerciseName].AddSet(70, 8)

	// define variable for Exercise object
	// e1 := w1.Exercises[exerciseName]

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./templates/index.html"))
		tmpl.Execute(w, w1)
	}
	http.HandleFunc("/", h1)

	startServer()
}

func startServer() {
	fmt.Println("Starting Server...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// func workoutHandler (workout workout.Workout) {
// 	// Takes workout object
// }
