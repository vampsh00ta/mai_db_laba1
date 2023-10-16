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

type main struct {
	Keyboard
}

var mainiknstance main = main{}
var oncemain sync.Once

func (m main) CallbackData(key int) string {
	return m.Pattern + strconv.Itoa(key)

}
func (m main) Markup() *models.InlineKeyboardMarkup {

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

func Main(pattern string) KeyboardI {
	oncemain.Do(func() {
		mainiknstance = main{Keyboard{Pattern: pattern}}
	})

	return mainiknstance
}
