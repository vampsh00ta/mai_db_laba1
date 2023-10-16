package psql

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type VehicleRepository interface {
	GetVehicleByPts(tx *gorm.DB, pts string) (*Vehicle, error)
	GetVehicleDptsByPts(tx *gorm.DB, pts string) ([]*Dtp, error)
	RegisterVehicle(tx *gorm.DB, passport int, vehicle *Vehicle) error
	GetVehicleOwners(tx *gorm.DB, pts string) ([]*Person, error)
	DeleteVechilesFromPerson(tx *gorm.DB, passport int, vehicle []*Vehicle) error
}

func (db Db) GetVehicleByPts(tx *gorm.DB, pts string) (*Vehicle, error) {

	var vehicle Vehicle
	err := tx.Model(&Vehicle{}).
		Where("pts = ?", pts).
		Find(&vehicle).Error
	if err != nil {
		description := fmt.Sprintf("repository:GetVehicleByPts: %s", err.Error())
		return nil, errors.New(description)
	}
	return &vehicle, nil
}
func (db Db) GetVehicleOwners(tx *gorm.DB, pts string) ([]*Person, error) {
	var persons []*Person
	//err := tx.Model(&Person{}).
	//
	//	Preload("Vehicles", db.Where("pts = ?", pts)).
	//	Find(&persons).Error
	err := tx.Raw(
		`select person.name,person.surname,person.patronymic,person.passport from person
			join person_vehicle on person_vehicle.person_id = person.id
			join (select * from vehicle where pts = ?) as vehicle on vehicle.id = person_vehicle.vehicle_id
		`, pts).
		Find(&persons).Error
	if err != nil {
		description := fmt.Sprintf("repository:GetVehicleOwners: %s", err.Error())
		return nil, errors.New(description)
	}
	return persons, nil
}
func (db Db) RegisterVehicle(tx *gorm.DB, passport int, vehicle *Vehicle) error {
	var person *Person
	err := tx.First(&person, "passport = ?", passport).Error
	if person == nil {
		description := fmt.Sprintf("repository:RegisterVehicle: %s", "no  such person")
		return errors.New(description)
	}
	if err != nil {
		description := fmt.Sprintf("repository:RegisterVehicle: %s", err.Error())
		return errors.New(description)
	}

	if err := tx.Model(&person).Association("Vehicles").Append([]*Vehicle{vehicle}); err != nil {
		description := fmt.Sprintf("repository:RegisterVehicle: %s", err.Error())
		return errors.New(description)
	}

	return nil
}

func (db Db) DeleteVechilesFromPerson(tx *gorm.DB, passport int, vehicle []*Vehicle) error {
	var person *Person
	err := tx.First(&person, "passport = ?", passport).Error
	if person == nil {
		description := fmt.Sprintf("repository:DeleteVechilesFromPerson: %s", "no  such person")
		return errors.New(description)
	}
	if err != nil {
		description := fmt.Sprintf("repository:DeleteVechilesFromPerson: %s", err.Error())
		return errors.New(description)
	}

	if err := tx.Model(&person).Association("Vehicles").Delete(vehicle); err != nil {
		description := fmt.Sprintf("repository:DeleteVechilesFromPerson: %s", err.Error())
		return errors.New(description)
	}

	return nil
}

func (db Db) GetVehicleDptsByPts(tx *gorm.DB, pts string) ([]*Dtp, error) {
	var dtps []*Dtp
	err := tx.Raw(`select dtp.id,dtp.date,dtp.category
		from (select * from vehicle where pts = ?) as vehicle
		join person_vehicle on vehicle.id = person_vehicle.vehicle_id
		join person on person.id = person_vehicle.person_id 
		join participant_of_dtp on participant_of_dtp.person_id = person.id
		join dtp on dtp.id = participant_of_dtp.dtp_id`, pts).
		Find(&dtps).Error
	if err != nil {
		description := fmt.Sprintf("repository:GetVehicleByPts: %s", err.Error())
		return nil, errors.New(description)
	}
	return dtps, nil
}
