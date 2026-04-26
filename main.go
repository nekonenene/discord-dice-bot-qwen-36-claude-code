package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"discord-dice-bot/dice"

	"github.com/bwmarrin/discordgo"
)

func main() {
	token := os.Getenv("DISCORD_BOT_TOKEN")
	if token == "" {
		log.Fatal("DISCORD_BOT_TOKEN environment variable is required")
	}

	s, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}

	s.AddHandler(messageHandler)

	if err := s.Open(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Bot is online. Press Ctrl+C to stop.")
	<-make(chan struct{})
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	content := strings.TrimSpace(m.Content)
	if content == "" {
		return
	}

	result, err := dice.ParseAndRoll(content)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("エラー: %v", err))
		return
	}

	s.ChannelMessageSend(m.ChannelID, result.Format())
}
