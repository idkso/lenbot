package commands

import (
	. "bot/config"
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

var HelpMap map[*Command]*Command
var HelpCommands []*Command

func init() {
	HelpMap = make(map[*Command]*Command)
	
	help := Command{
		Name: "help",
		Description: "help with commands",
		Usage: Config.Prefix + "help <command>",
		Args: 1,
		MinArgs: 0,
		Function: help,
	}
	
	Commands["help"] = &help
}

func InitHelp() {
	for _, v := range Commands {
		HelpMap[v] = v
	}

	for _, v := range HelpMap {
		HelpCommands = append(HelpCommands, v)
		fmt.Println("command", v.Name, "ready to go")
	}

	fmt.Println("help menu ready")
}

func help(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) > 0 {
		_, err := strconv.Atoi(args[0])
		if err != nil {
			helpCmd(s, m, args[0])
			return
		}
	}
	
	arg := 0
	embed := discordgo.MessageEmbed{}
	embed.Title = "help menu"
	embed.Footer = &discordgo.MessageEmbedFooter{
		Text: fmt.Sprintf("%d commands", len(HelpCommands)),
	}

	page := 1
	beginpage := 0
	if len(args) > 0 {
		n, err := strconv.Atoi(args[arg])
		if err == nil {
			arg++
			page = n
		}
	}
	
	var cmdslen int

	if page * 10 > len(HelpCommands) {
		page = 1
	}

	page *= 10
	beginpage = page - 10

	if len(HelpCommands) > 10 {
		cmdslen = 10
	} else {
		cmdslen = len(HelpCommands)
	}

	if page > len(HelpCommands) {
		page = cmdslen
		beginpage = 0
	}

	for _, cmd := range HelpCommands[beginpage:page] {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name: cmd.Name,
			Value: fmt.Sprintf("%s\n`%s`", cmd.Description, cmd.Usage),
			Inline: false,
		})
	}

	s.ChannelMessageSendEmbed(m.ChannelID, &embed)
}

func helpCmd(s *discordgo.Session, m *discordgo.MessageCreate, cmd string) {
	command, ok := Commands[cmd]
	if !ok {
		s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
			Content: fmt.Sprintf("invalid command '%s'", cmd),
			AllowedMentions: &discordgo.MessageAllowedMentions{
				Parse: []discordgo.AllowedMentionType{"users"},
			},
		})
		return
	}
	embed := discordgo.MessageEmbed{}
	embed.Title = command.Name
	embed.Description = command.Description

	embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
		Name: "usage",
		Value: "`" + command.Usage + "`",
		Inline: false,
	})

	embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
		Name: "args",
		Value: fmt.Sprintf("%d", command.Args),
		Inline: true,
	})

	embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
		Name: "minimum args",
		Value: fmt.Sprintf("%d", command.MinArgs),
		Inline: true,
	})

	s.ChannelMessageSendEmbed(m.ChannelID, &embed)
}
