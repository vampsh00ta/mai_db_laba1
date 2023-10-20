package service

import (
	"TgDbMai/internal/psql"
	"errors"
	"fmt"
	"strings"
)

type Spravki interface {
	IsPersonOwner(pts, fio string) (bool, error)
	GetPersonsVehiclesAndDtps(fio string) ([]*psql.Vehicle, error)
}

func (s service) IsPersonOwner(pts, fio string) (bool, error) {
	tx := s.rep.GetDb().Begin()
	defer tx.Commit()
	fioData := strings.Split(fio, " ")
	vehicle, err := s.rep.GetPersonByPts(tx, fioData[0], fioData[1], fioData[2], pts)
	if err != nil {
		return false, err
	}
	if tx.Error != nil {
		tx.Rollback()
		return false, errors.New("transaction error")
	}
	return vehicle.Id != 0, nil

}
func (s service) GetPersonsVehiclesAndDtps(fio string) ([]*psql.Vehicle, error) {
	tx := s.rep.GetDb().Begin()
	defer tx.Commit()
	fioData := strings.Split(fio, " ")
	vehicles, err := s.rep.GetPersonsVehicles(tx, fioData[0], fioData[1], fioData[2])
	if err != nil {
		return nil, err
	}
	var ptss []string
	for _, vehicle := range vehicles {
		ptss = append(ptss, vehicle.Pts)
	}
	dtps, err := s.rep.GetVehicleDptsByPts(tx, ptss...)
	if err != nil {
		return nil, err
	}
	fmt.Println(dtps)
	if tx.Error != nil {
		tx.Rollback()
		return nil, errors.New("transaction error")
	}
	return vehicles, nil

}
func (s service) GetPersonInfo() {

}
func (s service) GetOfficersInfo() {

}
func (s service) GetDtpsInfoNearMetro() {

}
func (s service) GetDtpsInfoRadiusMetro() {

}
