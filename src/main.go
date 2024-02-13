package main

import (
	"go-lift/handlers"
	"log"
	"net/http"
)

// func workoutsHandler(w http.ResponseWriter, r *http.Request) {
// 	workout := makeTestWorkout()
// 	tmpl, _ := template.ParseFiles("../templates/index.html")

// 	tmpl.Execute(w, workout)
// }

func main() {

	// define variable for Exercise object
	// e1 := w1.Exercises[exerciseName]

	// h1 := func(w http.ResponseWriter, r *http.Request) {
	// 	tmpl := template.Must(template.ParseFiles("../templates/index.html"))
	// 	tmpl.Execute(w, w1)
	// }
	// http.HandleFunc("/", h1)

	// Create route and handler to server tailwind CSS output
	http.HandleFunc("/styles/tailwind", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./css/output.css")
	})

	router := http.NewServeMux()

	router.HandleFunc("GET /workout", handlers.GetAllWorkouts)
	router.HandleFunc("GET /workout/{id}", handlers.GetWorkout)
	router.HandleFunc("DELETE /workout/{id}", handlers.DeleteWorkout)

	// http.HandleFunc("/workouts", workoutsHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
