package keyboard

import "github.com/go-telegram/bot/models"

const (
	IsPersonOwnerCommand      = "Проверить,принадлежит авто человеку"
	GetPersonsVehiclesCommand = "Вывести автомобили, принадлежащие человеку"

	GetPersonInfoCommand          = "Вывести данные человека"
	GetOfficersInfoCommand        = "ВЫвести данные сотрудника ГИБДД"
	DtpActionsCommand             = "Дтп"
	GetDtpsInfoNearAreaCommand    = "Вывести ДТП, произошедшие в конкретном районе"
	GetDtpsInfoRadiusMetroCommand = "Вывести ДТП, произошедшие в n радиуса от метро"
	ByPassportCommand             = "По паспорту"
	ByFIOCommand                  = "По ФИО"
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
				{Text: GetDtpsInfoNearAreaCommand},
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

func GetPersonInfo() *models.ReplyKeyboardMarkup {

	kb := &models.ReplyKeyboardMarkup{
		Keyboard: [][]models.KeyboardButton{
			{
				{Text: ByPassportCommand},
			}, {
				{Text: ByFIOCommand},
			},
			{
				{Text: BackCommand},
			},
		},
	}
	return kb
}
