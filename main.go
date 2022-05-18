package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	token = "OTU4ODY3ODI5MTU3OTI0ODc0.Gbz7QY.ppZMr-J1Afbb2LEBEWCzCOBNpISY-6hmTgslvM"
)

func main() {

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("erro creating Discord session", err)
		return
	}

	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.IntentGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("error when connect", err)
		return
	}

	fmt.Println("The Bot is Running...")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
