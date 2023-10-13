package main

import (
	"TgDbMai/config"
	"TgDbMai/internal/keyboard"
	"TgDbMai/internal/query_handlers"

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
	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData, "main_", tgbotapi.MatchTypePrefix, query_handlers.Main)
	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData, "gai_", tgbotapi.MatchTypePrefix, query_handlers.Gai)
	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData, "gaishnik_", tgbotapi.MatchTypePrefix, query_handlers.Gaishnik)

	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData, "back_", tgbotapi.MatchTypePrefix, query_handlers.Back)
	//
	//bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData, "gai", tgbotapi.MatchTypePrefix, callbackHandler)
	//bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData, "gagaishniki", tgbotapi.MatchTypePrefix, callbackHandler)

	//bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData, "dtp", tgbotapi.MatchTypePrefix, handlers.Dtp)

	bot.Start(ctx)
}

func helloHandler(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
	b.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Выберите",
		ReplyMarkup: keyboard.Main(),
	})
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
