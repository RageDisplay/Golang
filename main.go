package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	url      = "https://api.rasp.yandex.net/v3.0/schedule/"
	apiKey   = ""
	station  = "c146"
	platform = "1"
)

type Response struct {
	Metadata Metadata `json:"metadata"`
	Segments []struct {
		Thread  Thread  `json:"thread"`
		Arrival Arrival `json:"arrival"`
	} `json:"segments"`
}

type Metadata struct {
	Total int `json:"total"`
}

type Thread struct {
	Title string `json:"title"`
}

type Arrival struct {
	Time string `json:"time"`
}

func main() {
	db, err := sql.Open("mysql", "root:12345@/yandexapi")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	resp, err := http.Get(url + fmt.Sprintf("electronic/%s/%s/arrivals/?apikey=%s", station, platform, apiKey))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}

	for _, segment := range response.Segments {
		name := segment.Thread.Title
		value, err := strconv.ParseFloat(segment.Arrival.Time, 64)
		if err != nil {
			panic(err)
		}
		timestamp := time.Now().UTC()

		_, err = db.Exec("INSERT INTO metrics(name, value, timestamp) VALUES($1, $2, $3)", name, value, timestamp)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s: %.2f\n", name, value)
	}
}
