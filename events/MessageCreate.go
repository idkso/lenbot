package events

import (
	"bot/commands"
	. "bot/config"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var Prefix string
var Hooks []func(*discordgo.Session, *discordgo.MessageCreate)

func init() {
	Events = append(Events, MessageCreate)
}

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	for _, v := range Hooks {
		go v(s, m)
	}

	if !strings.HasPrefix(m.Content, Config.Prefix) {
		return
	}
	
	str := strings.TrimPrefix(m.Content, Config.Prefix)
	args := strings.Split(str, " ")
	cmd := args[0]

	args = args[1:]

	commands.RunCmd(s, m, cmd, args)
}
