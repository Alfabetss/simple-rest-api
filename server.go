package main

import (
	"fmt"

	"github.com/Alfabetss/simple-rest-api/config"
)

func main() {
	db, err := config.Connect()
	defer db.Close()

	if err != nil {
		fmt.Printf("failed to connect : %s", err.Error())
		return
	}
	fmt.Println("connect success")
}
