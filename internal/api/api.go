package api

import (
	"github.com/go-chi/chi/v5"
)

// регистрация обработчиков маршрутов
func Init(r *chi.Mux) {
	r.Get("/api/nextdate", nextDayHandler)
	r.Post("/api/task", addTaskHandler)

	// r.Get("/tasks", getTasks)
	// r.Post("/tasks", postTasks)
	// r.Get("/tasks/{id}", getTask)
	// r.Delete("/tasks/{id}", deleteTask)

}
