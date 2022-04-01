package util

import (
	. "bot/config"
	. "bot/database"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

func GetCoins(id string) (int64, int64, int64, int64) {
	var pocket int64
	var bank int64
	var banksize int64
	var lastdaily int64

	err := DB.QueryRow("SELECT pocket, bank, banksize, lastdaily FROM coins WHERE id = ?", id).
		Scan(&pocket, &bank, &banksize, &lastdaily)
	if err == sql.ErrNoRows {
		if !InitCoins(id) {
			fmt.Fprintln(os.Stderr, err)
		}
	} else if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	return pocket, bank, banksize, lastdaily
}

func InitCoins(id string) bool {
	_, err := DB.Exec(fmt.Sprintf("INSERT INTO coins (id, pocket, bank, banksize, lastdaily) VALUES (%s, %d, %d, %d, %d)",
		id, 0, 0, Config.DefaultBank, 0))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return false
	}
	return true
}

func UpdateCoins(id string, offset int64) bool {
	_, err := DB.Exec(
		fmt.Sprintf("UPDATE coins SET pocket = pocket + %d WHERE id = %s", offset, id),
	)
	if err == sql.ErrNoRows {
		if !InitCoins(id) {
			fmt.Fprintln(os.Stderr, err)
			return false
		}
	} else if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return false
	}

	return true
}

func DailyCoins(id string) bool {
	t := time.Now().Unix()
	
	_, err := DB.Exec(
		fmt.Sprintf("UPDATE coins SET pocket = pocket + %d, lastdaily = %d WHERE id = %s",
			Config.DailySize, t, id))
	if err == sql.ErrNoRows {
		if !InitCoins(id) {
			fmt.Fprintln(os.Stderr, err)
			return false
		}
	} else	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return false
	}
	return true
}

func GetMember(s *discordgo.Session, guildid string, id string) *discordgo.Member {
	guild, err := s.State.Guild(guildid)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil
	}

	for _, m := range guild.Members {
		if m.User.ID == id {
			return m
		}
	}

	return nil
}

func GetUser(s *discordgo.Session, id string) *discordgo.User {
	for _, g := range s.State.Guilds {
		if m := GetMember(s, g.ID, id); m != nil {
			return m.User
		}
	}

	return nil
}
