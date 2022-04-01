package commands

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Command struct {
	Name string
	Description string
	Usage string
	Args int
	MinArgs int
	Function func(*discordgo.Session, *discordgo.MessageCreate, []string)
}

var Commands map[string]*Command

func init() {
	Commands = make(map[string]*Command)
}

func RunCmd(s *discordgo.Session, m *discordgo.MessageCreate, cmd string, args []string) {
	cmdrun := false

	cmdlower := strings.ToLower(cmd)

	for k, v := range Commands {
		if k == cmdlower {
			if len(args) >= v.MinArgs {
				v.Function(s, m, args)
				cmdrun = true
			} else {
				s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
					Content: "`" + v.Usage + "`",
					AllowedMentions: &discordgo.MessageAllowedMentions{
						Parse: []discordgo.AllowedMentionType{"users"},
					},
				})
			}
		}
	}

	if !cmdrun {
		s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
			Content: fmt.Sprintf("invalid command '%s'", cmd),
			AllowedMentions: &discordgo.MessageAllowedMentions{
				Parse: []discordgo.AllowedMentionType{"users"},
			},
		})
	}
}
