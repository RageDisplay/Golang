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

		time.Sleep(60 * time.Minute)

		//time.Sleep(2 * time.Second)
	}
}
