package service

import (
	"bot-vk/internal/database"
	"bot-vk/internal/models"
	"encoding/json"
	"log"
	"os"
	"strings"
)

type KeyboardServiceImpl struct {
	keyboardData *database.DataBase
}

func NewKeyboardServiceImpl(database *database.DataBase) *KeyboardServiceImpl {
	return &KeyboardServiceImpl{keyboardData: database}
}

func GetJsonByteFromFile(name string) []byte {

	json_data, err := os.ReadFile(name)
	if err != nil {
		log.Panic(err)
	}
	return json_data

}

func (impl *KeyboardServiceImpl) GetMainKeyboard() (models.Keyboard, error) {
	json_bytes := GetJsonByteFromFile("internal\\keyboards\\main.json")

	var keys models.Keyboard
	if err := json.Unmarshal([]byte(json_bytes), &keys); err != nil {
		log.Panic(err)
	}
	return keys, nil
}

func (impl *KeyboardServiceImpl) GetCategoriesInlineKeyboard() (models.Keyboard, error) {

	var keys models.Keyboard
	var buttons [][]models.Button

	keys.OneTime = false
	keys.Inline = true

	categories := impl.keyboardData.CategoryData.GetCategories()

	for i := 0; i < len(categories); i += 1 {

		var button_row []models.Button
		var button models.Button

		button.Action.Label = categories[i].CategoryTitle
		button.Action.Payload = ""
		button.Action.Type = "text"
		button.Color = "primary"

		button_row = append(button_row, button)

		buttons = append(buttons, button_row)

	}

	keys.Buttons = append(keys.Buttons, buttons...)

	return keys, nil

}

func (impl *KeyboardServiceImpl) GetProductInlineKeyboard(category string) (models.LinkKeyboard, error) {

	var keys models.LinkKeyboard
	var buttons [][]models.LinkButton

	keys.OneTime = false
	keys.Inline = true

	products := impl.keyboardData.PorductData.GetProductsWhereCategory(category)

	var button_row []models.LinkButton

	for i := 0; i < len(products); i += 1 {

		var button models.LinkButton

		urlIsCorrect := strings.Contains(products[i].ProductBody, "2go")

		if urlIsCorrect {
			button.Action = models.LinkAction{
				Label:   products[i].ProductTitle,
				Payload: "",
				Link:    products[i].ProductBody,
				Type:    "open_link",
			}

			button_row = append(button_row, button)
		}

		if (len(button_row)%2 == 0) || (len(button_row)%2 != 0 && i+1 == len(products)) {
			buttons = append(buttons, button_row)
			button_row = nil
		}

		if i == 12 {
			keys.Buttons = append(keys.Buttons, buttons...)

			return keys, nil
		}

	}

	keys.Buttons = append(keys.Buttons, buttons...)

	return keys, nil

}

func (impl *KeyboardServiceImpl) GetOSAGOInlineKeyboard() (models.LinkKeyboard, error) {
	var keys models.LinkKeyboard
	var buttons [][]models.LinkButton

	keys.OneTime = false
	keys.Inline = true

	var button_row []models.LinkButton

	var button models.LinkButton

	button.Action = models.LinkAction{
		Label:   "OSAGO",
		Payload: "",
		Link:    "https://my.saleads.pro/s/2Jiig",
		Type:    "open_link",
	}

	button_row = append(button_row, button)

	buttons = append(buttons, button_row)

	keys.Buttons = append(keys.Buttons, buttons...)

	return keys, nil
}

func (impl *KeyboardServiceImpl) GetHistoryInlineKeyboard() (models.LinkKeyboard, error) {
	var keys models.LinkKeyboard
	var buttons [][]models.LinkButton

	keys.OneTime = false
	keys.Inline = true

	var button_row []models.LinkButton

	var button models.LinkButton

	button.Action = models.LinkAction{
		Label:   "Кредистория",
		Payload: "",
		Link:    "https://my.saleads.pro/s/6zJn8",
		Type:    "open_link",
	}

	button_row = append(button_row, button)

	buttons = append(buttons, button_row)

	keys.Buttons = append(keys.Buttons, buttons...)

	return keys, nil
}

func (impl *KeyboardServiceImpl) GetBankInlineKeyboard() (models.LinkKeyboard, error) {
	var keys models.LinkKeyboard
	var buttons [][]models.LinkButton

	keys.OneTime = false
	keys.Inline = true

	var button_row []models.LinkButton

	var button models.LinkButton

	button.Action = models.LinkAction{
		Label:   "Банкротство",
		Payload: "",
		Link:    "https://my.saleads.pro/s/kdzs0",
		Type:    "open_link",
	}

	button_row = append(button_row, button)

	buttons = append(buttons, button_row)

	keys.Buttons = append(keys.Buttons, buttons...)

	return keys, nil
}
