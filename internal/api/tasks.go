package api

import (
	"go_final_project/internal/storage"

	"net/http"
)

type TasksResp struct {
	Tasks []*storage.Task `json:"tasks"`
}

func getTasksHandler(w http.ResponseWriter, r *http.Request) {
	maxCount := 50 //задел на пагинацию
	tasks, err := storage.Tasks(maxCount)
	if err != nil {
		writeError(w, err.Error())
		return
	}

	writeJson(w, TasksResp{
		Tasks: tasks,
	})
}
