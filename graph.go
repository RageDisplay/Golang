package main

import (
"database/sql"
"fmt"
"log"
"os"
"gonum.org/v1/plot"
"gonum.org/v1/plot/plotter"
"gonum.org/v1/plot/vg"
_ "github.com/mattn/go-sqlite3"
)

func main() {
db, err := sql.Open("sqlite3", "./metrics.db")
if err != nil {
log.Fatal(err)
}
defer db.Close()

rows, err := db.Query("SELECT id, count FROM schedule ORDER BY id")
if err != nil {
log.Fatal(err)
}
defer rows.Close()

var counts plotter.Values
var ides []string
for rows.Next() {
var id string
var count float64
if err := rows.Scan(&id, &count); err != nil {
log.Fatal(err)
}
ides = append(ides, id)
counts = append(counts, count)
}

p:= plot.New()
//if err != nil {
//log.Fatal(err)
//}

p.Title.Text = "Flights per Day"
p.X.Label.Text = "id"
p.Y.Label.Text = "Count"

bars, err := plotter.NewBarChart(counts, vg.Points(20))
if err != nil {
log.Fatal(err)
}
bars.LineStyle.Width = vg.Length(0)

p.Add(bars)
p.NominalX(ides...)

f, err := os.Create("flights.png")
if err != nil {
log.Fatal(err)
}
defer f.Close()

if err := p.Save(10*vg.Inch, 10*vg.Inch, "flights.png"); err != nil {
log.Fatal(err)
}
fmt.Println("Graph created successfully")
}