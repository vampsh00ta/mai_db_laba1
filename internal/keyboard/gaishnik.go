package keyboard

import (
	"fmt"
	"github.com/go-telegram/bot/models"
	"strconv"
)

const (
	CheckVehicleOwnerKey = int(iota)
	AddDtpParticipantKey
	CheckVehicleKey
)

func Gaishnik() *models.InlineKeyboardMarkup {
	fmt.Println("gaishnik_"+strconv.Itoa(CheckVehicleOwnerKey),
		"gaishnik_"+strconv.Itoa(AddDtpParticipantKey),
		"gaishnik_"+strconv.Itoa(CheckVehicleKey))
	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Проверить владельца автомобиля", CallbackData: "gaishnik_" + strconv.Itoa(CheckVehicleOwnerKey)},
			}, {
				{Text: "Добавить участника дтп", CallbackData: "gaishnik_" + strconv.Itoa(AddDtpParticipantKey)},
			},
			{
				{Text: "Проверить автомобиль", CallbackData: "gaishnik_" + strconv.Itoa(CheckVehicleKey)},
			},
			//{
			//	{Text: "Выписать штраф", CallbackData: "gaishnik_"+string(CheckVehicle)},
			//},
			{
				{Text: "Назад", CallbackData: "back_" + strconv.Itoa(BackToMainKey)},
			},
		},
	}
	return kb
}
