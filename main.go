package main

import (
	"log"
	"ping-pong-bot/bot"
	"ping-pong-bot/config"
)

func main() {
	err := config.ReadAllConfigs()

	if err != nil {
		log.Fatal("erro when read config")
	}

	bot.Start()

	<-make(chan struct{})

	return
}
