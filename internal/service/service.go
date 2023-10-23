package service

import (
	rep "TgDbMai/internal/repository"
)

type service struct {
	rep rep.Repository
}

func New(r rep.Repository) Service {
	return &service{
		rep: r,
	}
}
