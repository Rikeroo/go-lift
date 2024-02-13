package main

import (
	"go-lift/pkg/workout"
	"log"
	"net/http"
	"text/template"
	"time"
)

func makeTestWorkout() workout.Workout {
	// Returns test workout with 2 exercises and 3
	// sets each

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

	return w1
}

func workoutsHandler(w http.ResponseWriter, r *http.Request) {
	workout := makeTestWorkout()
	tmpl, _ := template.ParseFiles("../templates/index.html")

	tmpl.Execute(w, workout)
}

func main() {

	// define variable for Exercise object
	// e1 := w1.Exercises[exerciseName]

	// h1 := func(w http.ResponseWriter, r *http.Request) {
	// 	tmpl := template.Must(template.ParseFiles("../templates/index.html"))
	// 	tmpl.Execute(w, w1)
	// }
	// http.HandleFunc("/", h1)

	http.HandleFunc("/styles/tailwind", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./css/output.css")
	})

	router := http.NewServeMux()
	

	http.HandleFunc("/workouts", workoutsHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
