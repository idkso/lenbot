package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

type ConfigStruct struct {
	Token string `toml:"token"`
	Prefix string `toml:"prefix"`
	DB string `toml:"db"`
	DailySize int64 `toml:"daily"`
	DefaultBank int64 `toml:"default_bank"`
}

var Config ConfigStruct

func init() {
	file, err := ioutil.ReadFile("config.toml")
	if err != nil {
		fmt.Fprintln(os.Stderr, "error opening or reading config.toml")
		os.Exit(1)
	}

	err = toml.Unmarshal(file, &Config)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error unmarshaling config.toml")
		os.Exit(1)
	}
}
