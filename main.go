package main

import (
	"bot/commands"
	. "bot/config"
	"bot/events"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	commands.InitHelp()
	dg, err := discordgo.New("Bot " + Config.Token)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	for _, event := range events.Events {
		dg.AddHandler(event)
	}
	
	dg.Identify.Intents = discordgo.IntentsAll

	err = dg.Open()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Println("bot is now running")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}
