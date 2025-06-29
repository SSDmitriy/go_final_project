package api

import (
	"encoding/json"
	"go_final_project/internal/storage"

	"net/http"
)

func addTaskHandler(w http.ResponseWriter, r *http.Request) {

	var task storage.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeError(w, "ошибка addTaskHandler - неверный формат JSON: "+err.Error())
		return
	}

	if task.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		writeError(w, "ошибка addTaskHandler - не указано название задачи")
		return
	}

	if err := checkDate(&task); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeError(w, err.Error())
		return
	}

	id, err := storage.AddTask(&task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeError(w, "ошибка добавления задачи в базу данных: "+err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	writeJson(w, map[string]interface{}{"id": id})
}
