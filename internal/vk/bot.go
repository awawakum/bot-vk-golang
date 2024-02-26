package vk

import (
	"bot-vk/internal/handler"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-vk-api/vk"
	lp "github.com/go-vk-api/vk/longpoll/user"
	"github.com/joho/godotenv"
)

type Bot struct {
	bot     *vk.Client
	handler handler.Handler
}

func NewBot(tgString string, handler *handler.Handler) (*Bot, error) {
	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	}

	bot, err := vk.NewClientWithOptions(
		vk.WithToken(os.Getenv("TOKEN")),
		vk.WithHTTPClient(&http.Client{Timeout: time.Second * 30}),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize bot API: %w", err)
	}
	return &Bot{
		bot:     bot,
		handler: *handler,
	}, nil
}

func (b *Bot) Run() error {

	longpoll, err := lp.NewWithOptions(b.bot, lp.WithMode(lp.ReceiveAttachments))
	if err != nil {
		log.Panic(err)
	}

	stream, err := longpoll.GetUpdatesStream(0)
	if err != nil {
		log.Panic(err)
	}

	for {
		select {

		case update, ok := <-stream.Updates:
			if !ok {
				return err
			}

			params := b.handler.HandleMessage(update)
			if params != nil {

				var sentMessageID int64

				if err = b.bot.CallMethod("messages.send", *params, &sentMessageID); err != nil {
					log.Panic(err)
				}
			}

		case err, ok := <-stream.Errors:
			if ok {
				//	stream.Stop()
				log.Panic(err)
			}
		}
	}
}
