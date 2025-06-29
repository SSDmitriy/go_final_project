package api

import (
	"go_final_project/internal/storage"
	"net/http"
)

func getSingleTaskHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.FormValue("id")

	task, err := storage.GetSingleTask(idStr)
	if err != nil {
		writeError(w, err.Error())
		return
	}

	writeJson(w, task)
}
