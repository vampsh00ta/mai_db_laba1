package query_handlers

import (
	"TgDbMai/internal/keyboard"
	"context"
	tgbotapi "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type Gai struct {
	*BotHandler
}

func NewGai(bot *tgbotapi.Bot, handler *BotHandler) {
	gai := Gai{handler}
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		keyboard.RegVehicleCommand,
		tgbotapi.MatchTypeExact,
		gai.RegVehicle())
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		keyboard.DtpHappenCommand,
		tgbotapi.MatchTypeExact,
		gai.DtpHappen())
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		keyboard.IsPersonOwnerCommand,
		tgbotapi.MatchTypeExact,
		gai.IsPersonOwner())
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		keyboard.GetPersonsVehiclesCommand,
		tgbotapi.MatchTypeExact,
		gai.GetPersonsVehicles())
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		keyboard.GetPersonInfoCommand,
		tgbotapi.MatchTypeExact,
		gai.GetPersonInfo())
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		keyboard.GetOfficersInfoCommand,
		tgbotapi.MatchTypeExact,
		gai.GetOfficersInfoCommand())
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		keyboard.GetDtpsInfoNearAreaCommand,
		tgbotapi.MatchTypeExact,
		gai.GetDtpsInfoNearMetro())
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		keyboard.GetDtpsInfoRadiusMetroCommand,
		tgbotapi.MatchTypeExact,
		gai.GetDtpsInfoRadiusMetro())
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		keyboard.ByFIOCommand,
		tgbotapi.MatchTypeExact,
		gai.GetPersonInfoFIO())
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		keyboard.ByPassportCommand,
		tgbotapi.MatchTypeExact,
		gai.GetPersonInfoPassport())
}

func (g Gai) DtpHappen() tgbotapi.HandlerFunc {
	return func(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {
		g.step.Dtp(ctx, bot, update)
		//err := g.step.Producer.WriteMessages()
	}
}

func (g Gai) RegVehicle() tgbotapi.HandlerFunc {
	return func(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {
		//g.step.Veh
	}
}

func (s Gai) IsPersonOwner() tgbotapi.HandlerFunc {
	return func(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {
		s.step.IsPersonOwner(ctx, bot, update)

	}
}

// добабив вывод дтп, в которых был автомобиль
func (s Gai) GetPersonsVehicles() tgbotapi.HandlerFunc {
	return func(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {
		s.step.GetPersonsVehicles(ctx, bot, update)
	}
}
func (s Gai) GetPersonInfo() tgbotapi.HandlerFunc {
	return func(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {
		s.back.Set(update.Message.Chat.ID,
			&Back{
				name:     keyboard.GaiCommand,
				keyboard: keyboard.Gai(),
			},
		)
		bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID:      update.Message.Chat.ID,
			Text:        keyboard.GetPersonInfoCommand,
			ReplyMarkup: keyboard.GetPersonInfo(),
		})
	}
}
func (s Gai) GetPersonInfoFIO() tgbotapi.HandlerFunc {
	return func(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {
		s.step.GetPersonInfoFIO(ctx, bot, update)
	}
}
func (s Gai) GetPersonInfoPassport() tgbotapi.HandlerFunc {
	return func(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {
		s.step.GetPersonInfoPassport(ctx, bot, update)
	}
}

func (s Gai) GetOfficersInfoCommand() tgbotapi.HandlerFunc {
	return func(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {
		s.step.GetOfficersInfo(ctx, bot, update)

	}
}
func (s Gai) GetDtpsInfoNearMetro() tgbotapi.HandlerFunc {
	return func(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {
		s.step.GetDtpsInfoNearArea(ctx, bot, update)
	}
}
func (s Gai) GetDtpsInfoRadiusMetro() tgbotapi.HandlerFunc {
	return func(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {
		s.step.GetDtpsInfoRadius(ctx, bot, update)
	}
}
