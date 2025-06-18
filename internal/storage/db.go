package storage

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

const dbFile = "scheduler.db"

const Schema = `
CREATE TABLE scheduler (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	date CHAR(8) NOT NULL DEFAULT '' CHECK(
        length(date) = 8 AND
        date NOT LIKE '%[^0-9]%' AND
        SUBSTR(date, 1, 4) BETWEEN '2025' AND '2100' AND
		SUBSTR(date, 5, 2) BETWEEN '01' AND '12' AND
        SUBSTR(date, 7, 2) BETWEEN '01' AND '31'
    ),
	title VARCHAR(256) NOT NULL DEFAULT 'Задача:',
	comment TEXT,
	repeat VARCHAR(128)
);

CREATE INDEX tasks_date ON scheduler (date);
`

var (
	db     *sql.DB
	dbPath string
)

func Init(dbFile string) error {

	absPath, err := filepath.Abs(dbFile)
	if err != nil {
		return fmt.Errorf("Ошибка 003 получения пути БД: %s", err)
	}

	dbPath = absPath

	_, err = os.Stat(dbPath)
	install := os.IsNotExist(err)

	if install {
		fmt.Println("База данных не найдена, будет создана.")
		db, err = sql.Open("sqlite", dbPath)
		if err != nil {
			fmt.Println("Ошибка 004 открытия базы данных: ", err)
			return nil
		}

		if _, err := db.Exec(Schema); err != nil {
			fmt.Println("Ошибка 005 создания таблицы базы данных: ", err)
			return nil
		}

		fmt.Println("База данных создана.")
		return nil
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return fmt.Errorf("ошибка подключения к БД: %w", err)
	}

	// db, err = sql.Open("sqlite", dbFile)
	// if err != nil {
	// 	fmt.Println("Ошибка 005 открытия базы данных: ", err)
	// }

	return nil
}

func GetDB() *sql.DB {
	return db
}
