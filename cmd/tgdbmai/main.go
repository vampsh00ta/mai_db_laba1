package main

import (
	"TgDbMai/config"
	"TgDbMai/internal/psql"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"golang.org/x/exp/slog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// callback
var callback = tgbotapi.NewInlineKeyboardMarkup(
	//добавляет строку
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("1.com", "http://1.com"),
		tgbotapi.NewInlineKeyboardButtonData("2", "xyu"),
		tgbotapi.NewInlineKeyboardButtonData("3", "3"),
	),
	//добавляет строку

	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("4", "4"),
		tgbotapi.NewInlineKeyboardButtonData("5", "5"),
		tgbotapi.NewInlineKeyboardButtonData("6", "6"),
	),
)

// main menu
var menu = tgbotapi.NewReplyKeyboard(
	//добавляет строку
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("1"),
		tgbotapi.NewKeyboardButton("2"),
		tgbotapi.NewKeyboardButton("vv3"),
	),
	//добавляет строку

	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("4"),
		tgbotapi.NewKeyboardButton("5"),
		tgbotapi.NewKeyboardButton("test"),
	),
)
var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Диспетчер", "dispatcher"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Сотрудник дпс", "dps_employee"),
	),
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow",
		cfg.Db.Host,
		cfg.Db.Username,
		cfg.Db.Password,
		cfg.Db.Name,
		cfg.Db.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	var logger *slog.Logger
	bot, err := tgbotapi.NewBotAPI(cfg.ApiToken)
	if err != nil {
		panic(err)
	}
	bot = bot
	rep := psql.New(db, logger)
	tx := db.Begin()
	person := psql.Person{Name: "ya", Surname: "ya", Patronymic: "ya", Passport: 11111}
	res, err := rep.AddParticipant(tx, "228", 1, "1488", "xyesos", person)
	//if err != nil {
	//	fmt.Println(err)
	//}
	if err := tx.Commit().Error; err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Id)
	//for _, car := range res {
	//	fmt.Println(car.Model)
	//}
	//bot.Debug = true
	//fmt.Sprintf(bot.Token)
	//log.Printf("Authorized on account %s", bot.Self.UserName)
	//
	//controllers.Controller(bot, rep)

}
