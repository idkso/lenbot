package events

import (
	"bot/personality"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

const (
	guildid        string = "959268017055862844"
	welcomechannel string = "959268019123671060"
)

func init() {
	Events = append(Events, GuildMemberAdd)
}

func GuildMemberAdd(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	if m.GuildID != guildid {
		return
	}

	var role string = "959269364039835660"

	if m.User.Bot {
		role = "959269108778688582"
	}
	
	s.GuildMemberRoleAdd(guildid, m.User.ID, role)

	s.ChannelMessageSend("959268019123671060", fmt.Sprintf(personality.RandomWelcome(), m.User.ID))
}
