package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"

	"github.com/nekonenene/discord-dice-bot-qwen-claude-code/dice"
)

func main() {
	token := os.Getenv("DISCORD_BOT_TOKEN")
	if token == "" {
		log.Fatal("DISCORD_BOT_TOKEN environment variable is required")
	}

	s, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("Invalid bot credentials: %v", err)
	}

	s.AddHandler(func(_ *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}
		if !strings.HasPrefix(m.Content, "<@") {
			return
		}

		mentionEnd := strings.Index(m.Content, ">")
		if mentionEnd == -1 {
			return
		}

		afterMention := strings.TrimSpace(m.Content[mentionEnd+1:])
		if afterMention == "" {
			reply(s, m, "Usage: `<@bot> 4D6`")
			return
		}

		result, err := dice.RollNotation(afterMention)
		if err != nil {
			reply(s, m, fmt.Sprintf("Error: %s", err.Error()))
			return
		}

		reply(s, m, result.String())
	})

	if err := s.Open(); err != nil {
		log.Fatalf("Failed to open session: %v", err)
	}

	log.Println("Bot is online!")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	s.Close()
	log.Println("Shutting down...")
}

func reply(s *discordgo.Session, m *discordgo.MessageCreate, content string) {
	if _, err := s.ChannelMessageSend(m.ChannelID, content); err != nil {
		log.Printf("Failed to send reply: %v", err)
	}
}
