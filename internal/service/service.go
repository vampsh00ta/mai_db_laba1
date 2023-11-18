package service

import (
	rep "TgDbMai/internal/repository"
)

type Service interface {
	Dtp
	Crew
	Person
	Fine
	Vehicle
}
type service struct {
	rep rep.Repository
}

func New(r rep.Repository) Service {
	return &service{
		rep: r,
	}
}
