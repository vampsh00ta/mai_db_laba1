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

var gainstance Gai = Gai{}
var oncegai sync.Once

type Gai struct {
	Keyboard
}

func NewGai(pattern string) KeyboardI {

	oncegai.Do(func() {
		gainstance = Gai{Keyboard{Pattern: pattern}}
	})
	return gainstance
}
func (g Gai) CallbackData(key int) string {
	return g.Pattern + strconv.Itoa(key)

}

func (g Gai) Markup() *models.InlineKeyboardMarkup {
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
