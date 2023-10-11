package psql

import (
	"golang.org/x/exp/slog"
	"gorm.io/gorm"
)

type Repository interface {
	DtpRepository
	//PersonI
	VehicleRepository
}

type Db struct {
	gorm.DB
	logger *slog.Logger
}

func New(db *gorm.DB, logger *slog.Logger) Repository {
	return &Db{
		DB:     *db,
		logger: logger}
}
