package service

import (
	rep "TgDbMai/internal/repository"
	"errors"
	"fmt"
)

type Dtp interface {
	RegDtp(dtp *rep.Dtp, officerCount int) (*rep.Dtp, error)
	VehicleDpts(pts string) ([]*rep.Dtp, error)
	AddDescriptionToDtp(dtpId int, text string) (*rep.DtpDescription, error)
	GetVehicleByPts(pts string) (*rep.Vehicle, error)
	GetVehicleOwners(pts string) ([]*rep.Person, error)
	GetCurrentDtp(id int) (*rep.Dtp, error)
	CloseDtp(id int) error
	AddParticipant(dtpId int, pts string, passport int, role string, lawNumber *string) (*rep.ParticipantOfDtp, error)
	IssueFine(passport int, amount int, reason string) (*rep.Person, error)
	GetFines(passport int) (*rep.Person, error)
}

func (s service) RegDtp(dtp *rep.Dtp, officerCount int) (*rep.Dtp, error) {
	tx := s.rep.GetDb().Begin()
	defer tx.Commit()
	dtp, err := s.rep.AddDtp(tx, dtp.Coords, dtp.Street, dtp.Area, dtp.Metro, dtp.Category)
	if err != nil || tx.Error != nil {
		tx.Rollback()
		return nil, err
	}
	crews, err := s.FindClosedCrews(dtp.Coords)
	if err != nil {
		return nil, err
	}
	crewIds := []int{}
	for _, crew := range crews {
		crewIds = append(crewIds, crew.Id)
	}

	_, err = s.rep.AddDescriptionToDtp(tx, dtp.Id, rep.CallCops)
	if err != nil || tx.Error != nil {
		tx.Rollback()
		return nil, err
	}
	_, err = s.rep.SetDutyCrews(tx, true, crewIds...)
	if err != nil || tx.Error != nil {
		tx.Rollback()
		return nil, err
	}
	var crewsToCall []*rep.Crew
	count := 0
	for i := 0; count < officerCount && i < len(crews); i++ {
		crewsToCall = append(crewsToCall, crews[i])
		count += len(crews[i].PoliceOfficers)
	}
	_, err = s.rep.AddCrewsToDtp(tx, dtp.Id, crewIds...)
	if err != nil || tx.Error != nil {
		tx.Rollback()
		return nil, err
	}
	dtp.Crews = crews

	return dtp, nil
}
func (s service) VehicleDpts(pts string) ([]*rep.Dtp, error) {
	tx := s.rep.GetDb().Begin()
	defer tx.Commit()
	dpts, err := s.rep.GetVehicleDptsByPts(tx, pts)

	if err != nil || tx.Error != nil {
		return nil, err
	}
	if tx.Error != nil {
		return nil, errors.New("transaction error")
	}
	return dpts, nil

}

func (s service) AddDescriptionToDtp(id int, text string) (*rep.DtpDescription, error) {
	tx := s.rep.GetDb().Begin()
	defer tx.Commit()
	if id == 0 {
		return nil, nil
	}
	dtpDescription, err := s.rep.AddDescriptionToDtp(tx, id, text)
	if err != nil || tx.Error != nil {
		return nil, err
	}
	return dtpDescription, nil
}

func (s service) GetVehicleByPts(pts string) (*rep.Vehicle, error) {
	tx := s.rep.GetDb().Begin()
	defer tx.Commit()
	vehicle, err := s.rep.GetVehicleByPts(tx, pts)
	if err != nil || tx.Error != nil {
		return nil, err
	}
	return vehicle, nil
}
func (s service) GetVehicleOwners(pts string) ([]*rep.Person, error) {
	tx := s.rep.GetDb().Begin()
	defer tx.Commit()
	persons, err := s.rep.GetVehicleOwners(tx, pts)
	if err != nil || tx.Error != nil {
		return nil, err
	}
	return persons, nil
}

func (s service) GetCurrentDtp(id int) (*rep.Dtp, error) {
	tx := s.rep.GetDb().Begin()
	defer tx.Commit()
	dtp, err := s.rep.GetCurrentDtpByPersonId(tx, id)
	if err != nil || tx.Error != nil {
		return nil, err
	}
	if dtp == nil {
		return nil, nil
	}
	update, err := s.rep.GetLastUpdateByDtpId(tx, dtp.Id)
	dtp.DtpDescriptions = append(dtp.DtpDescriptions, update)

	return dtp, nil
}

func (s service) CloseDtp(id int) error {
	tx := s.rep.GetDb().Begin()
	defer tx.Commit()
	err := s.rep.CloseDtp(tx, id)
	fmt.Println(err)
	if err != nil {
		return err
	}
	if tx.Error != nil {
		return errors.New("transaction error")
	}
	return nil
}
func (s service) AddParticipant(dtpId int, pts string, passport int, role string, lawNumber *string) (*rep.ParticipantOfDtp, error) {
	tx := s.rep.GetDb().Begin()
	defer tx.Commit()
	parpicipant, err := s.rep.AddParticipant(tx, dtpId, pts, passport, role, lawNumber)
	if err != nil || tx.Error != nil {
		return nil, err
	}
	if err != nil || tx.Error != nil {
		return nil, err
	}
	return parpicipant, nil
}
func (s service) IssueFine(passport int, amount int, reason string) (*rep.Person, error) {
	tx := s.rep.GetDb().Begin()
	defer tx.Commit()
	person, err := s.rep.IssueFine(tx, passport, amount, reason)
	if err != nil || tx.Error != nil {
		return nil, err
	}

	return person, nil
}
func (s service) GetFines(passport int) (*rep.Person, error) {
	tx := s.rep.GetDb().Begin()
	defer tx.Commit()
	person, err := s.rep.GetFines(tx, passport)
	if err != nil || tx.Error != nil {
		return nil, err
	}

	return person, nil
}
