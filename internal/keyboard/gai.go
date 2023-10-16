package keyboard

import (
	"github.com/go-telegram/bot/models"
	"strconv"
	"sync"
)

const (
	RegDtpKey = iota
	RegVehicleKey
	_
)

var gainstance gai = gai{}
var oncegai sync.Once

type gai struct {
	Keyboard
}

func (g gai) CallbackData(key int) string {
	return g.Pattern + strconv.Itoa(key)

}

func (g gai) Markup() *models.InlineKeyboardMarkup {
	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Случилось дтп", CallbackData: g.CallbackData(RegDtpKey)},
			}, {
				{Text: "Зарегистрировать автомобиль", CallbackData: g.CallbackData(RegVehicleKey)},
			},
			{
				{Text: "Проверить автомобиль", CallbackData: g.CallbackData(CheckVehicleKey)},
			},
			{
				{Text: "Назад", CallbackData: BackCallbackData(BackToMainKey)},
			},
		},
	}
	return kb
}

func Gai(pattern string) KeyboardI {

	oncegai.Do(func() {
		gainstance = gai{Keyboard{Pattern: pattern}}
	})
	return gainstance
}
