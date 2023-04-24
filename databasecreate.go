package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func create() error {
	// Подключение к базе данных SQLite
	db, err := sql.Open("sqlite3", "./metrics.db")
	if err != nil {
		fmt.Println("Ошибка открытия БД: ", err)
		return nil
		//log.Fatal("Ошибка открытия БД: ", err)
	}
	defer db.Close()

	// Создание таблицы в базе данных, если она не существует
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS schedule (id INTEGER PRIMARY KEY AUTOINCREMENT, date TEXT NOT NULL, departureAirport TEXT NOT NULL, count INTEGER NOT NULL)")
	if err != nil {
		fmt.Println("Ошибка создания таблицы в БД: ", err)
		return nil
		//log.Fatal("Ошибка создания таблицы в БД: ", err)
	}
	return nil
}
