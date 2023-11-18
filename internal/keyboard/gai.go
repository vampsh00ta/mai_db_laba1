package keyboard

import (
	"github.com/go-telegram/bot/models"
)

const (
	RegDtpKey = iota
	RegVehicleKey
	_
)
const (
	DtpHappenCommand              = "Случилось дтп"
	IsPersonOwnerCommand          = "Проверить,принадлежит авто человеку"
	GetPersonsVehiclesCommand     = "Вывести автомобили, принадлежащие человеку"
	RegVehicleCommand             = "Зарегистировать автомобиль"
	GetPersonInfoCommand          = "Вывести данные человека"
	GetOfficersInfoCommand        = "ВЫвести данные сотрудника ГИБДД"
	DtpActionsCommand             = "Дтп"
	GetDtpsInfoNearAreaCommand    = "Вывести ДТП, произошедшие в конкретном районе"
	GetDtpsInfoRadiusMetroCommand = "Вывести ДТП, произошедшие в n радиуса от метро"
	ByPassportCommand             = "По паспорту"
	ByFIOCommand                  = "По ФИО"
)

func Gai() *models.ReplyKeyboardMarkup {
	kb := &models.ReplyKeyboardMarkup{
		Keyboard: [][]models.KeyboardButton{
			{
				{Text: DtpHappenCommand},
			}, {
				{Text: RegVehicleCommand},
			},
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
