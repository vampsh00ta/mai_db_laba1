package psql

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type VehicleRepository interface {
	GetVehicleByPts(tx *gorm.DB, pts string) (*Vehicle, error)
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
	err := tx.Model(&Person{}).
		Preload("Vehicles", db.Where("vehicle.pts = ?", pts)).
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
