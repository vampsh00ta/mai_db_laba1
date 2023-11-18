package service

import (
	rep "TgDbMai/internal/repository"
	"errors"
	"strings"
)

type Vehicle interface {
	GetVehicleByPts(pts string) (*rep.Vehicle, error)
	GetPersonVehicleByPts(pts, fio string) (*rep.Vehicle, error)
	GetPersonsVehicles(fio string) ([]*rep.Vehicle, error)
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

func (s service) GetPersonsVehicles(fio string) ([]*rep.Vehicle, error) {
	tx := s.rep.GetDb().Begin()
	defer tx.Commit()
	fioData := strings.Split(fio, " ")
	if len(fioData) != 3 {
		return nil, errors.New("неправильные данные")
	}
	vehicles, err := s.rep.GetPersonsVehicles(tx, fioData[0],
		fioData[1],
		fioData[2])
	if err != nil {
		return nil, err
	}
	var ptss []string
	for _, vehicle := range vehicles {
		ptss = append(ptss, vehicle.Pts)
	}

	if tx.Error != nil {
		tx.Rollback()
		return nil, errors.New("transaction error")
	}
	return vehicles, nil

}
func (s service) GetPersonVehicleByPts(pts, fio string) (*rep.Vehicle, error) {
	tx := s.rep.GetDb().Begin()
	defer tx.Commit()
	fioData := strings.Split(fio, " ")
	if len(fioData) != 3 {
		return nil, errors.New("неправильные данные")
	}
	vehicle, err := s.rep.GetPersonByPts(tx,
		fioData[0],
		fioData[1],
		fioData[2], pts)
	if err != nil {
		return nil, err
	}
	if tx.Error != nil {
		tx.Rollback()
		return nil, errors.New("transaction error")
	}
	return vehicle, nil

}
