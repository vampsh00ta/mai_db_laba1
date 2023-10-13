package keyboard

import (
	"github.com/go-telegram/bot/models"
)

func Gai() *models.InlineKeyboardMarkup {
	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Случилось дтп", CallbackData: "test"},
			}, {
				{Text: "Зарегистрировать автомобиль", CallbackData: "gai_register_auto"},
			},
			{
				{Text: "Проверить автомобиль", CallbackData: "gai_check_auto"},
			},
			{
				{Text: "Назад", CallbackData: "back_main"},
			},
		},
	}
	return kb
}
