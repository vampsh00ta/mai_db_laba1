package keyboard

import (
	"github.com/go-telegram/bot/models"
)

const (
	RegDtpKey = iota
	RegVehicleKey
	_
)
const (
	DtpHappenCommand = "Случилось дтп"
)

func Gai() *models.ReplyKeyboardMarkup {
	kb := &models.ReplyKeyboardMarkup{
		Keyboard: [][]models.KeyboardButton{
			{
				{Text: DtpHappenCommand},
			}, {
				{Text: RegVehicleCommand},
			},
			{
				{Text: CheckVehicleCommand},
			},
			{
				{Text: BackCommand},
			},
		},
	}
	return kb
}
