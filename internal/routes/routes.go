package routes

import (
	"slo19/frontend-masters-project/internal/app"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	//GET
	r.Get("/health", app.HealthCheck)
	r.Get("/workouts/{id}", app.WorkoutHandler.HandleGetWorkoutById)

	//POST
	r.Post("/workouts", app.WorkoutHandler.HandleCreateWorkout)

	//PUT
	r.Put("/workouts/{id}", app.WorkoutHandler.HandleUpdateWorkoutById)
	return r
}
