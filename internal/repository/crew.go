package psql

import (
	"gorm.io/gorm"
)

type CrewRepository interface {
	GetAllCrews(tx *gorm.DB) ([]*Crew, error)
}

func (db Db) GetAllCrews(tx *gorm.DB) ([]*Crew, error) {
	var crews []*Crew
	//var a []*PoliceOfficer

	err := tx.Model(&Crew{}).
		Preload("PoliceOfficers.Person").
		Preload("Gai").
		Find(&crews).Error
	//err = tx.Model(&PoliceOfficer{}).
	//	Preload("Crews").
	//	Find(&a).Error
	if err != nil {
		return nil, Error("GetPersonInfoByPassport", err)

	}
	return crews, nil
}
