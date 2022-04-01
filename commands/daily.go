package commands

import (
	. "bot/config"
	"bot/personality"
	"bot/util"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func init() {
	daily := Command{
		Name: "daily",
		Description: "collects your daily amount of coins",
		Usage: Config.Prefix + "daily",
		Args: 0,
		MinArgs: 0,
		Function: daily,
	}
	
	Commands["daily"] = &daily
}

func daily(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	t := time.Now().Unix()
	var lastdaily int64
	_, _, _, lastdaily = util.GetCoins(m.Author.ID)

	embed := discordgo.MessageEmbed{}
	embed.Title = fmt.Sprintf("%s's daily", m.Author.Username)
	embed.Footer = &discordgo.MessageEmbedFooter{
		Text: personality.RandomFooter(),
		IconURL: m.Author.AvatarURL("128"),
	}
	
	if t - 86400 < lastdaily {
		embed.Description =
			fmt.Sprintf("u gotta wait like %s b4 u can daily again", fmtDuration((lastdaily + 86400) - t))
		
		embed.Color = 0xed050c
	} else {
		util.DailyCoins(m.Author.ID)

		embed.Description =
			fmt.Sprintf("%d coins r in ur bank now lol", Config.DailySize)
		
		embed.Color = 0x24c712
	}

	s.ChannelMessageSendEmbed(m.ChannelID, &embed)
}

func fmtDuration(d int64) string {
	m := d / 60
	h := 0

	for m > 59 {
		m -= 60
		h++
	}
	
	return fmt.Sprintf("%d hours and %d min", h, m)
}

