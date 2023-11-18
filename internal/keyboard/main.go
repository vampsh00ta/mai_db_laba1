package keyboard

import (
	"github.com/go-telegram/bot/models"
)

const (
	GaiCommand        = "ГАИ"
	GaishnikCommand   = "Сотрудник ГИБДД"
	MainСommand       = "Главное меню"
	SpravkiCommand    = "Справки"
	MasterDataCommand = "Master Data"
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
