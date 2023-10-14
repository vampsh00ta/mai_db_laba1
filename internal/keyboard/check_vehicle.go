package keyboard

import (
	"github.com/go-telegram/bot/models"
	"strconv"
	"sync"
)

const (
	CheckVehicleDtpsKey = iota
	CheckVehicleOwnerKey
)

var checkVehicleInstance = CheckVehicle{}
var checkVehicleOnce sync.Once

func NewCheckVehicle(pattern string) KeyboardI {
	checkVehicleOnce.Do(func() {
		checkVehicleInstance = CheckVehicle{Keyboard{Pattern: pattern}}
	})
	return checkVehicleInstance
}

type CheckVehicle struct {
	Keyboard
}

func (c CheckVehicle) CallbackData(key int) string {
	return c.Pattern + strconv.Itoa(key)
}
func (c CheckVehicle) Markup() *models.InlineKeyboardMarkup {
	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Дтп ,в которых учавствовал автомобиль", CallbackData: c.CallbackData(CheckVehicleDtpsKey)},
			}, {
				{Text: "Владелец автомобиля", CallbackData: c.CallbackData(CheckVehicleOwnerKey)},
			},
			//{
			//	{Text: "Проверить автомобиль", CallbackData: "gai_" + strconv.Itoa(CheckVehicleKey)},
			//},
			{
				{Text: "Назад", CallbackData: "back_" + strconv.Itoa(BackToGaishnikKey)},
			},
		},
	}
	return kb
}
