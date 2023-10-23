package step_handlers

import (
	"TgDbMai/internal/service"
	"TgDbMai/internal/service/auth"
	log "TgDbMai/pkg/logger"
)

type StepHandler struct {
	s    service.Service
	log  *log.Logger
	Auth *auth.Auth
}

func New(s service.Service, logger *log.Logger, auth *auth.Auth) *StepHandler {
	return &StepHandler{
		s,
		logger,
		auth,
	}
}
