package helpers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"strings"
)

func CreateMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	channel, err := s.Channel(m.ChannelID)
	if err != nil {
		fmt.Printf("Failed to get channel: %v\n", err)
		return
	}
	if channel.ParentID != os.Getenv("DISCORD_HELP_FORUM_CHANNEL_ID") {
		return
	}

	if strings.Contains(m.Content, "```go") && strings.Contains(m.Content, "```") {
		code := strings.TrimPrefix(m.Content, "```go")
		code = strings.ReplaceAll(code, "```", "")

		if strings.Contains(m.Content, "!RunMyCode") {
			code = strings.ReplaceAll(code, "!RunMyCode", "")
			compiledCode, err := compileAndRun(code)
			if err == nil {
				message := fmt.Sprintf("<@%s>, here's the output of your code:\n```%s```", m.Author.ID, compiledCode)
				_, err := s.ChannelMessageSend(m.ChannelID, message)
				if err != nil {
					fmt.Println("error sending message,", err)
					return
				}
			}
		}
	}
}
