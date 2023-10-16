package response

import (
	"TgDbMai/internal/psql"
	"github.com/go-telegram/bot/models"
	"strconv"
)

//	func (g gai) Markup() *models.InlineKeyboardMarkup {
//		kb := &models.InlineKeyboardMarkup{
//			InlineKeyboard: [][]models.InlineKeyboardButton{
func VehicleDpts(dtps []*psql.Dtp) *models.InlineKeyboardMarkup {

	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Id", CallbackData: "pass"},
				{Text: "Время", CallbackData: "pass"},
				{Text: "Category", CallbackData: "pass"},
			},
		},
	}
	for _, dtp := range dtps {
		res := []models.InlineKeyboardButton{
			{
				Text: strconv.Itoa(dtp.Id), CallbackData: "pass",
			},
			{
				Text: dtp.Date.String(), CallbackData: "pass",
			},
			{
				Text: dtp.Category, CallbackData: "pass",
			},
		}
		kb.InlineKeyboard = append(kb.InlineKeyboard, res)
	}
	return kb
}

func VehicleOwners(persons []*psql.Person) *models.InlineKeyboardMarkup {

	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Имя", CallbackData: "pass"},
				{Text: "Фамилия", CallbackData: "pass"},
				{Text: "Отчество", CallbackData: "pass"},
				{Text: "Паспорт", CallbackData: "pass"},
			},
		},
	}
	for _, person := range persons {
		res := []models.InlineKeyboardButton{

			{
				Text: person.Name, CallbackData: "pass",
			},
			{
				Text: person.Surname, CallbackData: "pass",
			},
			{
				Text: person.Patronymic, CallbackData: "pass",
			},
			{
				Text: strconv.Itoa(person.Passport), CallbackData: "pass",
			},
		}
		kb.InlineKeyboard = append(kb.InlineKeyboard, res)
	}
	return kb
}
