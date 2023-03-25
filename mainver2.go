package main

import (
"fmt"
"log"
"database/sql"
"net/http"
"encoding/json"
_ "github.com/mattn/go-sqlite3"
)

type ScheduleResponse struct {
Dates []struct {
Date string `json:"date"`
Schedule []struct {
Thread struct {
Flight string `json:"flight"`
} `json:"thread"`
} `json:"schedule"`
} `json:"dates"`
}

func main() {
// Установка значений параметров для запроса
departureAirport := "SVO"
date := "2023-03-25"

// Формирование URL для запроса к API Яндекс.Расписание
url := fmt.Sprintf("https://api.rasp.yandex.net/v3.0/schedule/?apikey=%s&format=json&station=%s&date=%s&transport_types=plane", "751ddb3c-701e-480c-bf88-9327b8543f92", departureAirport, date)

// Выполнение запроса к API Яндекс.Расписание
response, err := http.Get(url)
if err != nil {
log.Fatal(err)
}
defer response.Body.Close()

// Чтение ответа API и декодирование его в структуру ScheduleResponse
var scheduleResponse ScheduleResponse
err = json.NewDecoder(response.Body).Decode(&scheduleResponse)
if err != nil {
log.Fatal(err)
}

// Подключение к базе данных SQLite
db, err := sql.Open("sqlite3", "metrics.db")
if err != nil {
log.Fatal(err)
}
defer db.Close()

// Создание таблицы в базе данных, если она не существует
_, err = db.Exec("CREATE TABLE IF NOT EXISTS metrics (date TEXT, airport TEXT, count INTEGER)")
if err != nil {
log.Fatal(err)
}

// Вычисление количества рейсов и сохранение результата в базе данных
count := 0
/*for _, dateData := range scheduleResponse.Dates {
for _, scheduleData := range dateData.Schedule {
count++
}
}*/
_, err = db.Exec("INSERT INTO metrics VALUES (?, ?, ?)", date, departureAirport, count)
if err != nil {
log.Fatal(err)
}
}