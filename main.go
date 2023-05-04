package main

import (
	"log"
	"time"
)

func main() {

	for {
		delpic()

		//time.Sleep(5 * time.Second)
	mark1:
		err := create()
		if err != nil {
			log.Print("Ошибка в создании таблицы в БД: ", err)
			goto mark1
		}

		//time.Sleep(5 * time.Second)

		err = clear()
		if err != nil {
			log.Print("Ошибка в очистке старых данных БД: ", err)
			goto mark1
		}

		//time.Sleep(5 * time.Second)
	mark2:
		err = response()
		if err != nil {
			log.Print("Ошибка в запросе", err)
			goto mark2
		}
		//time.Sleep(5 * time.Second)
	mark3:
		err = clear()
		if err != nil {
			log.Print("Ошибка в очистке старых данных БД: ", err)
			goto mark1
		}

		//time.Sleep(5 * time.Second)

		err = graph()
		if err != nil {
			log.Print("Ошибка в создании графика: ", err)
			goto mark3
		}

		//time.Sleep(5 * time.Second)

		err = send()
		if err != nil {
			log.Print("Ошибка в отправке графика: ", err)
			goto mark3
		}

		time.Sleep(60 * time.Minute)

		//time.Sleep(5 * time.Second)
	}
}
