package psql

import (
	"gorm.io/gorm"
)

type PersonRepository interface {
	GetPersonByPts(tx *gorm.DB, name, surname, patronymic, pts string) (*Vehicle, error)
	GetPersonsVehicles(tx *gorm.DB, name, surname, patronymic string) ([]*Vehicle, error)
	GetPersonInfoByFIO(tx *gorm.DB, name, surname, patronymic string) ([]*Person, error)
	GetPersonInfoByPassport(tx *gorm.DB, passport int) (*Person, error)
	GetOfficersInfoByFIO(tx *gorm.DB, name, surname, patronymic string) (*PoliceOfficer, error)
	GetOfficersCrewByOfficerId(tx *gorm.DB, id int) ([]*Crew, error)
	GetOfficersCrewByOfficerFIO(tx *gorm.DB, name, surname, patronymic string) ([]*Crew, error)
	GetOfficers(tx *gorm.DB, name, surname, patronymic string) ([]*Crew, error)

	GetOfficerByFIO(tx *gorm.DB, name, surname, patronymic string) (*PoliceOfficer, error)
}

func (db Db) GetPersonByPts(tx *gorm.DB, name, surname, patronymic, pts string) (*Vehicle, error) {
	var vehicle Vehicle
	err := tx.Raw(`
		select vehicle.id,pts, model, category from (select * from vehicle where pts = ?) as vehicle
		join person_vehicle on person_vehicle.vehicle_id = vehicle.id
		join (select * from person  where name = ?and surname = ? and patronymic = ?) as person
		                            on person.id = person_vehicle.person_id
		`, pts, name, surname, patronymic).
		First(&vehicle).Error
	if err != nil {
		return nil, Error("GetPersonByPts", err)

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
		return nil, Error("GetPersonsVehicles", err)

	}
	return vehicles, nil

}
func (db Db) GetPersonInfoByFIO(tx *gorm.DB, name, surname, patronymic string) ([]*Person, error) {
	var persons []*Person
	err := tx.Raw(`
		select * from  person where name = ? and surname = ? and patronymic = ?
		
		`, name, surname, patronymic).
		Find(&persons).Error
	if err != nil {
		return nil, Error("GetPersonInfoByFIO", err)

	}
	return persons, nil
}
func (db Db) GetPersonInfoByPassport(tx *gorm.DB, passport int) (*Person, error) {
	var person *Person
	err := tx.Raw(`
		select * from  person where passport = ?
		
		`, passport).
		Find(&person).Error
	if err != nil {
		return nil, Error("GetPersonInfoByPassport", err)

	}
	return person, nil
}
func (db Db) GetOfficersInfoByFIO(tx *gorm.DB, name, surname, patronymic string) (*PoliceOfficer, error) {
	var person *Person
	err := tx.Model(&Person{}).Preload("PoliceOfficer").
		Where("name = ? and surname = ? and patronymic = ?", name, surname, patronymic).
		First(&person).Error
	if err != nil {
		return nil, Error("GetPersonInfoByFIO", err)

	}
	return person.PoliceOfficer, nil
}

func (db Db) GetOfficersCrewByOfficerId(tx *gorm.DB, id int) ([]*Crew, error) {
	var crews []*Crew
	err := tx.
		Model(&crews).
		Joins("join police_officer on police_officer.id = crew.p_officer_id_2 or police_officer.id = crew.p_officer_id_1").
		Joins("join person on person.id = police_officer.person_id where person.id = ?", id).
		Find(&crews).Error
	if err != nil {
		return nil, Error("GetOfficersCrewByOfficerId", err)

	}
	return crews, nil
}

//	func (db Db) GetOfficerByFIO(tx *gorm.DB, name, surname, patronymic string) (*PoliceOfficer, error) {
//		var person Person
//		err := tx.Model(&person).
//			Preload("PoliceOfficer.Crews1.Dtps").
//			Preload("PoliceOfficer.Crews2.Dtps").
//			Where(" name = ? and surname = ? and patronymic = ?", name, surname, patronymic).
//			First(&person).Error
//		if err != nil {
//			return nil, Error("GetPersonInfoByFIO", err)
//
//		}
//		return person.PoliceOfficer, nil
//	}
//func (db Db) GetOfficersCrewByOfficerFIO(tx *gorm.DB, name, surname, patronymic string) ([]*Crew, error) {
//	var crews []*Crew
//	err := tx.Model(&Crew{}).
//		Preload("PoliceOfficer1.Person",
//			db.Where("person.name = ? and  surname = ? and  patronymic = ?", name, surname, patronymic)).
//		Preload("PoliceOfficer2.Person",
//			db.Where("person.name = ? and  surname = ? and  patronymic = ?", name, surname, patronymic)).
//		Find(&crews).Error
//	if err != nil {
//		return nil, Error("GetOfficersCrewByOfficerId", err)
//
//	}
//	return crews, nil
//}
