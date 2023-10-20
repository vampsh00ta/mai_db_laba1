package keyboard

import (
	"github.com/go-telegram/bot/models"
)

const (
	CheckVehicleDtpsKey = iota
	CheckVehicleOwnerKey
)
const (
	VehicleDtpsCommand      = "Дтп,в которых учавствовал автомобиль"
	VehicleOwnerCommand     = "Владелец автомобиля"
	VehicleInfoByPtsCommand = "Информация автомобиля по его номеру"
)

func CheckVehicle() *models.ReplyKeyboardMarkup {
	kb := &models.ReplyKeyboardMarkup{
		Keyboard: [][]models.KeyboardButton{
			{
				{Text: VehicleDtpsCommand},
			}, {
				{Text: VehicleOwnerCommand},
			},
			{
				{Text: VehicleInfoByPtsCommand},
			},
			{
				{Text: BackCommand},
			},
		},
	}
	return kb
}
