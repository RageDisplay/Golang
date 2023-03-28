package main

import (
	"time"
)

func main() {
	for {
		response()

		clear()

		graph()

		time.Sleep(24 * time.Hour)
	}
}
