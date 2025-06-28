package api

import (
	"fmt"
	"go_final_project/internal/storage"
	"go_final_project/internal/util"
	"time"
)

func checkDate(task *storage.Task) error {
	now := time.Now()
	var nextDate string

	if task.Date == "" {
		task.Date = now.Format(util.DateFormat)
	}

	t, err := time.Parse(util.DateFormat, task.Date)
	if err != nil {
		return fmt.Errorf("ошибка checkDate - неверный формат даты")
	}

	if task.Repeat != "" {
		nextDate, err = util.NextTaskDate(now, task.Date, task.Repeat)
		if err != nil {
			return err
		}
	}

	if util.AfterNow(t, now) {
		if len(task.Repeat) == 0 {
			// если правила повторения нет, то берём сегодняшнее число
			task.Date = now.Format(util.DateFormat)
		} else {
			// в противном случае, берём вычисленную ранее следующую дату
			task.Date = nextDate
		}
	}

	return nil
}
