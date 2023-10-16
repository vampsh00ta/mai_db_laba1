package keyboard

import (
	"github.com/go-telegram/bot/models"
	"strconv"
	"sync"
)

const (
	CheckVehicleDtpsKey = iota
	CheckVehicleOwnerKey
	CheckVehicleInfoByPts
)

var checkVehicleInstance = checkVehicle{}
var checkVehicleOnce sync.Once

type checkVehicle struct {
	Keyboard
}

func (c checkVehicle) CallbackData(key int) string {
	return c.Pattern + strconv.Itoa(key)
}
func (c checkVehicle) Markup() *models.InlineKeyboardMarkup {
	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Дтп ,в которых учавствовал автомобиль", CallbackData: c.CallbackData(CheckVehicleDtpsKey)},
			}, {
				{Text: "Владелец автомобиля", CallbackData: c.CallbackData(CheckVehicleOwnerKey)},
			},
			{
				{Text: "Информация автомобиля по его номеру", CallbackData: c.CallbackData(CheckVehicleInfoByPts)},
			},
			{
				{Text: "Назад", CallbackData: "back_" + strconv.Itoa(BackToGaishnikKey)},
			},
		},
	}
	return kb
}

func CheckVehicle(pattern string) KeyboardI {
	checkVehicleOnce.Do(func() {
		checkVehicleInstance = checkVehicle{Keyboard{Pattern: pattern}}
	})
	return checkVehicleInstance
}
