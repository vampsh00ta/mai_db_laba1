package psql

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type DtpRepository interface {
	GetAllDtps(tx *gorm.DB) ([]*Dtp, error)
	GetDtpById(tx *gorm.DB, id int) (*Dtp, error)
	AddDtp(tx *gorm.DB, coords, street, area, metro, category string) (*Dtp, error)
	AddDtpDescription(tx *gorm.DB, dtpId int, text string) (*DtpDescription, error)
	GetDtpByIdCars(tx *gorm.DB, dtpId int) ([]*Vehicle, error)
	AddParticipant(tx *gorm.DB, lawNumber string, dtpId int, pts, role string, personData Person) (*ParticipantOfDtp, error)
	GetPoliceOfficerDtpsById(tx *gorm.DB, id int) ([]*Dtp, error)
	GetDtpByArea(tx *gorm.DB, area string) ([]*Dtp, error)
	//select id from violation  where law_number = ?;
	//select  id  from person where name  = ? and surname  = ? and  patronymic = ? and passport = ?;
	//insert into participant_of_dtp (violation_id,vehicle_id , person_id , dtp_id , role)
	//values (?,?,?,?);
	//DeleteDtp() error
}

func (db Db) AddDtp(tx *gorm.DB, coords, street, area, metro, category string) (*Dtp, error) {
	dtp := Dtp{Coords: coords, Street: street, Area: area, Metro: metro, Category: category}
	result := tx.Select("Coords", "Street", "Area", "Metro", "Category").Create(&dtp)
	if result.Error != nil {
		description := fmt.Sprintf("repository:AddDtp: %s", result.Error.Error())
		return nil, errors.New(description)
	}

	return &dtp, nil

}

func (db Db) GetAllDtps(tx *gorm.DB) ([]*Dtp, error) {
	var dtps []*Dtp
	err := tx.Model(&dtps).
		Find(&dtps).Error
	if err != nil {
		description := fmt.Sprintf("repository:GetAllDtps: %s", err.Error())
		return nil, errors.New(description)
	}
	return dtps, nil
}

func (db Db) GetDtpById(tx *gorm.DB, id int) (*Dtp, error) {
	var dtp Dtp
	err := tx.Model(&dtp).
		Where("id = ?", id).
		Find(&dtp).Error
	if err != nil {
		description := fmt.Sprintf("repository:GetDtpById: %s", err.Error())
		return nil, errors.New(description)
	}
	return &dtp, nil
}
func (db Db) AddDtpDescription(tx *gorm.DB, dtpId int, text string) (*DtpDescription, error) {
	var dtp Dtp
	err := tx.Model(&dtp).
		Where("id = ?", dtpId).
		Find(&dtp).Error
	if err != nil {
		description := fmt.Sprintf("repository:AddDtpDescription: %s", err.Error())
		return nil, errors.New(description)
	}
	var dtpDescription DtpDescription
	dtpDescription = DtpDescription{
		Time:  time.Now(),
		Text:  text,
		DtpId: dtpId,
	}
	err = tx.Model(&dtpDescription).Create(&dtpDescription).Error
	if err != nil {
		description := fmt.Sprintf("repository:AddDtpDescription: %s", err.Error())
		return nil, errors.New(description)
	}
	return &dtpDescription, nil
}
func (db Db) GetDtpByIdCars(tx *gorm.DB, dtpId int) ([]*Vehicle, error) {
	var vehicles []*Vehicle

	err := tx.Table("(?) as dtp", db.Model(&Dtp{}).Where("id = ?", dtpId)).
		Select("Model", "Pts", "Vehicle.Category").
		Joins("join participant_of_dtp on  participant_of_dtp.dtp_id = dtp.id").
		Joins("join person on person.id = participant_of_dtp.person_id").
		Joins("join person_vehicle on person_vehicle.person_id = person.id").
		Joins("join  vehicle on person_vehicle.vehicle_id = vehicle.id").
		Find(&vehicles).Error
	if err != nil {
		description := fmt.Sprintf("repository:GetDtpByIdCars: %s", err.Error())
		return nil, errors.New(description)
	}
	return vehicles, nil
}
func (db Db) AddParticipant(tx *gorm.DB, lawNumber string, dtpId int, pts, role string, personData Person) (*ParticipantOfDtp, error) {
	var err error

	var dtp Dtp
	err = tx.Model(&dtp).
		Where("id = ?", dtpId).
		Find(&dtp).Error
	if err != nil {
		description := fmt.Sprintf("repository:AddParticipant: %s", err.Error())
		return nil, errors.New(description)
	}

	var violation Violation
	err = tx.Model(&violation).
		Where("law_number = ?", lawNumber).
		Find(&violation).Error
	if err != nil {
		description := fmt.Sprintf("repository:AddParticipant: %s", err.Error())
		return nil, errors.New(description)
	}
	var person Person
	err = tx.Model(&person).
		Where("name  = ? and surname  = ? and  patronymic = ? and passport = ?",
			personData.Name, personData.Surname, personData.Patronymic, personData.Passport).
		Find(&person).Error
	if err != nil {
		description := fmt.Sprintf("repository:AddParticipant: %s", err.Error())
		return nil, errors.New(description)
	}

	vehicle, err := db.GetVehicleByPts(tx, pts)
	if err != nil {
		description := fmt.Sprintf("repository:AddParticipant: %s", err.Error())
		return nil, errors.New(description)
	}

	participantOfDtp := ParticipantOfDtp{
		DtpId:       dtpId,
		ViolationId: violation.Id,
		PersonId:    person.Id,
		VehicleId:   vehicle.Id,
		Role:        role,
	}

	err = tx.Model(&participantOfDtp).Create(&participantOfDtp).Error
	if err != nil {
		description := fmt.Sprintf("repository:AddParticipant: %s", err.Error())
		return nil, errors.New(description)
	}
	return &participantOfDtp, nil
}

func (db Db) GetPoliceOfficerDtpsById(tx *gorm.DB, id int) ([]*Dtp, error) {
	var dtps []*Dtp
	err := tx.Raw(`
	select dtp.id , dtp.date from dtp
	join crew_dtp on crew_dtp.dtp_id = dtp.id
	join crew on crew.id = crew_dtp.crew_id
	join (select * from police_officer where id = 1) as police_officer
	on crew.p_officer_id_1 = police_officer.id or crew.p_officer_id_2 = police_officer.id
	
`).
		Where("id = ?", id).
		Find(&dtps).Error
	if err != nil {
		description := fmt.Sprintf("repository:GetDtpById: %s", err.Error())
		return nil, errors.New(description)
	}
	return dtps, nil
}
func (db Db) GetDtpByArea(tx *gorm.DB, area string) ([]*Dtp, error) {
	var dtps []*Dtp
	err := tx.Model(&dtps).
		Where("area = ?", area).
		Find(&dtps).Error
	if err != nil {
		description := fmt.Sprintf("repository:GetDtpById: %s", err.Error())
		return nil, errors.New(description)
	}
	return dtps, nil
}
