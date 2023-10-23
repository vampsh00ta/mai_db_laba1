package psql

import (
	"errors"
	"fmt"
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
}

func (db Db) GetOfficersCrewByOfficerFIO(tx *gorm.DB, name, surname, patronymic string) ([]*Crew, error) {
	//TODO implement me
	panic("implement me")
}

func (db Db) GetOfficerByFIO(tx *gorm.DB, name, surname, patronymic string) (*PoliceOfficer, error) {
	//TODO implement me
	panic("implement me")
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
