package service

import (
	rep "TgDbMai/internal/repository"
	"errors"
	"strconv"
	"strings"
)

type Person interface {
	GetPersonInfoFIO(fio string) ([]*rep.Person, error)
	GetPersonInfoPassport(passport string) ([]*rep.Person, error)
	GetOfficersInfo(fio string) (*rep.Person, []*rep.Dtp, error)
	GetVehicleOwners(pts string) ([]*rep.Person, error)
	AddParticipant(dtpId int, pts string, passport int, role string, lawNumber *string) (*rep.ParticipantOfDtp, error)
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
func (s service) GetOfficersInfo(fio string) (*rep.Person, []*rep.Dtp, error) {
	tx := s.rep.GetDb().Begin()
	defer tx.Commit()
	fioData := strings.Split(fio, " ")
	if len(fioData) != 3 {
		return nil, nil, errors.New("неправильные данные")
	}

	officer, err := s.rep.GetOfficersInfoByFIO(tx, fioData[0],
		fioData[1],
		fioData[2])
	if err != nil {
		return nil, nil, err
	}

	if tx.Error != nil {
		tx.Rollback()
		return nil, nil, errors.New("transaction error")
	}

	return officer, nil, nil

}
func (s service) GetPersonInfoPassport(passport string) ([]*rep.Person, error) {
	tx := s.rep.GetDb().Begin()
	defer tx.Commit()
	passportInt, _ := strconv.Atoi(passport)
	person, err := s.rep.GetPersonInfoByPassport(tx, passportInt)
	if err != nil {
		return nil, err
	}
	if tx.Error != nil {
		tx.Rollback()
		return nil, errors.New("transaction error")
	}
	return []*rep.Person{person}, nil
}
func (s service) GetPersonInfoFIO(fio string) ([]*rep.Person, error) {
	tx := s.rep.GetDb().Begin()
	defer tx.Commit()
	fioData := strings.Split(fio, " ")
	if len(fioData) != 3 {
		return nil, errors.New("неправильные данные")
	}
	person, err := s.rep.GetPersonInfoByFIO(tx, fioData[0],
		fioData[1],
		fioData[2])
	if err != nil {
		return nil, err
	}
	if tx.Error != nil {
		tx.Rollback()
		return nil, errors.New("transaction error")
	}
	return person, nil
}
