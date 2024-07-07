package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/daniil49926/tg_job_notifier/src/utils"
	"log"
	"net/http"
)

func SendMessage(mess string, cfg utils.Config) error {
	data := map[string]interface{}{
		"chat_id": cfg.Telegram.AdminUid,
		"text":    mess,
	}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	reader := bytes.NewBuffer(dataBytes)
	resp, err := http.Post(
		fmt.Sprintf("%s%s/sendMessage", cfg.Telegram.TgUrl, cfg.Telegram.BotToken),
		"application/json",
		reader,
	)
	if err != nil {
		return err
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	return nil
}
