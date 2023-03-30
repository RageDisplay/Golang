package main

import (
	"time"
)

func main() {
	for {

		create()

		clear()

		response()

		clear()

		graph()

		send()

		//time.Sleep(24 * time.Hour)

		time.Sleep(30 * time.Second)
	}
}
