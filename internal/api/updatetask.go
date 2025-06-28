package api

import (
	"encoding/json"
	"go_final_project/internal/storage"

	"net/http"
)

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {

	var task storage.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeError(w, "ошибка updateTaskHandler - неверный формат JSON: "+err.Error())
		return
	}

	if task.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		writeError(w, "ошибка updateTaskHandler - не указано название задачи")
		return
	}

	if err := checkDate(&task); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeError(w, err.Error())
		return
	}

	err := storage.UpdateTask(&task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeError(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	writeJson(w, struct{}{})
}
