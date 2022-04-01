package commands

import (
	. "bot/config"
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"

	"github.com/bwmarrin/discordgo"
)

func init() {
	cmd := Command{
		Name: "color",
		Description: "sends a color image from a hex code",
		Usage: Config.Prefix + "color #<hex>",
		Args: 1,
		MinArgs: 1,
		Function: colorcmd,
	}
	
	Commands["color"] = &cmd
}

func colorcmd(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	width := 500
	height := 500
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	thecolor, err := ParseHexColor(args[0])

	if err != nil {
		fmt.Println("error parsing hex color:", err.Error())
		return
	}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, thecolor)
		}
	}

	var b []byte
	lol := bytes.NewBuffer(b)

	png.Encode(lol, img)
	s.ChannelFileSend(m.ChannelID, "color.png", lol)
	lol.Reset()
}

// stolen from stackoverflow
var errInvalidFormat = errors.New("invalid format")

func ParseHexColor(s string) (c color.RGBA, err error) {
	c.A = 0xff

	if s[0] != '#' {
		return c, errInvalidFormat
	}

	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		err = errInvalidFormat
		return 0
	}

	switch len(s) {
	case 7:
		c.R = hexToByte(s[1])<<4 + hexToByte(s[2])
		c.G = hexToByte(s[3])<<4 + hexToByte(s[4])
		c.B = hexToByte(s[5])<<4 + hexToByte(s[6])
	case 4:
		c.R = hexToByte(s[1]) * 17
		c.G = hexToByte(s[2]) * 17
		c.B = hexToByte(s[3]) * 17
	default:
		err = errInvalidFormat
	}
	return
}
