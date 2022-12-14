package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"yt_music_bot_go/pkg/ytfiles"
)

const commandStart = "start"

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Я не знаю такой команды х(")
	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(&msg)
	default:
		return b.handleUnknownCommand(&msg)
	}
}

func (b *Bot) handleMessage(message *tgbotapi.Message) {

	infos, err := ytfiles.PrepareForSending(message.From.ID, message.Text)
	if err != nil {
		log.Fatal(err)
	}
	defer ytfiles.DeleteDir(message.From.ID)

	if len(*infos) == 1 {
		msg := tgbotapi.NewAudioUpload(message.Chat.ID, (*infos)[0].FullPath)
		_, err := b.bot.Send(msg)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		for _, file := range *infos {
			msg := tgbotapi.NewAudioUpload(message.Chat.ID, file.FullPath)
			_, err := b.bot.Send(msg)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func (b *Bot) handleStartCommand(msg *tgbotapi.MessageConfig) error {
	msg.Text = "Ну шо, погнали!11"
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleUnknownCommand(msg *tgbotapi.MessageConfig) error {
	_, err := b.bot.Send(msg)
	return err
}
