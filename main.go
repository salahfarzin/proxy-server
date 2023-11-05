package main

import (
	"github.com/joho/godotenv"
	server "github.com/salahfarzin/api/src"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	server.Start()
}
