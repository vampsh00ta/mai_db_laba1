package query_handlers

import (
	tgbotapi "github.com/go-telegram/bot"
)

const (
	CheckVehiclePattern = "check_vehicle_"
	GaiPattern          = "gai_"
	GaishnikPattern     = "gaishnik_"
	MainPattern         = "main_"
	BackPattern         = "back_"
)

func New(bot *tgbotapi.Bot) {
	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData, CheckVehiclePattern, tgbotapi.MatchTypePrefix, NewVehicle(CheckVehiclePattern))

	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData, GaiPattern, tgbotapi.MatchTypePrefix, NewGai(GaiPattern))
	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData, GaishnikPattern, tgbotapi.MatchTypePrefix, NewGaishnik(GaishnikPattern))
	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData, MainPattern, tgbotapi.MatchTypePrefix, NewMain(MainPattern))

	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData, BackPattern, tgbotapi.MatchTypePrefix, NewBack(BackPattern))
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText, "/start", tgbotapi.MatchTypeExact, NewStart)

}
