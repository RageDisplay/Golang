package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func clear() error {
	// Открытие базы данных SQLite
	db, err := sql.Open("sqlite3", "./metrics.db")
	if err != nil {
		return err
	}
	defer db.Close()

	// Получение количества записей в таблице
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM schedule").Scan(&count)
	if err != nil {
		return err
	}

	// Проверка количества записей и удаление первой, если оно превышает 12
	if count > 12 {
		_, err = db.Exec("DELETE FROM schedule WHERE id = (SELECT MIN(id) FROM schedule)")
		if err != nil {
			return err
		}
	}
	return nil
}

func clearlast() error {
	// Открытие базы данных SQLite
	db, err := sql.Open("sqlite3", "./metrics.db")
	if err != nil {
		return err
	}
	defer db.Close()

	// Получение количества записей в таблице
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM schedule").Scan(&count)
	if err != nil {
		return err
	}

	// Проверка количества записей и удаление первой, если оно превышает 12
	_, err = db.Exec("DELETE FROM schedule WHERE id = (SELECT MAX(id) FROM schedule)")
	if err != nil {
		return err
	}
	return nil
}
