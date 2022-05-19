package bot

import (
	"fmt"
	"log"
	"ping-pong-bot/config"

	"github.com/bwmarrin/discordgo"
)

func Start() {
	discordBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		log.Fatal("Error while creating a discord session")
	}

	discordBot.AddHandler(messageCreate)

	discordBot.Identify.Intents = discordgo.IntentGuildMessages

	err = discordBot.Open()

	if err != nil {
		log.Fatal("error when connect ", err)
	}

	fmt.Println("Bot is running...")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == config.BotPrefix+"pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}

	if m.Content == config.BotPrefix+"ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}
}
