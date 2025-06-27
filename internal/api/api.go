package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// регистрация обработчиков маршрутов
func Init(r *chi.Mux) {
	r.Get("/api/nextdate", nextDayHandler)
	r.Post("/api/task", addTaskHandler)
	r.Get("/api/tasks", getTasksHandler)
	r.Get("/api/task", getSingleTaskHandler)

	// r.Post("/tasks", postTasks)
	// r.Delete("/tasks/{id}", deleteTask)

}

func AwriteJson(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(data)

}
