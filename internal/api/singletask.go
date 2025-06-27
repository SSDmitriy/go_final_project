package api

import (
	"go_final_project/internal/storage"
	"net/http"
)

func getSingleTaskHandler(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json; charset=utf-8")
	id := r.FormValue("id")

	task, err := storage.GetSingleTask(id)
	if err != nil {
		writeError(w, "ошибка при получении задач из базы данных: "+err.Error())
		return
	}

	AwriteJson(w, task)
}
