package psql

import (
	"golang.org/x/exp/slog"
	"gorm.io/gorm"
)

type Repository interface {
	DtpRepository
	PersonI
	VehicleRepository
	GetDb() *gorm.DB
}

type Db struct {
	*gorm.DB
	logger *slog.Logger
}

func (db Db) GetDb() *gorm.DB {
	return db.DB
}

func New(db *gorm.DB, logger *slog.Logger) Repository {

	return &Db{
		DB:     db,
		logger: logger}
}
