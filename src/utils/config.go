package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Telegram struct {
		AdminUid int
		BotToken string
		TgUrl    string
	}
}

func setEnvFromFile() {
	envFile, err := os.Open(".env")
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = envFile.Close()
	}()
	scanner := bufio.NewScanner(envFile)
	for scanner.Scan() {
		row := scanner.Text()
		spRow := strings.Split(row, " = ")
		if len(spRow) != 2 {
			continue
		}
		_ = os.Setenv(spRow[0], spRow[1])
	}
}

func InitConfig() *Config {
	setEnvFromFile()
	var cfg Config
	cfg.Telegram.AdminUid, _ = strconv.Atoi(os.Getenv("adminTgUid"))
	cfg.Telegram.BotToken = os.Getenv("botToken")
	cfg.Telegram.TgUrl = "https://api.telegram.org/bot"
	return &cfg
}
