package keyboard

import "github.com/go-telegram/bot/models"

const (
	RegVehicleCommand  = "автомобиль "
	RegPersonCommand   = "человека"
	RegGaishnikCommand = "сотрудника ГИБДД"
	AddCommand         = "Добавить"
	DeleteCommand      = "Удалить"
	UpdateCommand      = "Изменить"
	GetCommand         = "Получить"
)

func MasterData() *models.ReplyKeyboardMarkup {

	kb := &models.ReplyKeyboardMarkup{
		Keyboard: [][]models.KeyboardButton{
			{
				{Text: AddCommand},
			}, {
				{Text: DeleteCommand},
			},
			{
				{Text: UpdateCommand},
			},
			{
				{Text: GetCommand},
			},
			{
				{Text: BackCommand},
			},
		},
	}
	return kb
}
func Add() *models.ReplyKeyboardMarkup {

	kb := &models.ReplyKeyboardMarkup{
		Keyboard: [][]models.KeyboardButton{
			{
				{Text: RegVehicleCommand},
			}, {
				{Text: RegPersonCommand},
			},
			{
				{Text: RegGaishnikCommand},
			},

			{
				{Text: BackCommand},
			},
		},
	}
	return kb
}
