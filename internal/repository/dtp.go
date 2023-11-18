package psql

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type DtpRepository interface {
	GetAllDtps(tx *gorm.DB) ([]*Dtp, error)
	GetDtpById(tx *gorm.DB, id int) (*Dtp, error)
	AddDtp(tx *gorm.DB, coords, street, area, metro, category string) (*Dtp, error)
	GetDtpByIdCars(tx *gorm.DB, dtpId int) ([]*Vehicle, error)
	AddParticipant(tx *gorm.DB, dtpId int, pts string, passport int, role string, lawNumber *string) (*ParticipantOfDtp, error)
	GetPoliceOfficerDtpsById(tx *gorm.DB, id int) ([]*Dtp, error)
	GetDtpByArea(tx *gorm.DB, area string) ([]*Dtp, error)
	GetCurrentDtpByPersonId(tx *gorm.DB, id int) (*Dtp, error)
	CloseDtp(tx *gorm.DB, id int) error
}

func (db Db) AddDtp(tx *gorm.DB, coords, street, area, metro, category string) (*Dtp, error) {
	dtp := Dtp{Coords: coords, Street: street, Area: area, Metro: metro, Category: category}
	result := tx.Select("Coords", "Street", "Area", "Category").Create(&dtp)
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
func (db Db) AddParticipant(tx *gorm.DB, dtpId int, pts string, passport int, role string, lawNumber *string) (*ParticipantOfDtp, error) {
	var err error
	if err != nil {
		description := fmt.Sprintf("repository:AddParticipant: %s", err.Error())
		return nil, errors.New(description)
	}
	var violation *Violation
	err = tx.Model(&violation).
		Where("law_number = ?", lawNumber).
		Find(&violation).Error
	if err != nil {
		description := fmt.Sprintf("repository:AddParticipant: %s", err.Error())
		return nil, errors.New(description)
	}
	var person Person
	err = tx.Model(&person).
		Where("passport = ?",
			passport).
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
	participant := &ParticipantOfDtp{
		DtpId:       dtpId,
		ViolationId: violation.Id,
		PersonId:    person.Id,
		VehicleId:   vehicle.Id,
		Role:        role,
	}
	err = tx.Model(&participant).Create(&participant).Error
	if err != nil {
		description := fmt.Sprintf("repository:AddParticipant: %s", err.Error())
		return nil, errors.New(description)
	}
	return participant, nil
}

func (db Db) GetPoliceOfficerDtpsById(tx *gorm.DB, id int) ([]*Dtp, error) {
	var dtps []*Dtp
	err := tx.Raw(`
	select dtp.id , dtp.date from dtp
	join crew_dtp on crew_dtp.dtp_id = dtp.id
	join crew on crew.id = crew_dtp.crew_id
	join (select * from police_officer where id = ?) as police_officer
	on crew.p_officer_id_1 = police_officer.id or crew.p_officer_id_2 = police_officer.id
	
`).
		Where("id = ?", id).
		Find(&dtps).Error
	if err != nil {
		description := fmt.Sprintf("repository:GetPoliceOfficerDtpsById: %s", err.Error())
		return nil, errors.New(description)
	}
	return dtps, nil
}
func (db Db) GetDtpByArea(tx *gorm.DB, area string) ([]*Dtp, error) {
	var dtps []*Dtp

	err := tx.Raw(`
		select dtp.id , dtp_description.time as date,dtp.category from (select * from dtp where area = ?) as dtp
	    join (select * from dtp_description where text = 'Зарегестрировано') as dtp_description
		on dtp_description.dtp_id = dtp.id
		`, area).
		Find(&dtps).Error
	if err != nil {
		description := fmt.Sprintf("repository:GetDtpByArea: %s", err.Error())
		return nil, errors.New(description)
	}
	fmt.Println()
	return dtps, nil
}

func (db Db) GetCurrentDtpByPersonId(tx *gorm.DB, id int) (*Dtp, error) {
	var dtp *Dtp
	//err := tx.Model(&dtp).
	//	Preload("Crews", db.Where("crew.duty = true")).
	//	Preload("Crews.PoliceOfficer", db.Where("police_officer.person_id = ?", id)).
	//	//Preload("DtpDescriptions", db.Where("text <> ?", ClosedDtp)).
	//	Find(&dtp).Error
	res := tx.Raw(`
		select dtp.id,dtp.coords,dtp.category,time as date from dtp 
	join crew_dtp on crew_dtp.dtp_id = dtp.id
	join (select * from crew where duty = true) as crew on crew.id = crew_dtp.crew_id
	join crew_police_officer on crew_police_officer.crew_id = crew.id
	join (select * from police_officer where person_id = ?) as police_officer on police_officer.id  = crew_police_officer.po_id
	join (select dtp_id,text from dtp_description where ? != dtp_description.text ) as dtp_description on dtp_description.dtp_id = dtp.id
	order by dtp.id desc limit 1
`, id, ClosedDtp).
		Find(&dtp)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	err := res.Error
	if err != nil {
		description := fmt.Sprintf("repository:GetCurrentDtpByPersonId: %s", err.Error())
		return nil, errors.New(description)
	}
	return dtp, nil
}
func (db Db) CloseDtp(tx *gorm.DB, id int) error {
	err := tx.Raw(`
	update crew set duty = false where crew.id in (
	    select crew.id from (select * from crew_dtp where dtp_id = ?) as crew_dtp join crew on crew.id = crew_dtp.crew_id

	)
`, id).Find(&DtpDescription{}).Error
	if err != nil {
		description := fmt.Sprintf("repository:CloseDtp: %s", err.Error())
		return errors.New(description)
	}
	return nil
}

//select crew.id from (select * from crew_dtp where dtp_id = 55) as crew_dtp join crew on crew.id = crew_dtp.crew_id
