package keyboard

import (
	"github.com/go-telegram/bot/models"
)

func Main() *models.InlineKeyboardMarkup {
	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "ГАИ", CallbackData: "main_gai"},
			}, {
				{Text: "Сотрудник ГИБДД", CallbackData: "main_gaishnik"},
			},
		},
	}
	return kb
}
