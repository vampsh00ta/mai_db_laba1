package main

import (
	"TgDbMai/config"
	handlers "TgDbMai/internal/handlers"
	repository "TgDbMai/internal/repository"
	"TgDbMai/internal/service"
	authentication "TgDbMai/internal/service/auth"
	"TgDbMai/internal/step_handlers"
	log "TgDbMai/pkg/logger"
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlog "gorm.io/gorm/logger"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow",
		cfg.PG.Host,
		cfg.PG.Username,
		cfg.PG.Password,
		cfg.PG.Name,
		cfg.PG.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormlog.Default.LogMode(gormlog.Info),
	})
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	rep := repository.New(db)
	//tx := rep.GetDb()
	//res, err := rep.GetOfficersCrewByOfficerId(tx, 1)
	//fmt.Println(res[0].PoliceOfficer1, err)
	auth := &authentication.Auth{DB: make(map[int64]*authentication.User)}
	srvc := service.New(rep)
	logger := log.New(cfg.Level)
	stepH := step_handlers.New(srvc, logger, auth)
	opts := []tgbotapi.Option{
		tgbotapi.WithDefaultHandler(handler),
		tgbotapi.WithMiddlewares(auth.AuthMiddleware()),
	}

	bot, err := tgbotapi.New(cfg.Apitoken, opts...)
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
