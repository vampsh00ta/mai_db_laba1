package keyboard

import (
	"github.com/go-telegram/bot/models"
	"strconv"
)

const (
	GaiKey = iota
	GaishnikKey
)
const (
	BackToMainKey = iota
	BackToGaiKey
	BackToGaishnikKey
)

func Main() *models.InlineKeyboardMarkup {
	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "ГАИ", CallbackData: "main_" + strconv.Itoa(GaiKey)},
			}, {
				{Text: "Сотрудник ГИБДД", CallbackData: "main_" + strconv.Itoa(GaishnikKey)},
			},
		},
	}
	return kb
}
