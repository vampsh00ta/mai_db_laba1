package service

import (
	"TgDbMai/internal/psql"
	"errors"
)

type Service interface {
	RegDtp(dtp *psql.Dtp) (int, error)
	VehicleDpts(pts string) ([]*psql.Dtp, error)
	AddDtpDescription(dtpId int, text string) (*psql.DtpDescription, error)
	GetVehicleByPts(pts string) (*psql.Vehicle, error)
	GetVehicleOwners(pts string) ([]*psql.Person, error)
}

func (s service) RegDtp(dtp *psql.Dtp) (int, error) {
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
func (s service) VehicleDpts(pts string) ([]*psql.Dtp, error) {
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

func (s service) AddDtpDescription(id int, text string) (*psql.DtpDescription, error) {
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

func (s service) GetVehicleByPts(pts string) (*psql.Vehicle, error) {
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
func (s service) GetVehicleOwners(pts string) ([]*psql.Person, error) {
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
