package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/gopulse/helpers"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	discord, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		return
	}

	discord.AddHandler(helpers.CreateMessage)

	discord.Identify.Intents = discordgo.IntentsGuildMessages

	err = discord.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
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
