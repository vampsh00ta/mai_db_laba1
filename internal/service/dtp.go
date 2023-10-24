package service

import (
	rep "TgDbMai/internal/repository"
	"errors"
)

type Dtp interface {
	RegDtp(dtp *rep.Dtp, officerCount int) (*rep.Dtp, error)
	VehicleDpts(pts string) ([]*rep.Dtp, error)
	AddDtpDescription(dtpId int, text string) (*rep.DtpDescription, error)
	GetVehicleByPts(pts string) (*rep.Vehicle, error)
	GetVehicleOwners(pts string) ([]*rep.Person, error)
}

func (s service) RegDtp(dtp *rep.Dtp, officerCount int) (*rep.Dtp, error) {
	tx := s.rep.GetDb().Begin()
	defer tx.Commit()
	dtp, err := s.rep.AddDtp(tx, dtp.Coords, dtp.Street, dtp.Area, dtp.Metro, dtp.Category)
	if err != nil {
		return nil, err
	}
	crews, err := s.FindClosedCrews(dtp.Coords)
	if err != nil {
		return nil, err
	}
	var crewsToCall []*rep.Crew
	count := 0
	for i := 0; count < officerCount && i < len(crews); i++ {
		crewsToCall = append(crewsToCall, crews[i])
		count += len(crews[i].PoliceOfficers)
	}
	dtp.Crews = crews
	err = tx.Model(&dtp).Association("Crews").Append(&crews)
	if err != nil {
		return nil, err
	}

	if tx.Error != nil {
		tx.Rollback()
		return nil, errors.New("transaction error")
	}

	return dtp, nil
}
func (s service) VehicleDpts(pts string) ([]*rep.Dtp, error) {
	tx := s.rep.GetDb().Begin()
	defer tx.Commit()
	dpts, err := s.rep.GetVehicleDptsByPts(tx, pts)
	if err != nil {
		return nil, err
	}
	if tx.Error != nil {
		return nil, errors.New("transaction error")
	}
	return dpts, nil

}

func (s service) AddDtpDescription(id int, text string) (*rep.DtpDescription, error) {
	tx := s.rep.GetDb().Begin()
	defer tx.Commit()
	dtpDescription, err := s.rep.AddDtpDescription(tx, id, text)
	if err != nil {
		return nil, err
	}
	if tx.Error != nil {
		return nil, errors.New("transaction error")
	}
	return dtpDescription, nil
}

func (s service) GetVehicleByPts(pts string) (*rep.Vehicle, error) {
	tx := s.rep.GetDb().Begin()
	defer tx.Commit()
	vehicle, err := s.rep.GetVehicleByPts(tx, pts)
	if err != nil {
		return nil, err
	}
	if tx.Error != nil {
		return nil, errors.New("transaction error")
	}
	return vehicle, nil
}
func (s service) GetVehicleOwners(pts string) ([]*rep.Person, error) {
	tx := s.rep.GetDb().Begin()
	defer tx.Commit()
	persons, err := s.rep.GetVehicleOwners(tx, pts)
	if err != nil {
		return nil, err
	}
	if tx.Error != nil {
		return nil, errors.New("transaction error")
	}
	return persons, nil
}
