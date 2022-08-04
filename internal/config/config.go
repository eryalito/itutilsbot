package config

import (
	"flag"
	"log"
	"os"
)

type ArgFlags struct {
	Config struct {
		AuthPath    string
		DisableAuth bool
		Token       string
	}
}

func (af *ArgFlags) Init() {
	flag.StringVar(&af.Config.AuthPath, "a", "auth", "Path to auth policy folder")
	flag.BoolVar(&af.Config.DisableAuth, "d", false, "Disable auth policies load")
	flag.StringVar(&af.Config.Token, "t", "", "Bot token. (API_TOKEN environment variable)")
}

func (af *ArgFlags) Parse() {
	flag.Parse()

	if af.Config.Token == "" {
		env_token, ok := os.LookupEnv("API_TOKEN")
		if !ok {
			log.Fatal("Bot token not provided")
		} else {
			af.Config.Token = env_token
		}
	}
}
