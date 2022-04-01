package commands

import (
	. "bot/config"
	"bot/personality"
	"bot/util"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

func init() {
	gamble := Command{
		Name: "gamble",
		Description: "gambles an amount of coins",
		Usage: Config.Prefix + "gamble <amount>",
		Args: 1,
		MinArgs: 1,
		Function: gamble,
	}
	
	Commands["gamble"] = &gamble
}

func gamble(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	n, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "`" + Commands["gamble"].Usage + "`")
		return
	}

	pocket, _, _, _ := util.GetCoins(m.Author.ID)
	if n > pocket {
		s.ChannelMessageSend(m.ChannelID, "lol ur too poor to gamble that much loser")
		return
	} else if n < 25 {
		s.ChannelMessageSend(m.ChannelID, "tf are u trying to gamble? dirt?")
		return
	}

	s1 := rand.NewSource(time.Now().UnixMilli())
	r1 := rand.New(s1)

	opguess := r1.Intn(100)

	s2 := rand.NewSource(time.Now().UnixMilli() - int64(r1.Intn(69420)))
	r2 := rand.New(s2)

	urguess := r2.Intn(100)

	embed := discordgo.MessageEmbed{}
	embed.Title = fmt.Sprintf("%s's gambling result", m.Author.Username)
	embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
		Name: "you",
		Value: fmt.Sprintf("%d", urguess),
		Inline: true,
	})
	
	embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
		Name: "opponent",
		Value: fmt.Sprintf("%d", opguess),
		Inline: true,
	})

	embed.Footer = &discordgo.MessageEmbedFooter{
		Text: personality.RandomFooter(),
		IconURL: m.Author.AvatarURL("128"),
	}

	var offset int64

	if opguess > urguess {
		if !util.UpdateCoins(m.Author.ID, -n) {
			fmt.Fprintln(os.Stderr, err)
			s.ChannelMessageSend(m.ChannelID, "server did a fucky wucky lol")
			return
		}
		offset = -n;
		embed.Description = fmt.Sprintf(
			"-%d coin <:shock:948766805911044116>", n,
		)
		embed.Color = 0xed050c
	} else {
		howmuchuwon := int64(urguess - opguess)
		amount := int64((float64(n) / 100) * float64(howmuchuwon)) + 10
		if !util.UpdateCoins(m.Author.ID, amount) {
			fmt.Fprintln(os.Stderr, err)
			s.ChannelMessageSend(m.ChannelID, "server did a fucky wucky lol")
			return
		}
		offset = amount
		embed.Description = fmt.Sprintf(
			"+%d coin <:socialcredit:948766805885849701>", amount,
		)
		embed.Color = 0x24c712
	}

	embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
		Name: "coins",
		Value: fmt.Sprintf("%d", pocket + offset),
	})
	
	s.ChannelMessageSendEmbed(m.ChannelID, &embed)
}
