package keyboard

import (
	"github.com/go-telegram/bot/models"
)

type Keyboard struct {
	Pattern string
}
type KeyboardI interface {
	CallbackData(key int) string
	Markup() *models.InlineKeyboardMarkup
}

//func (Keyboard *Keyboard) CallbackData(key int) string{
//	return Keyboard.Pattern + strconv.Itoa(key)
//}
