package query_handlers

import (
	"TgDbMai/internal/keyboard"
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"strconv"
)

func Gaishnik(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
	b.AnswerCallbackQuery(ctx, &tgbotapi.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		ShowAlert:       false,
	})
	switch update.CallbackQuery.Data {
	case "gaishnik_" + strconv.Itoa(keyboard.CheckVehicleOwnerKey):
		fmt.Printf(strconv.Itoa(keyboard.CheckVehicleOwnerKey))

	case "gaishnik_" + strconv.Itoa(keyboard.AddDtpParticipantKey):
		fmt.Printf(strconv.Itoa(keyboard.AddDtpParticipantKey))

	case "gaishnik_" + strconv.Itoa(keyboard.CheckVehicleKey):
		fmt.Printf(strconv.Itoa(keyboard.CheckVehicleKey))

	}

}
