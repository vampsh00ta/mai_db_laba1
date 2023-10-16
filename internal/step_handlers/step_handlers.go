package step_handlers

import "TgDbMai/internal/service"

type StepHandler struct {
	s service.Service
}

func New(s service.Service) *StepHandler {
	return &StepHandler{
		s,
	}
}
