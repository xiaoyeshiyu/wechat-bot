package main

import (
	"log"

	"wechat-bot/config"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	sugaredLogger := NewLog()
	server, err := initServer(sugaredLogger)
	if err != nil {
		sugaredLogger.Fatal(err)
	}

	err = server.Start()
	if err != nil {
		sugaredLogger.Fatal(err)
	}
}
