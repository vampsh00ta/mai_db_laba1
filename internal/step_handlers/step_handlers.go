package step_handlers

import (
	"TgDbMai/internal/service"
	"TgDbMai/internal/service/auth"
	log "TgDbMai/pkg/logger"
)

type StepHandler struct {
	s        service.Service
	log      *log.Logger
	Auth     auth.Auth
	Producer *service.Producer
}

func New(s service.Service, logger *log.Logger, auth auth.Auth, producer *service.Producer) *StepHandler {
	return &StepHandler{
		s,
		logger,
		auth,
		producer,
	}
}
