package api

import (
	"go_final_project/internal/storage"
	"net/http"
)

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.FormValue("id")

	err := storage.DeleteTask(idStr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeError(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	writeJson(w, struct{}{})
}
