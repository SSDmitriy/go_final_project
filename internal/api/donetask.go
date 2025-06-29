package api

import (
	"go_final_project/internal/storage"

	"net/http"
)

func doneTaskHandler(w http.ResponseWriter, r *http.Request) {

	idStr := r.FormValue("id")

	task, err := storage.GetSingleTask(idStr)
	if err != nil {
		writeError(w, err.Error())
		return
	}

	if task.Repeat == "" {
		err := storage.DeleteTask(idStr)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			writeError(w, err.Error())
			return
		}
	} else {
		err := storage.UpdateDate(idStr)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			writeError(w, err.Error())
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	writeJson(w, struct{}{})
}
