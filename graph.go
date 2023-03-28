package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func graph() {
	db, err := sql.Open("sqlite3", "./metrics.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT date, count FROM schedule ORDER BY date")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var counts plotter.Values
	var dates []string
	for rows.Next() {
		var date string
		var count float64
		if err := rows.Scan(&date, &count); err != nil {
			log.Fatal(err)
		}
		dates = append(dates, date)
		counts = append(counts, count)
	}

	p := plot.New()
	//if err != nil {
	//log.Fatal(err)
	//}

	p.Title.Text = "Flights per Day"
	p.X.Label.Text = "date"
	p.Y.Label.Text = "Count"

	bars, err := plotter.NewBarChart(counts, vg.Points(20))
	if err != nil {
		log.Fatal(err)
	}
	bars.LineStyle.Width = vg.Length(0)

	p.Add(bars)
	p.NominalX(dates...)

	f, err := os.Create("flights.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := p.Save(10*vg.Inch, 10*vg.Inch, "flights.png"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("График успешно создан")
}
