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

var dicePattern = regexp.MustCompile(`^(\d+)d(\d+)$`)

func parseDice(s string) (count int, faces int, ok bool) {
	s = strings.TrimSpace(s)
	m := dicePattern.FindStringSubmatch(s)
	if m == nil {
		return 0, 0, false
	}
	count, err1 := strconv.Atoi(m[1])
	faces, err2 := strconv.Atoi(m[2])
	if err1 != nil || err2 != nil {
		return 0, 0, false
	}
	return count, faces, true
}

func rollDice(count, faces int) (sum int, results []int) {
	results = make([]int, count)
	sum = 0
	for i := 0; i < count; i++ {
		v := rand.Intn(faces) + 1
		results[i] = v
		sum += v
	}
	return sum, results
}

func main() {
	botToken := os.Getenv("DISCORD_BOT_TOKEN")
	if botToken == "" {
		log.Fatal("DISCORD_BOT_TOKEN environment variable is required")
	}

	session, err := discordgo.New("Bot " + botToken)
	if err != nil {
		log.Fatal(err)
	}

	session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.Bot {
			return
		}

		mentioned := false
		for _, u := range m.Mentions {
			if u.ID == s.State.User.ID {
				mentioned = true
				break
			}
		}
		if !mentioned {
			return
		}

		content := m.ContentWithMentionsReplaced()

		count, faces, ok := parseDice(content)
		if !ok {
			s.ChannelMessageSend(m.ChannelID, "ダイスを指定してください。例: `4D6`")
			return
		}

		if count < 1 || count > 100 {
			s.ChannelMessageSend(m.ChannelID, "ダイスの個数は1~100を指定してください。")
			return
		}
		if faces < 1 || faces > 100 {
			s.ChannelMessageSend(m.ChannelID, "ダイスの面数は1~100を指定してください。")
			return
		}

		sum, results := rollDice(count, faces)

		if count <= 10 {
			resp := fmt.Sprintf("%d (%s)", sum, strings.Join(toStringSlice(results), ", "))
			s.ChannelMessageSend(m.ChannelID, resp)
		} else {
			s.ChannelMessageSend(m.ChannelID, strconv.Itoa(sum))
		}
	})

	err = session.Open()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Bot is ready. Listening for mentions...")
	select {}
}

func toStringSlice(nums []int) []string {
	strs := make([]string, len(nums))
	for i, n := range nums {
		strs[i] = strconv.Itoa(n)
	}
	return strs
}
