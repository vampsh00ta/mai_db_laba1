package query_handlers

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// key: func
func NewQueryHandler(key_func ...any) tgbotapi.HandlerFunc {
	funcs := make(map[string]tgbotapi.HandlerFunc)
	fmt.Println(key_func)
	for i := 0; i < len(key_func)/2; i += 2 {
		key := key_func[i].(string)
		f := key_func[i+1].(func(ctx context.Context, bot *tgbotapi.Bot, update *models.Update))
		funcs[key] = f
	}
	return func(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
		b.AnswerCallbackQuery(ctx, &tgbotapi.AnswerCallbackQueryParams{
			CallbackQueryID: update.CallbackQuery.ID,
			ShowAlert:       false,
		})
		key := update.CallbackQuery.Data
		func_to_call := funcs[key]
		func_to_call(ctx, b, update)
	}
}
