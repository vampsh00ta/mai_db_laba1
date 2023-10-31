package keyboard

import (
	rep "TgDbMai/internal/repository"
	"github.com/go-telegram/bot/models"
)

const (
// _ = int(iota)
// AddDtpParticipantKey
// CheckVehicleKey"Проверить автомобиль"
)
const (
	AddParticipantDtpCommand = "Добавить участника в текущее дтп"
	AddChangesToDtpCommand   = "Добавить комментарии по дтп"
	GetCurrentDtpCommand     = "Текущее дтп"
	CheckVehicleCommand      = "Проверить автомобиль"
	ExitCommand              = "Выйти"
	BackCommand              = "Назад"
	CurrentDtpCommand        = "Текущее дтп"
	IssueFineCommand         = "Выписать штраф"
	CheckFinesCommand        = "Проверить неоплаченные штрафы"
)

func Gaishnik() *models.ReplyKeyboardMarkup {
	kb := &models.ReplyKeyboardMarkup{
		Keyboard: [][]models.KeyboardButton{
			{
				{Text: AddParticipantDtpCommand},
			},
			{
				{Text: CurrentDtpCommand},
			},
			{
				{Text: AddChangesToDtpCommand},
			},
			{
				{Text: CheckVehicleCommand},
			},

			{
				{Text: IssueFineCommand},
			},
			{
				{Text: CheckFinesCommand},
			},
			{
				{Text: ExitCommand},
			},
			{
				{Text: BackCommand},
			},
		},
	}
	return kb
}

func DescriptionUpdate() *models.ReplyKeyboardMarkup {
	kb := &models.ReplyKeyboardMarkup{
		Keyboard: [][]models.KeyboardButton{
			{
				{Text: rep.DoneProtocol},
			},
			{
				{Text: rep.TestOfAlcDrugs},
			},
			{
				{Text: rep.HandOverToPolice},
			},
			{
				{Text: rep.ClosedDtp},
			},
		},
	}
	return kb
}
