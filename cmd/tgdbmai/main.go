package main

import (
	"TgDbMai/config"
	handlers "TgDbMai/internal/handlers"
	"TgDbMai/internal/psql"
	"TgDbMai/internal/service"
	"TgDbMai/internal/step_handlers"

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
	srvc := service.New(rep)
	stepH := step_handlers.New(srvc)
	opts := []tgbotapi.Option{
		tgbotapi.WithDefaultHandler(handler),
	}

	bot, err := tgbotapi.New(cfg.ApiToken, opts...)
	if err != nil {
		panic(err)
	}
	handlers.New(bot, stepH)
	bot.Start(ctx)

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
