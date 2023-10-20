package psql

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type PersonI interface {
	GetPersonByPts(tx *gorm.DB, name, surname, patronymic, pts string) (*Vehicle, error)
	GetPersonsVehicles(tx *gorm.DB, name, surname, patronymic string) ([]*Vehicle, error)
	GetPersonInfo(tx *gorm.DB, name, surname, patronymic string) (*Person, error)
}

func (db Db) GetPersonByPts(tx *gorm.DB, name, surname, patronymic, pts string) (*Vehicle, error) {
	var vehicle Vehicle
	err := tx.Raw(`
		select pts, model, category from (select * from vehicle where pts = ?) as vehicle
		join person_vehicle on person_vehicle.vehicle_id = vehicle.id
		join (select * from person  where name = ?and surname = ? and patronymic = ?) as person
		                            on person.id = person_vehicle.person_id
		`, pts, name, surname, patronymic).
		Find(&vehicle).Error
	if err != nil {
		description := fmt.Sprintf("repository:GetVehicleByPts: %s", err.Error())
		return nil, errors.New(description)
	}
	return &vehicle, nil
}
func (db Db) GetPersonsVehicles(tx *gorm.DB, name, surname, patronymic string) ([]*Vehicle, error) {
	var vehicles []*Vehicle
	err := tx.Raw(`
		select pts, model, category from  vehicle
		join person_vehicle on person_vehicle.vehicle_id = vehicle.id
		join (select * from person  where name = ?and surname = ? and patronymic = ?) as person
		                            on person.id = person_vehicle.person_id
		`, name, surname, patronymic).
		Find(&vehicles).Error
	if err != nil {
		description := fmt.Sprintf("repository:GetVehicleByPts: %s", err.Error())
		return nil, errors.New(description)
	}
	return vehicles, nil

}
func (db Db) GetPersonInfo(tx *gorm.DB, name, surname, patronymic string) (*Person, error) {
	return nil, nil
}
