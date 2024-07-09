package main

import (
	"log"

	"github.com/EmilioCliff/car-rental/api"
	"github.com/EmilioCliff/car-rental/utils"
)

func main() {
	sender := utils.NewGmailSender("Andy Safaris", "clifftest33@gmail.com", "tcvouixpfzgttlxq")

	server := api.NewServer(*sender)
	log.Println("starting server at local host: ", "8081")
	if err := server.Start("0.0.0.0:5000"); err != nil {
		log.Println(err)
	}
}
