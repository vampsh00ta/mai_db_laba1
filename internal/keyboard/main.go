package keyboard

import (
	"github.com/go-telegram/bot/models"
)

const (
	GaiKey = iota
	GaishnikKey
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
			{
				{Text: SpravkiCommand},
			},
			{
				{Text: MasterDataCommand},
			},
		},
	}
	return kb
}
