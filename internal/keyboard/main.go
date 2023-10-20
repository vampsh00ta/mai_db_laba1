package keyboard

import (
	"github.com/go-telegram/bot/models"
)

const (
	GaiKey = iota
	GaishnikKey
)

const (
	GaiCommand      = "ГАИ"
	GaishnikCommand = "Сотрудник ГИБДД"
	MainСommand     = "Главное меню"
	Spravka         = "Справки"
)

func Main() *models.ReplyKeyboardMarkup {

	kb := &models.ReplyKeyboardMarkup{
		Keyboard: [][]models.KeyboardButton{
			{
				{Text: GaiCommand},
			}, {
				{Text: GaishnikCommand},
			},
		},
	}
	return kb
}
