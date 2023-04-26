package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func clear() error {
	// Открытие базы данных SQLite
	db, err := sql.Open("sqlite3", "./metrics.db")
	if err != nil {
		fmt.Println(err)
		return nil
		//panic(err)
	}
	defer db.Close()

	// Получение количества записей в таблице
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM schedule").Scan(&count)
	if err != nil {
		fmt.Println(err)
		return nil
		//panic(err)
	}

	// Проверка количества записей и удаление первой, если оно превышает 12
	if count > 12 {
		_, err = db.Exec("DELETE FROM schedule WHERE id = (SELECT MIN(id) FROM schedule)")
		if err != nil {
			fmt.Println(err)
			return nil
			//panic(err)
		}
	}
	return nil
}

func clearlast() error {
	// Открытие базы данных SQLite
	db, err := sql.Open("sqlite3", "./metrics.db")
	if err != nil {
		fmt.Println(err)
		return nil
		//panic(err)
	}
	defer db.Close()

	// Получение количества записей в таблице
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM schedule").Scan(&count)
	if err != nil {
		fmt.Println(err)
		return nil
		//panic(err)
	}

	_, err = db.Exec("DELETE FROM schedule WHERE id = (SELECT MAX(id) FROM schedule)")
	if err != nil {
		fmt.Println(err)
		return nil
		//panic(err)
	}
	return nil
}
