package main

import (
	"TgDbMai/config"
	"TgDbMai/internal/keyboard"

	handlers "TgDbMai/internal/handler"
	"TgDbMai/internal/psql"

	tgbotapi "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"

	"context"
	"fmt"
	"golang.org/x/exp/slog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
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
	if err != nil {
		panic(err)
	}
	rep := psql.New(db, logger)
	rep = rep
	opts := []tgbotapi.Option{
		tgbotapi.WithDefaultHandler(handler),
	}

	bot, err := tgbotapi.New(cfg.ApiToken, opts...)
	if err != nil {
		panic(err)
	}
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText, "/start", tgbotapi.MatchTypeExact, helloHandler)
	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData, "main", tgbotapi.MatchTypePrefix, callbackHandler)

	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData, "back_main", tgbotapi.MatchTypeExact, backMain)

	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData, "gai", tgbotapi.MatchTypePrefix, callbackHandler)
	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData, "gagaishniki", tgbotapi.MatchTypePrefix, callbackHandler)

	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData, "test", tgbotapi.MatchTypePrefix, handlers.GaiDtpName)

	bot.Start(ctx)
}

func helloHandler(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
	b.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Выберите",
		ReplyMarkup: keyboard.Main(),
	})
}
func backMain(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
	b.AnswerCallbackQuery(ctx, &tgbotapi.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		ShowAlert:       false,
	})
	b.EditMessageReplyMarkup(ctx, &tgbotapi.EditMessageReplyMarkupParams{
		ChatID:      update.CallbackQuery.Message.Chat.ID,
		MessageID:   update.CallbackQuery.Message.ID,
		ReplyMarkup: keyboard.Main(),
	})
}
func test(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {

}
func callbackHandler(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {

	b.AnswerCallbackQuery(ctx, &tgbotapi.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		ShowAlert:       false,
	})
	if update.CallbackQuery.Data == "main_gai" {

		b.EditMessageReplyMarkup(ctx, &tgbotapi.EditMessageReplyMarkupParams{
			ChatID:      update.CallbackQuery.Message.Chat.ID,
			MessageID:   update.CallbackQuery.Message.ID,
			ReplyMarkup: keyboard.Gai(),
		})

	} else if update.CallbackQuery.Data == "main_gaishnik" {
		b.EditMessageReplyMarkup(ctx, &tgbotapi.EditMessageReplyMarkupParams{
			ChatID:      update.CallbackQuery.Message.Chat.ID,
			MessageID:   update.CallbackQuery.Message.ID,
			ReplyMarkup: keyboard.Gaishnik(),
		})
	}

}
func handler(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
	if update.InlineQuery == nil {
		return
	}

	results := []models.InlineQueryResult{
		&models.InlineQueryResultArticle{ID: "1", Title: "Foo 1", InputMessageContent: &models.InputTextMessageContent{MessageText: "foo 1"}},
		&models.InlineQueryResultArticle{ID: "2", Title: "Foo 2", InputMessageContent: &models.InputTextMessageContent{MessageText: "foo 2"}},
		&models.InlineQueryResultArticle{ID: "3", Title: "Foo 3", InputMessageContent: &models.InputTextMessageContent{MessageText: "foo 3"}},
	}

	b.AnswerInlineQuery(ctx, &tgbotapi.AnswerInlineQueryParams{
		InlineQueryID: update.InlineQuery.ID,
		Results:       results,
	})
}
func Gai_name(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
	//if v := ctx.Value("1"); v != nil {
	//	fmt.Println("found value:", v)
	//}
	//fmt.Println("key not found:", "1")
	var person psql.Person
	person.Name = update.Message.Text
	//user, _ := b.GetMe(ctx)
	//handlers["register_gai_person_"+strconv.Itoa(int(user.ID))] = person
	//idToUnregister := handlers["register_gai_person_id"].(string)
	//b.UnregisterHandler(idToUnregister)

	//id := b.RegisterHandler(tgbotapi.HandlerTypeMessageText, "", tgbotapi.MatchTypeExact, Gai_surname)
	//handlers["register_gai_person_id"] = id

	b.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "вв-те фамилию",
	})
}
func Gai_surname(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
	//if v := ctx.Value("1"); v != nil {
	//	fmt.Println("found value:", v)
	//}
	//fmt.Println("key not found:", "1")
	//user, _ := b.GetMe(ctx)

	//person := handlers["register_gai_person_"+strconv.Itoa(int(user.ID))].(psql.Person)
	//person.Surname = update.Message.Text
	//handlers["register_gai_person_"+strconv.Itoa(int(user.ID))] = person
	//idToUnregister := handlers["register_gai_person_id"].(string)
	//b.UnregisterHandler(idToUnregister)

	b.RegisterHandler(tgbotapi.HandlerTypeMessageText, "", tgbotapi.MatchTypeExact, helloHandler)

	b.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   update.Message.Text,
	})
}
