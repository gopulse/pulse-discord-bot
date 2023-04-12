package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	discord, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		return
	}

	discord.AddHandler(messageCreate)

	err = discord.Open()
	if err != nil {
		panic(err)
	}

	fmt.Println("Bot is running")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	err = discord.Close()
	if err != nil {
		return
	}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.ChannelID != "1095489906450628690" {
		return
	}

	// Ignore messages from bots
	if m.Author.Bot {
		return
	}

	fmt.Printf("[%s] %s: %s\n", m.ChannelID, m.Author.Username, m.Content)
}
