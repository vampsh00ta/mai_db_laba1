package keyboard

import (
	"github.com/go-telegram/bot/models"
)

const (
// _ = int(iota)
// AddDtpParticipantKey
// CheckVehicleKey"Проверить автомобиль"
)
const (
	AddParticipantDtpCommand = "Добавить участника дтп"
	CheckVehicleCommand      = "Проверить автомобиль"
	BackCommand              = "Назад"
)

func Gaishnik() *models.ReplyKeyboardMarkup {

	kb := &models.ReplyKeyboardMarkup{
		Keyboard: [][]models.KeyboardButton{
			{
				{Text: AddParticipantDtpCommand},
			},
			{
				{Text: CheckVehicleCommand},
			},
			//{
			//	{Text: "Выписать штраф", CallbackData: "gaishnik_"+string(CheckVehicle)},
			//},
			{
				{Text: BackCommand},
			},
		},
	}
	return kb
}
