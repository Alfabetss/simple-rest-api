package main

import (
	"fmt"

	"github.com/Alfabetss/simple-rest-api/config"
)

func main() {
	err := config.Connect()
	if err != nil {
		fmt.Printf("failed to connect : %s", err.Error())
		return
	}
	fmt.Println("connect success")
}
