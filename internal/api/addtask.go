package api

import (
	"encoding/json"
	"fmt"
	"go_final_project/internal/storage"
	"go_final_project/internal/util"

	"net/http"
	"time"
)

func addTaskHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var task storage.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Неверный формат JSON: " + err.Error()})
		return
	}

	if task.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Не указано название задачи"})
		return
	}

	if err := checkDate(&task); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	id, err := storage.AddTask(&task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Ошибка добавления задачи в базу данных: " + err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"id": id})
}

func checkDate(task *storage.Task) error {
	now := time.Now()
	var nextDate string

	if task.Date == "" {
		task.Date = now.Format("20060102")
	}

	t, err := time.Parse("20060102", task.Date)
	if err != nil {
		return fmt.Errorf("ошибка - неверный формат даты")
	}

	if task.Repeat != "" {
		nextDate, err = util.NextTaskDate(now, task.Date, task.Repeat)
	}

	if util.AfterNow(now, t) {
		if len(task.Repeat) == 0 {
			// если правила повторения нет, то берём сегодняшнее число
			task.Date = now.Format("20060102")
		} else {
			// в противном случае, берём вычисленную ранее следующую дату
			task.Date = nextDate
		}
	}

	return nil
}
