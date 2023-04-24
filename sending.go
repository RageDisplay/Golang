package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func send() error {
	// URL адрес для отправки POST-запроса с изображением
	url := "https://api.imgbb.com/1/upload?key=4759b9b02b8b77232df840dee6d7991d"

	// имя (ключ) файла на сервере
	fileKey := "image"

	// имя файла для отправки
	filename := "flights.png"

	// открытие файла
	file, err := os.Open(filename)
	if err != nil {
		clearlast()
		fmt.Println(err)
		return nil
		//log.Fatal(err)
	}
	defer file.Close()

	// создание тела запроса
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// добавление файла в тело запроса
	part, err := writer.CreateFormFile(fileKey, filename)
	if err != nil {
		clearlast()
		fmt.Println(err)
		return nil
		//log.Fatal(err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		fmt.Println(err)
		return nil
		//log.Fatal(err)
	}

	// закрытие тела запроса
	err = writer.Close()
	if err != nil {
		fmt.Println(err)
		return nil
		//log.Fatal(err)
	}

	// создание POST-запроса с телом запроса
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		fmt.Println(err)
		return nil
		//log.Fatal(err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// отправка запроса
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		clearlast()
		fmt.Println(err)
		return nil
		//log.Fatal(err)
	} else {
		fmt.Println("График успешно отправлен")
	}

	defer res.Body.Close()

	//resBody, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println(string(resBody))
	return nil
}

func delpic() error {
	err := os.Remove("flights.png")
	if err != nil {
		//fmt.Println(err)
		return nil
	}
	return nil
}
