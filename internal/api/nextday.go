package api

import (
	"go_final_project/internal/util"
	"net/http"
	"time"
)

// пример запроса POST: api/nextdate?now=20240126&date=20240126&repeat=y
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
			w.Write([]byte("ошибка nextDayHandler - не удалось распарсить дату"))
			return
		}
	}

	startDateStr := r.FormValue("date")
	repeatRuleStr := r.FormValue("repeat")

	nextDay, err := util.NextTaskDate(nowDate, startDateStr, repeatRuleStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(nextDay))
}
