package service

import (
	rep "TgDbMai/internal/repository"
	"errors"
)

type Service interface {
	RegDtp(dtp *rep.Dtp) (int, error)
	VehicleDpts(pts string) ([]*rep.Dtp, error)
	AddDtpDescription(dtpId int, text string) (*rep.DtpDescription, error)
	GetVehicleByPts(pts string) (*rep.Vehicle, error)
	GetVehicleOwners(pts string) ([]*rep.Person, error)
	Spravki
}

func (s service) RegDtp(dtp *rep.Dtp) (int, error) {
	tx := s.rep.GetDb().Begin()
	defer tx.Commit()
	dtp, err := s.rep.AddDtp(tx, dtp.Coords, dtp.Street, dtp.Area, dtp.Metro, dtp.Category)
	if err != nil {
		return 0, err
	}
	if tx.Error != nil {
		tx.Rollback()
		return 0, errors.New("transaction error")
	}
	return dtp.Id, nil
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
