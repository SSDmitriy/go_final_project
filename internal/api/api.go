package api

import (
	"go_final_project/internal/util"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

// регистрация обработчиков маршрутов
func Init(r *chi.Mux) {
	r.HandleFunc("/api/nextdate", nextDayHandler)

	// r.Get("/tasks", getTasks)
	// r.Post("/tasks", postTasks)
	// r.Get("/tasks/{id}", getTask)
	// r.Delete("/tasks/{id}", deleteTask)

}

// "api/nextdate?now=20240126&date=20240126&repeat=y"
func nextDayHandler(w http.ResponseWriter, r *http.Request) {
	nowStr := r.FormValue("now")

	var nowDate time.Time
	if nowStr == "" {
		nowDate = time.Now()
	} else {
		var err error
		nowDate, err = time.Parse(util.DateFormat, nowStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("ошибка 201 парсинга текущей даты"))
			return
		}
	}

	startDate := r.FormValue("date")
	repeatRule := r.FormValue("repeat")

	nextDay, err := util.NextTaskDate(nowDate, startDate, repeatRule)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(nextDay))
}
