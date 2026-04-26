package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var dicePattern = regexp.MustCompile(`^(\d+)D(\d+)$`)

func parseDice(s string) (count int, faces int, ok bool) {
	m := dicePattern.FindStringSubmatch(strings.ToUpper(s))
	if m == nil {
		return 0, 0, false
	}
	count, err1 := strconv.Atoi(m[1])
	faces, err2 := strconv.Atoi(m[2])
	if err1 != nil || err2 != nil {
		return 0, 0, false
	}
	if count < 1 || count > 100 || faces < 1 || faces > 100 {
		return 0, 0, false
	}
	return count, faces, true
}

func rollDice(count, faces int) []int {
	results := make([]int, count)
	for i := range results {
		results[i] = rand.Intn(faces) + 1
	}
	return results
}

func formatResult(count, faces int, results []int) string {
	sum := 0
	parts := make([]string, count)
	for i, v := range results {
		parts[i] = strconv.Itoa(v)
		sum += v
	}
	return fmt.Sprintf("%dD%d: [%s] (合計: %d)", count, faces, strings.Join(parts, ", "), sum)
}

func main() {
	token := os.Getenv("DISCORD_BOT_TOKEN")
	if token == "" {
		log.Fatal("DISCORD_BOT_TOKEN environment variable is required")
	}

	session, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}

	session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == session.State.User.ID {
			return
		}

		content := strings.TrimSpace(m.Content)

		botMention := "<@" + session.State.User.ID + ">"
		botMentionWithPrefix := "<@!" + session.State.User.ID + ">"

		var diceStr string
		switch {
		case strings.HasPrefix(content, botMention+" "):
			diceStr = strings.TrimPrefix(content, botMention+" ")
		case strings.HasPrefix(content, botMentionWithPrefix+" "):
			diceStr = strings.TrimPrefix(content, botMentionWithPrefix+" ")
		default:
			return
		}

		diceStr = strings.TrimSpace(diceStr)
		if diceStr == "" {
			return
		}

		count, faces, ok := parseDice(diceStr)
		if !ok {
			sendEmbed(s, m.ChannelID, "エラー", "ダイス指定が不正です。\n有効な形式: ND M (例: 4D6)\n面数: 1~100, 個数: 1~100")
			return
		}

		results := rollDice(count, faces)
		msg := formatResult(count, faces, results)
		sendEmbed(s, m.ChannelID, fmt.Sprintf("%dD%d を振りました", count, faces), msg)
	})

	err = session.Open()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Bot online! Waiting for messages...")
	select {}
}

func sendEmbed(s *discordgo.Session, channelID, title, message string) {
	color := rand.Intn(0xFFFFFF)
	embed := &discordgo.MessageEmbed{
		Title:       title,
		Description: message,
		Color:       color,
	}
	_, err := s.ChannelMessageSendEmbed(channelID, embed)
	if err != nil {
		log.Printf("Failed to send embed: %v", err)
	}
}
