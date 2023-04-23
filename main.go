package main

import (
	"os"
	"os/exec"
	"time"
)

func main() {
	for {

		create()

		clear()

		response()

		clear()

		graph()

		//time.Sleep(10 * time.Second)

		send()

		//time.Sleep(10 * time.Second)

		time.Sleep(60 * time.Minute)

		//time.Sleep(2 * time.Second)
	}
}

func restart() {
	//cmd := exec.Command("cmd.exe", "/C", "autorun.bat")    //Windows
	cmd := exec.Command("make") //Linux
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
