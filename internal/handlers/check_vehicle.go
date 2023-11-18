package query_handlers

import (
	"TgDbMai/internal/keyboard"
	"context"
	tgbotapi "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type CheckVehicle struct {
	*BotHandler
}

func NewCheckVehicle(bot *tgbotapi.Bot, handler *BotHandler) {
	checkVehicle := CheckVehicle{handler}
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		keyboard.VehicleDtpsCommand,
		tgbotapi.MatchTypeExact,
		checkVehicle.VehicleDtps())
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		keyboard.VehicleOwnerCommand,
		tgbotapi.MatchTypeExact,
		checkVehicle.VehicleOwner())
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		keyboard.VehicleInfoByPtsCommand,
		tgbotapi.MatchTypeExact,
		checkVehicle.VehicleInfoByPts())
}

func (g CheckVehicle) VehicleDtps() tgbotapi.HandlerFunc {
	return func(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {
		g.step.VehicleDtps(ctx, bot, update)
	}
}
func (g CheckVehicle) VehicleInfoByPts() tgbotapi.HandlerFunc {
	return func(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {
		g.step.VehicleInfo(ctx, bot, update)

	}
}
func (g CheckVehicle) VehicleOwner() tgbotapi.HandlerFunc {
	return func(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {
		g.step.CheckVehicleOwners(ctx, bot, update)
	}
}
