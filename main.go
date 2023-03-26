package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type ScheduleResponse struct {
	Dates []struct {
		Date     string `json:"date"`
		Schedule []struct {
			Thread struct {
				Number string `json:"number"`
			} `json:"thread"`
		} `json:"schedule"`
	} `json:"dates"`
}

func main() {
	// Установка значений параметров для запроса
	departureAirport := "s9600213"
	date := "2023-03-25"
	url := fmt.Sprintf("https://api.rasp.yandex.net/v3.0/schedule/?apikey=%s&format=json&station=%s&date=%s&transport_types=plane", "751ddb3c-701e-480c-bf88-9327b8543f92", departureAirport, date)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	//Создание JSON

	file, err := os.Create("SVO.json")
	if err != nil {
		fmt.Println("Ошибка в создании json:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.Write(body)

	// Разбор JSON-ответа
	var scheduleResponse ScheduleResponse
	err = json.Unmarshal(body, &scheduleResponse)
	if err != nil {
		log.Fatal("Ошибка разбора JSON-ответа API Яндекс.Расписание: ", err)
	}

	// Подключение к базе данных SQLite
	db, err := sql.Open("sqlite3", "metrics.db")
	if err != nil {
		log.Fatal("Ошибка открытия БД: ", err)
	}
	defer db.Close()

	// Создание таблицы в базе данных, если она не существует
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS schedule (id INTEGER PRIMARY KEY AUTOINCREMENT, date TEXT NOT NULL, departureAirport TEXT NOT NULL, count INTEGER NOT NULL)")
	if err != nil {
		log.Fatal("Ошибка создания таблицы в БД: ", err)
	}

	// Запись метрики в базу данных
	count := len(scheduleResponse.Dates[0].Schedule)
	stmt, err := db.Prepare("INSERT INTO train_schedule(date, departureAirport, count) values(?,?,?)")
	if err != nil {
		log.Fatal("Ошибка подготовки запроса к БД: ", err)
	}
	_, err = stmt.Exec(scheduleResponse.Dates[0].Date, departureAirport, count)
	if err != nil {
		log.Fatal("Ошибка выполнения запроса к БД: ", err)
	}

	log.Printf("Метрика успешно записана в базу данных (Дата: %s, Станция: %s, Количество рейсов: %d)", scheduleResponse.Dates[0].Date, departureAirport, count)
}
