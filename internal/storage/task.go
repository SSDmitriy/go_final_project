package storage

import (
	"database/sql"
	"fmt"
	"strconv"
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

	if err == nil {
		id, err = res.LastInsertId()
	}
	return id, err
}

func Tasks(limit int) ([]*Task, error) {
	tasks := make([]*Task, 0)

	//selectQuery := fmt.Sprintf("SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date LIMIT %d", limit)
	//selectQuery := fmt.Sprintf("SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date LIMIT :limit", sql.Named("limit", limit))
	selectQuery := "SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date LIMIT ?"

	rows, err := db.Query(selectQuery, limit)
	if err != nil {
		return tasks, err
	}
	defer rows.Close()

	for rows.Next() {
		task := new(Task)
		err := rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			return tasks, err
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return tasks, err
	}

	return tasks, nil
}

func GetSingleTask(idStr string) (*Task, error) {
	task := &Task{}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, fmt.Errorf("ошибка - невозможно конвертировать id в int: %s", err)
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
			return nil, fmt.Errorf("задача с id %d не найдена", id)
		}
		// Возвращаем другие ошибки бд
		return nil, fmt.Errorf("ошибка при получении задачи из базы данных: %v", err)
	}

	return task, nil
}
