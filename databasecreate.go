package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func create() error {
	// Подключение к базе данных SQLite
	db, err := sql.Open("sqlite3", "./metrics.db")
	if err != nil {
		return err
	}
	defer db.Close()

	// Создание таблицы в базе данных, если она не существует
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS schedule (id INTEGER PRIMARY KEY AUTOINCREMENT, date TEXT NOT NULL, departureAirport TEXT NOT NULL, count INTEGER NOT NULL)")
	if err != nil {
		return err
	}
	return nil
}
