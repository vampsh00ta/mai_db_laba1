package keyboard

import (
	"github.com/go-telegram/bot/models"
	"strconv"
	"sync"
)

const (
	_ = int(iota)
	AddDtpParticipantKey
	CheckVehicleKey
)

var gaishniknstance Gaishnik = Gaishnik{}
var oncegaishnik sync.Once

type Gaishnik struct {
	Keyboard
}

func NewGaishnik(pattern string) KeyboardI {

	oncegaishnik.Do(func() {
		gaishniknstance = Gaishnik{Keyboard{Pattern: pattern}}
	})

	return gaishniknstance
}
func (g Gaishnik) CallbackData(key int) string {
	return g.Pattern + strconv.Itoa(key)

}
func (g Gaishnik) Markup() *models.InlineKeyboardMarkup {

	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Добавить участника дтп", CallbackData: g.CallbackData(AddDtpParticipantKey)},
			},
			{
				{Text: "Проверить автомобиль", CallbackData: g.CallbackData(CheckVehicleKey)},
			},
			//{
			//	{Text: "Выписать штраф", CallbackData: "gaishnik_"+string(CheckVehicle)},
			//},
			{
				{Text: "Назад", CallbackData: BackCallbackData(BackToMainKey)},
			},
		},
	}
	return kb
}
