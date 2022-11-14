package main

import (
	"countries-api/http/gin/server"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err.Error())
	}
	s := server.NewServer()
	s.Start()
}
