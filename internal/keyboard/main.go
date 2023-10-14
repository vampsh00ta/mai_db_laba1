package keyboard

import (
	"github.com/go-telegram/bot/models"
	"strconv"
	"sync"
)

const (
	GaiKey = iota
	GaishnikKey
)

const ()

type Main struct {
	Keyboard
}

var mainiknstance Main = Main{}
var oncemain sync.Once

func NewMain(pattern string) KeyboardI {
	oncemain.Do(func() {
		mainiknstance = Main{Keyboard{Pattern: pattern}}
	})

	return mainiknstance
}
func (m Main) CallbackData(key int) string {
	return m.Pattern + strconv.Itoa(key)

}
func (m Main) Markup() *models.InlineKeyboardMarkup {

	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "ГАИ", CallbackData: m.CallbackData(GaiKey)},
			}, {
				{Text: "Сотрудник ГИБДД", CallbackData: m.CallbackData(GaishnikKey)},
			},
		},
	}
	return kb
}
