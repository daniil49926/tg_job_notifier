package main

import (
	"github.com/daniil49926/tg_job_notifier/src/logger"
	"github.com/daniil49926/tg_job_notifier/src/telegram"
	"github.com/daniil49926/tg_job_notifier/src/utils"
)

func main() {
	mLogger := logger.InitLogger()
	mLogger.Println("Start job notifier")
	cfg := utils.InitConfig()
	err := telegram.SendMessage("Hi", *cfg)
	if err != nil {
		mLogger.Println(err)
	}
}
