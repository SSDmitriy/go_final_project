package api

import (
	"encoding/json"
	"go_final_project/internal/storage"

	"net/http"
)

type TasksResp struct {
	Tasks []*storage.Task `json:"tasks"`
}

func getTasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	maxCount := 50
	tasks, err := storage.Tasks(maxCount)
	if err != nil {
		writeError(w, "ошибка при получении задач из базы данных: "+err.Error())
		return
	}

	writeJson(w, TasksResp{
		Tasks: tasks,
	})
}

func writeError(w http.ResponseWriter, errorMsg string) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(struct {
		Error string `json:"error"`
	}{
		Error: errorMsg,
	})
}

func writeJson(w http.ResponseWriter, data interface{}) {
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "ошибка формирования JSON", http.StatusInternalServerError)
	}
}
