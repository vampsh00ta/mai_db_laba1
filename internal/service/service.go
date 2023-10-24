package service

import (
	rep "TgDbMai/internal/repository"
)

type Service interface {
	Dtp
	Spravki
	CrewI
}
type service struct {
	rep rep.Repository
}

func New(r rep.Repository) Service {
	return &service{
		rep: r,
	}
}
