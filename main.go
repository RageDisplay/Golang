package main

import (
	"time"
)

func main() {

	for {
		delpic()

		//time.Sleep(5 * time.Second)

		create()

		//time.Sleep(5 * time.Second)

		clear()

		//time.Sleep(5 * time.Second)

		response()

		//time.Sleep(5 * time.Second)

		clear()

		//time.Sleep(5 * time.Second)

		graph()

		//time.Sleep(5 * time.Second)

		send()

		time.Sleep(60 * time.Minute)

		//time.Sleep(5 * time.Second)
	}
}
