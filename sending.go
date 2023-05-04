package main

import (
	"bytes"
	"io"
	"log"
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
		return err
	}
	defer file.Close()

	// создание тела запроса
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// добавление файла в тело запроса
	part, err := writer.CreateFormFile(fileKey, filename)
	if err != nil {
		clearlast()
		return err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}

	// закрытие тела запроса
	err = writer.Close()
	if err != nil {
		return err

	}

	// создание POST-запроса с телом запроса
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// отправка запроса
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		clearlast()
		return err
	} else {
		log.Println("График успешно отправлен")
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
		return nil
	}
	return nil
}
