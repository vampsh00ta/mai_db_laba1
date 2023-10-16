package service

import (
	"TgDbMai/internal/psql"
)

type service struct {
	rep psql.Repository
}

func New(r psql.Repository) Service {
	return &service{r}
}
