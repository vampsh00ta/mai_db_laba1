package keyboard

import (
	"github.com/go-telegram/bot/models"
	"strconv"
)

const (
	RegDtpKey = iota
	RegVehicleKey
	_
	BackMainKey
)

func Gai() *models.InlineKeyboardMarkup {
	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Случилось дтп", CallbackData: "gai_" + strconv.Itoa(RegDtpKey)},
			}, {
				{Text: "Зарегистрировать автомобиль", CallbackData: "gai_" + strconv.Itoa(RegVehicleKey)},
			},
			{
				{Text: "Проверить автомобиль", CallbackData: "gai_" + strconv.Itoa(CheckVehicleKey)},
			},
			{
				{Text: "Назад", CallbackData: "back_" + strconv.Itoa(BackToMainKey)},
			},
		},
	}
	return kb
}
