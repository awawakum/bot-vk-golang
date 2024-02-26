package handler

import (
	"bot-vk/internal/service"
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-vk-api/vk"
	lp "github.com/go-vk-api/vk/longpoll/user"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) HandleMessage(update *lp.Update) *vk.RequestParams {

	switch data := update.Data.(type) {
	case *lp.NewMessage:

		h.services.ViewService.AddView(data.Text, fmt.Sprint(data.PeerID), "VK")

		switch data.Text {
		case "Начать":

			var message string

			h.services.UserService.AddUser(data.PeerID, "Vk_name", "Vk_lastname", "Vk_username", "VK")

			keys, err := h.services.KeyboardService.GetMainKeyboard()
			if err != nil {
				log.Panic(err)
			}

			message = "Срочно нужны деньги? 😳\nДорогие друзья! Добро пожаловать в бот финансовой помощи!\nЕсли вы попали в тяжёлую финансовую ситуацию или не знаете где взять деньги до ЗП, то милости просим! 🤗\nНиже представлены самые надежные и проверенные сервисы микрозаймов с высоким одобрением. На большинстве ресурсов есть акция \"Первый займ под 0%\". Каждый раз используйте новую организацию, чтобы избежать переплат. Не забывайте укладываться в беспроцентный период ⏳"

			b_keys, err := json.Marshal(keys)
			if err != nil {
				log.Panic(err)
			}

			obj := &vk.RequestParams{
				"peer_id":   data.PeerID,
				"message":   message,
				"keyboard":  string(b_keys),
				"random_id": 0,
			}
			return obj
		case "Займы":
			var message string

			keys, err := h.services.KeyboardService.GetCategoriesInlineKeyboard()
			if err != nil {
				log.Panic(err)
			}

			message = "Доступные страны: "

			b_keys, err := json.Marshal(keys)
			if err != nil {
				log.Panic(err)
			}

			obj := &vk.RequestParams{
				"peer_id":   data.PeerID,
				"message":   message,
				"keyboard":  string(b_keys),
				"random_id": 0,
			}
			return obj
		case "Россия":
			var message string

			keys, err := h.services.KeyboardService.GetProductInlineKeyboard(data.Text)
			if err != nil {
				log.Panic(err)
			}

			message = "Доступные ресурсы: "

			b_keys, err := json.Marshal(keys)
			if err != nil {
				log.Panic(err)
			}

			obj := &vk.RequestParams{
				"peer_id":   data.PeerID,
				"message":   message,
				"keyboard":  string(b_keys),
				"random_id": 0,
			}
			return obj
		case "Казахстан":
			var message string

			keys, err := h.services.KeyboardService.GetProductInlineKeyboard(data.Text)
			if err != nil {
				log.Panic(err)
			}

			message = "Доступные ресурсы: "

			b_keys, err := json.Marshal(keys)
			if err != nil {
				log.Panic(err)
			}

			obj := &vk.RequestParams{
				"peer_id":   data.PeerID,
				"message":   message,
				"keyboard":  string(b_keys),
				"random_id": 0,
			}
			return obj
		case "Кредистория":
			var message string

			keys, err := h.services.KeyboardService.GetHistoryInlineKeyboard()
			if err != nil {
				log.Panic(err)
			}

			message = "Нажмите на кнопку для получения кредистории"

			b_keys, err := json.Marshal(keys)
			if err != nil {
				log.Panic(err)
			}

			obj := &vk.RequestParams{
				"peer_id":   data.PeerID,
				"message":   message,
				"keyboard":  string(b_keys),
				"random_id": 0,
			}
			return obj
		case "Банкротство":
			var message string

			keys, err := h.services.KeyboardService.GetBankInlineKeyboard()
			if err != nil {
				log.Panic(err)
			}

			message = "Нажмите на кнопку для оформления банкротства"

			b_keys, err := json.Marshal(keys)
			if err != nil {
				log.Panic(err)
			}

			obj := &vk.RequestParams{
				"peer_id":   data.PeerID,
				"message":   message,
				"keyboard":  string(b_keys),
				"random_id": 0,
			}
			return obj
		case "ОСАГО":
			var message string

			keys, err := h.services.KeyboardService.GetOSAGOInlineKeyboard()
			if err != nil {
				log.Panic(err)
			}

			message = "Нажмите на кнопку для оформления ОСАГО"

			b_keys, err := json.Marshal(keys)
			if err != nil {
				log.Panic(err)
			}

			obj := &vk.RequestParams{
				"peer_id":   data.PeerID,
				"message":   message,
				"keyboard":  string(b_keys),
				"random_id": 0,
			}
			return obj
		}
	}
	return nil
}
