package storage

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

const Schema = `
CREATE TABLE scheduler (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	date CHAR(8) NOT NULL DEFAULT '' CHECK(
        length(date) = 8
		AND (date NOT LIKE '%[^0-9]%')
		AND (SUBSTR(date, 1, 4) BETWEEN '2000' AND '2100')
		AND (SUBSTR(date, 5, 2) BETWEEN '01' AND '12')
		AND (SUBSTR(date, 7, 2) BETWEEN '01' AND '31')
    ),
	title VARCHAR(256) NOT NULL DEFAULT 'Задача:',
	comment TEXT,
	repeat VARCHAR(128)
);

CREATE INDEX tasks_date ON scheduler (date);
`

var (
	db *sql.DB
)

func Init(dbFile string) error {

	_, err := os.Stat(dbFile)
	install := os.IsNotExist(err)

	db, err = sql.Open("sqlite", dbFile)
	if err != nil {
		return fmt.Errorf("ошибка Init - не удалось открыть базу данных: %v", err)
	}

	if install {
		fmt.Println("База данных не найдена, будет создана новая.")

		if _, err := db.Exec(Schema); err != nil {
			return fmt.Errorf("ошибка Init - не удалось создать таблицу базы данных: %v", err)
		}

		fmt.Println("База данных создана.")
	}

	return nil
}

func GetDB() *sql.DB {
	return db
}

func Close() error {
	if db != nil {
		return db.Close()
	}
	return nil
}
