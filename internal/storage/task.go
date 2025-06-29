package storage

import (
	"database/sql"
	"fmt"
	"go_final_project/internal/util"
	"strconv"
	"time"
)

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

func AddTask(task *Task) (int64, error) {
	var id int64

	insertQuery := `INSERT INTO scheduler (date, title, comment, repeat) VALUES (:date, :title, :comment, :repeat)`

	res, err := db.Exec(insertQuery,
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat))

	if err != nil {
		return -1, fmt.Errorf("ошибка AddTask - не удалось выполнить INSERT %v", err)
	} else {
		id, err = res.LastInsertId()

		if err != nil {
			return -2, fmt.Errorf("ошибка LastInsertId %v", err)
		}
	}
	return id, nil
}

func Tasks(limit int) ([]*Task, error) {
	tasks := make([]*Task, 0)

	selectQuery := "SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date LIMIT ?"

	rows, err := db.Query(selectQuery, limit)
	if err != nil {
		return tasks, fmt.Errorf("ошибка Tasks - не удалось выполнить запрос SELECT: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		task := new(Task)
		err := rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			return tasks, fmt.Errorf("ошибка rows.Scan %v", err)
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return tasks, fmt.Errorf("ошибка rows.Err %v", err)
	}

	return tasks, nil
}

func GetSingleTask(idStr string) (*Task, error) {
	task := &Task{}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, fmt.Errorf("ошибка - невозможно конвертировать id в int: %v", err)
	}

	selectSingleQuery := "SELECT * FROM scheduler WHERE id = ?"

	err = db.QueryRow(selectSingleQuery, id).Scan(
		&task.ID,
		&task.Date,
		&task.Title,
		&task.Comment,
		&task.Repeat,
	)

	if err != nil {
		//ошибка, если задача не найдена
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("ошибка GetSingleTask - задача с id %v не найдена", id)
		}
		// Возвращаем другие ошибки бд
		return nil, fmt.Errorf("ошибка GetSingleTask при получении задачи из базы данных: %v", err)
	}

	return task, nil
}

func UpdateTask(task *Task) error {

	updateQuery := `UPDATE scheduler SET date = :date, title = :title, comment = :comment, repeat = :repeat WHERE id = :id`

	res, err := db.Exec(updateQuery,
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat),
		sql.Named("id", task.ID))

	if err != nil {
		return fmt.Errorf("ошибка UpdateTask - не удалось выполнить UPDATE: %v", err)
	}

	// метод RowsAffected() возвращает количество записей к которым
	// была применена SQL команда
	count, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка res.RowsAffected: %v", err)
	}
	if count == 0 {
		return fmt.Errorf("ошибка UpdateTask - некорректный id задачи")
	}

	return nil
}

func UpdateDate(idStr string) error {

	task, err := GetSingleTask(idStr)
	if err != nil {
		return fmt.Errorf("ошибка UpdateDate - невозможно получить задачу из базы: %s", err)
	}

	newDate, err := util.NextTaskDate(time.Now(), task.Date, task.Repeat)
	if err != nil {
		return fmt.Errorf("ошибка UpdateDate - не удалось вычислить следующую дату: %s", err)
	}

	updateDateQuery := `UPDATE scheduler SET date = :newDate WHERE id = :id`

	res, err := db.Exec(updateDateQuery,
		sql.Named("newDate", newDate),
		sql.Named("id", task.ID))

	if err != nil {
		return fmt.Errorf("ошибка UpdateDate - не удалось выполнить запрос UPDATE: %v", err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка res.RowsAffected: %v", err)
	}
	if count == 0 {
		return fmt.Errorf("ошибка UpdateDate - некорректный id задачи")
	}

	return nil
}

func DeleteTask(idStr string) error {

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("ошибка DeleteTask - невозможно конвертировать id в int: %v", err)
	}

	deleteQuery := `DELETE FROM scheduler WHERE id = :id`

	res, err := db.Exec(deleteQuery,
		sql.Named("id", id))
	if err != nil {
		return fmt.Errorf("ошибка DeleteTask - не удалось выполнить запрос DELETE")
	}

	count, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка RowsAffected: %v", err)
	}
	if count == 0 {
		return fmt.Errorf("ошибка DeleteTask - задача с таким id не найдена")
	}

	return nil
}
