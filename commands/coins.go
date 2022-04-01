package commands

import (
	. "bot/config"
	"bot/personality"
	"bot/util"
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func init() {
	coins := Command{
		Name: "coins",
		Description: "shows a users coins",
		Usage: Config.Prefix + "coins <@user|userid>",
		Args: 1,
		MinArgs: 0,
		Function: coins,
	}
	
	Commands["coins"] = &coins
	Commands["balance"] = &coins
	Commands["bal"] = &coins
}

func coins(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	var id string
	if len(args) > 0 {
		if _, err := strconv.ParseInt(args[0], 10, 64); err != nil {
			id = m.Mentions[0].ID
		} else {
			id = args[0]
		}
	} else {
		id = m.Author.ID
	}

	pocket, bank, banksize, _ := util.GetCoins(m.Author.ID)

	user := util.GetMember(s, m.GuildID, id).User

	embed := discordgo.MessageEmbed{}
	embed.Title = fmt.Sprintf("%s's coins", user.Username)
	embed.Footer = &discordgo.MessageEmbedFooter{
		Text: personality.RandomFooter(),
		IconURL: user.AvatarURL("128"),
	}
	embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
		Name: "pocket",
		Value: fmt.Sprintf("%d", pocket),
		Inline: true,
	})
	
	embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
		Name: "bank",
		Value: fmt.Sprintf("%d/%d", bank, banksize),
		Inline: true,
	})

	s.ChannelMessageSendEmbed(m.ChannelID, &embed)
}
