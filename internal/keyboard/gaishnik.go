package keyboard

import (
	"TgDbMai/internal/psql"
	"github.com/go-telegram/bot/models"
	"strconv"
	"sync"
)

const (
	_ = int(iota)
	AddDtpParticipantKey
	CheckVehicleKey
)

var gaishniknstance gaishnik = gaishnik{}
var oncegaishnik sync.Once

type gaishnik struct {
	Keyboard
}

func (g gaishnik) CallbackData(key int) string {
	return g.Pattern + strconv.Itoa(key)

}
func (g gaishnik) Markup() *models.InlineKeyboardMarkup {

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
func VehicleDpts(dtps []*psql.Dtp) *models.InlineKeyboardMarkup {

	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Id", CallbackData: ""},
				{Text: "Метро", CallbackData: ""},
				{Text: "Координаты", CallbackData: ""},
				{Text: "Улица", CallbackData: ""},
				{Text: "Время", CallbackData: ""},
			},
		},
	}
	for _, dtp := range dtps {
		res := []models.InlineKeyboardButton{
			{
				Text: string(dtp.Id),
			},
			{
				Text: dtp.Metro,
			},
			{
				Text: dtp.Coords,
			},
			{
				Text: dtp.Street,
			},
			{
				Text: dtp.Date.String(),
			},
		}
		kb.InlineKeyboard = append(kb.InlineKeyboard, res)
	}
	return kb
}

func Gaishnik(pattern string) KeyboardI {

	oncegaishnik.Do(func() {
		gaishniknstance = gaishnik{Keyboard{Pattern: pattern}}
	})

	return gaishniknstance
}
