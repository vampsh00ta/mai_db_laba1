package query_handlers

import (
	"TgDbMai/internal/handler"
	"TgDbMai/internal/keyboard"
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"strconv"
)

func Gai(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {

	b.AnswerCallbackQuery(ctx, &tgbotapi.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		ShowAlert:       false,
	})
	switch update.CallbackQuery.Data {
	case "gai_" + strconv.Itoa(keyboard.RegDtpKey):
		handler.Dtp(ctx, b, update)
	case "gai_" + strconv.Itoa(keyboard.RegVehicleKey):
		fmt.Printf(strconv.Itoa(keyboard.RegVehicleKey))
	case "gai_" + strconv.Itoa(keyboard.BackMainKey):

	}

}
