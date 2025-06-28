package api

import (
	"go_final_project/internal/storage"
	"net/http"
)

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	err := storage.DeleteTask(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeError(w, "ошибка удаления задачи в базе данных: "+err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	writeJson(w, struct{}{})
}
