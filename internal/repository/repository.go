package psql

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type Repository interface {
	DtpRepository
	PersonRepository
	DescriptionRepository
	CrewRepository
	VehicleRepository
	GetDb() *gorm.DB
}

type Db struct {
	*gorm.DB
}

func (db Db) GetDb() *gorm.DB {
	return db.DB
}

func New(db *gorm.DB) Repository {

	return &Db{
		DB: db}
}

func Error(funcName string, err error) error {
	description := fmt.Sprintf("repository:%s: %s", funcName, err.Error())
	return errors.New(description)
}
