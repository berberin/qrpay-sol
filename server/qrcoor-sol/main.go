package main

import (
	"log"
	"qrcoor-sol/config"
	"qrcoor-sol/librarian"
	"qrcoor-sol/server"
)

func main() {
	config.LoadConfig("config.json")
	librarian.Init()
	err := server.RunServer()
	if err != nil {
		log.Println(err)
	}
}
