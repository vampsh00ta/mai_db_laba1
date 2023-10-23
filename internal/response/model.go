package response

import (
	rep "TgDbMai/internal/repository"
	"fmt"
	"github.com/go-telegram/bot/models"
	"strconv"
)

//	func (g gai) Markup() *models.InlineKeyboardMarkup {
//		kb := &models.InlineKeyboardMarkup{
//			InlineKeyboard: [][]models.InlineKeyboardButton{
func VehicleDpts(dtps []*rep.Dtp) *models.InlineKeyboardMarkup {

	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Id", CallbackData: "pass"},
				{Text: "Время", CallbackData: "pass"},
				{Text: "Category", CallbackData: "pass"},
			},
		},
	}
	parseCategory := func(str string) string {
		if str == "" {
			return "Пусто"
		}
		return str
	}
	for _, dtp := range dtps {
		res := []models.InlineKeyboardButton{
			{
				Text: strconv.Itoa(dtp.Id), CallbackData: "pass",
			},
			{
				Text: dtp.Date.String()[0:10], CallbackData: "pass",
			},
			{
				Text: parseCategory(dtp.Category), CallbackData: "pass",
			},
		}
		kb.InlineKeyboard = append(kb.InlineKeyboard, res)
	}
	return kb
}

func VehicleOwners(persons []*rep.Person) *models.InlineKeyboardMarkup {

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
func GetPersonsVehicles(vehicles []*rep.Vehicle) *models.InlineKeyboardMarkup {

	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Модель", CallbackData: "pass"},
				{Text: "Номер", CallbackData: "pass"},
				{Text: "Категория", CallbackData: "pass"},
			},
		},
	}
	for _, vehicle := range vehicles {
		res := []models.InlineKeyboardButton{

			{
				Text: vehicle.Model, CallbackData: "pass",
			},
			{
				Text: vehicle.Pts, CallbackData: "pass",
			},
			{
				Text: vehicle.Category, CallbackData: "pass",
			},
		}
		kb.InlineKeyboard = append(kb.InlineKeyboard, res)
	}
	return kb
}
func GetPersonInfo(persons []*rep.Person) *models.InlineKeyboardMarkup {

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

func GetOfficerInfo(persons []*rep.PoliceOfficer) *models.InlineKeyboardMarkup {
	fmt.Println(persons[0].Person)
	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Оффицер", CallbackData: "pass"},
			},
			{
				{Text: "Имя", CallbackData: "pass"},
				{Text: "Фамилия", CallbackData: "pass"},
				{Text: "Отчество", CallbackData: "pass"},
				{Text: "Ранг", CallbackData: "pass"},

				{Text: "Паспорт", CallbackData: "pass"},
			},
		},
	}
	for _, person := range persons {
		res := []models.InlineKeyboardButton{

			{
				Text: person.Person.Name, CallbackData: "pass",
			},
			{
				Text: person.Person.Surname, CallbackData: "pass",
			},
			{
				Text: person.Person.Patronymic, CallbackData: "pass",
			},
			{
				Text: person.Rank, CallbackData: "pass",
			},
			{
				Text: strconv.Itoa(person.Person.Passport), CallbackData: "pass",
			},
		}
		kb.InlineKeyboard = append(kb.InlineKeyboard, res)
	}
	return kb
}
func Сrew(crews []*rep.Crew) *models.InlineKeyboardMarkup {

	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Взвод", CallbackData: "pass"},
			},
			{
				{Text: "Время", CallbackData: "pass"},
				{Text: "Работа", CallbackData: "pass"},
			},
		},
	}
	toStringBool := func(b bool) string {
		if b {
			return "Работают"
		}
		return "Не работают"
	}
	for _, crew := range crews {
		res := []models.InlineKeyboardButton{

			{
				Text: crew.Time.String()[0:10], CallbackData: "pass",
			},
			{
				Text: toStringBool(crew.Duty), CallbackData: "pass",
			},
		}
		kb.InlineKeyboard = append(kb.InlineKeyboard, res)
	}
	return kb
}
