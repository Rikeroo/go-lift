package workout

import (
	"go-lift/pkg/exercise"
	"time"
)

type Workout struct {
	Date      time.Time
	Exercises []exercise.Exercise
}

func (w *Workout) AddExercise(name string) {
	// Creates new Exercise within a workout with
	// empty Sets slice
	exercise := exercise.Exercise{Name: name}
	w.Exercises = append(w.Exercises, exercise)

}
