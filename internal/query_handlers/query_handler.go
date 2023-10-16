package query_handlers

import (
	"TgDbMai/internal/step_handlers"
	tgbotapi "github.com/go-telegram/bot"
)

const (
	CheckVehiclePattern = "check_vehicle_"
	GaiPattern          = "gai_"
	GaishnikPattern     = "gaishnik_"
	MainPattern         = "main_"
	BackPattern         = "back_"
)

func New(step *step_handlers.StepHandler, bot *tgbotapi.Bot) *BotHandler {
	return &BotHandler{
		step,
		bot,
	}
}

func (bot BotHandler) Init() {
	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData, CheckVehiclePattern, tgbotapi.MatchTypePrefix, bot.CheckVehicle(CheckVehiclePattern))

	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData, GaiPattern, tgbotapi.MatchTypePrefix, bot.Gai(GaiPattern))
	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData, GaishnikPattern, tgbotapi.MatchTypePrefix, bot.Gaishnik(GaishnikPattern))
	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData, MainPattern, tgbotapi.MatchTypePrefix, bot.Main(MainPattern))

	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData, BackPattern, tgbotapi.MatchTypePrefix,
		bot.Back(BackPattern))
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText, "/start", tgbotapi.MatchTypeExact, bot.Start)

}
