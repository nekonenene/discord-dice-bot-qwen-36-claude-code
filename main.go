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
			replyUsage(s, m)
			return
		}

		result, err := dice.RollNotation(afterMention)
		if err != nil {
			replyError(s, m, err)
			return
		}

		replyRoll(s, m, result)
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

func replyUsage(s *discordgo.Session, m *discordgo.MessageCreate) {
	embed := &discordgo.MessageEmbed{
		Title: "🎲 ダイスの使い方",
		Fields: []*discordgo.MessageEmbedField{
			{Name: "表記形式", Value: "`NDM` — N個のM面ダイス"},
			{Name: "例", Value: "`4D6` (6面ダイスを4個)\n`2D100` (100面ダイスを2個)"},
			{Name: "範囲", Value: "個数: 1〜100 / 面数: 1〜100"},
		},
		Color: 0x3498DB,
	}
	sendEmbed(s, m, embed)
}

func replyError(s *discordgo.Session, m *discordgo.MessageCreate, err error) {
	embed := &discordgo.MessageEmbed{
		Title:       "❌ エラー",
		Description: fmt.Sprintf("Error: %s", err.Error()),
		Color:       0xDC3545,
	}
	sendEmbed(s, m, embed)
}

func replyRoll(s *discordgo.Session, m *discordgo.MessageCreate, result dice.Result) {
	values := make([]string, len(result.Values))
	for i, v := range result.Values {
		values[i] = fmt.Sprintf("%d", v)
	}
	embed := &discordgo.MessageEmbed{
		Title: fmt.Sprintf("🎲 %s を振りました！", diceNotation(result.Count, result.Sides)),
		Fields: []*discordgo.MessageEmbedField{
			{Name: "各目", Value: fmt.Sprintf("[%s]", strings.Join(values, ", ")), Inline: true},
			{Name: "合計", Value: fmt.Sprintf("**%d**", result.Total), Inline: true},
		},
		Color: 0x3DAE78,
	}
	sendEmbed(s, m, embed)
}

func diceNotation(count, sides int) string {
	return fmt.Sprintf("%dD%d", count, sides)
}

func sendEmbed(s *discordgo.Session, m *discordgo.MessageCreate, embed *discordgo.MessageEmbed) {
	if _, err := s.ChannelMessageSendEmbed(m.ChannelID, embed); err != nil {
		log.Printf("Failed to send embed: %v", err)
	}
}
