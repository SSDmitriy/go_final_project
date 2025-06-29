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

	if len(task.Date) == 0 {
		task.Date = now.Format(util.DateFormat)
	}

	t, err := time.Parse(util.DateFormat, task.Date)
	if err != nil {
		return fmt.Errorf("ошибка checkDate - неверный формат даты")
	}

	if len(task.Repeat) != 0 {
		nextDate, err = util.NextTaskDate(now, t.Format(util.DateFormat), task.Repeat)
		if err != nil {
			return err
		}
	}

	if util.AfterNow(now, t) {
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
