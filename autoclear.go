package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func clear() {
	// Открытие базы данных SQLite
	db, err := sql.Open("sqlite3", "./metrics.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Получение количества записей в таблице
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM schedule").Scan(&count)
	if err != nil {
		panic(err)
	}

	// Проверка количества записей и удаление первой, если оно превышает 15
	if count > 15 {
		_, err = db.Exec("DELETE FROM schedule WHERE id = (SELECT MIN(id) FROM schedule)")
		if err != nil {
			panic(err)
		}
	}
}
