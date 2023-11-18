package main

import (
	"TgDbMai/config"
	handlers "TgDbMai/internal/handlers"
	repository "TgDbMai/internal/repository"
	"TgDbMai/internal/service"
	authentication "TgDbMai/internal/service/auth"
	log "TgDbMai/pkg/logger"
	"github.com/confluentinc/confluent-kafka-go/kafka"

	"TgDbMai/internal/step_handlers"
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram/bot"
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
	//a, err := rep.GetCurrentDtpByPersonId(tx, 5)
	//fmt.Println(a, err)
	auth := &authentication.AuthMap{DB: make(map[int64]*authentication.User)}
	auth.LogIn(564764193, 955, 2)
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"client.id":         "kafka",
		"acks":              "all"})
	if err != nil {
		fmt.Println(err)
	}

	srvc := service.New(rep)

	logger := log.New(cfg.Level)
	stepH := step_handlers.New(srvc, logger, auth, service.NewProducer(producer))
	opts := []tgbotapi.Option{
		//tgbotapi.WithMiddlewares(auth.AuthMiddleware()),
	}

	bot, err := tgbotapi.New(cfg.Apitoken, opts...)
	if err != nil {
		panic(err)
	}
	handlers.New(bot, stepH)
	bot.Start(ctx)

}
