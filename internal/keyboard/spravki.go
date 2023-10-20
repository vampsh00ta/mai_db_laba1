package keyboard

import "github.com/go-telegram/bot/models"

const (
	IsPersonOwnerCommand      = "Проверить,принадлежит авто человеку"
	GetPersonsVehiclesCommand = "Вывести автомобили, принадлежащие человеку"

	GetPersonInfoCommand          = "Вывести данные человека"
	GetOfficersInfoCommand        = "ВЫвести данные сотрудника ГИБДД"
	DtpActionsCommand             = "Дтп"
	GetDtpsInfoNearMetroCommand   = "Вывести ДТП, произошедшие у конкретного метро"
	GetDtpsInfoRadiusMetroCommand = "Вывести ДТП, произошедшие в n радиуса от метро"
)

func Spravki() *models.ReplyKeyboardMarkup {

	kb := &models.ReplyKeyboardMarkup{
		Keyboard: [][]models.KeyboardButton{
			{
				{Text: IsPersonOwnerCommand},
			}, {
				{Text: GetPersonsVehiclesCommand},
			},
			{
				{Text: GetPersonInfoCommand},
			},
			{
				{Text: GetOfficersInfoCommand},
			},
			{
				{Text: GetDtpsInfoNearMetroCommand},
			},
			{
				{Text: GetDtpsInfoRadiusMetroCommand},
			},
			{
				{Text: BackCommand},
			},
		},
	}
	return kb
}
