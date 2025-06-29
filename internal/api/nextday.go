package api

import (
	"go_final_project/internal/util"
	"net/http"
	"time"
)

func nextDayHandler(w http.ResponseWriter, r *http.Request) {
	var nowDate time.Time
	nowStr := r.FormValue("now")

	if len(nowStr) == 0 {
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
