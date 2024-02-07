package workout

import (
	"go-lift/pkg/exercise"
	"time"
)

type ExercisesDict map[string]*exercise.Exercise

type Workout struct {
	Date      time.Time
	Exercises ExercisesDict
}

func (w *Workout) AddExercise(name string) {
	// Creates new Exercise within a workout with
	// empty Sets slice

	// If workout doesn't have Exercise dictionary initialised,
	// then initialise one
	if w.Exercises == nil {
		w.Exercises = make(ExercisesDict)
	}

	exercise := exercise.Exercise{Name: name}
	w.Exercises[name] = &exercise

}

func (w *Workout) AddSet(name string, weight float32, reps int) {
	w.Exercises[name].AddSet(weight, reps)

}
