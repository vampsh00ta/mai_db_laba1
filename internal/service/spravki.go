package service

import (
	rep "TgDbMai/internal/repository"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Spravki interface {
	GetPersonVehicleByPts(pts, fio string) (*rep.Vehicle, error)
	GetPersonsVehiclesAndDtps(fio string) ([]*rep.Vehicle, error)
	GetPersonInfoFIO(fio string) ([]*rep.Person, error)
	GetPersonInfoPassport(passport string) ([]*rep.Person, error)
	GetOfficersInfo(fio string) (*rep.PoliceOfficer, []*rep.Dtp, error)
	//GetOfficersInfo(fio string) (*rep.PoliceOfficer, error)
	GetDtpsInfoNearArea(area string) ([]*rep.Dtp, error)
}

func (s service) GetPersonVehicleByPts(pts, fio string) (*rep.Vehicle, error) {
	tx := s.rep.GetDb().Begin()
	defer tx.Commit()
	fioData := strings.Split(fio, " ")
	vehicle, err := s.rep.GetPersonByPts(tx,
		strings.ToLower(fioData[0]),
		strings.ToLower(fioData[1]),
		strings.ToLower(fioData[2]), pts)
	if err != nil {
		return nil, err
	}
	if tx.Error != nil {
		tx.Rollback()
		return nil, errors.New("transaction error")
	}
	return vehicle, nil

}
func (s service) GetPersonsVehiclesAndDtps(fio string) ([]*rep.Vehicle, error) {
	tx := s.rep.GetDb().Begin()
	defer tx.Commit()
	fioData := strings.Split(fio, " ")
	vehicles, err := s.rep.GetPersonsVehicles(tx, strings.ToLower(fioData[0]),
		strings.ToLower(fioData[1]),
		strings.ToLower(fioData[2]))
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
func (s service) GetPersonInfoFIO(fio string) ([]*rep.Person, error) {
	tx := s.rep.GetDb().Begin()
	defer tx.Commit()
	fioData := strings.Split(fio, " ")
	person, err := s.rep.GetPersonInfoByFIO(tx, strings.ToLower(fioData[0]),
		strings.ToLower(fioData[1]),
		strings.ToLower(fioData[2]))
	if err != nil {
		return nil, err
	}
	if tx.Error != nil {
		tx.Rollback()
		return nil, errors.New("transaction error")
	}
	return person, nil
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

func (s service) GetOfficersInfo(fio string) (*rep.PoliceOfficer, []*rep.Dtp, error) {
	tx := s.rep.GetDb().Begin()
	defer tx.Commit()
	fioData := strings.Split(fio, " ")

	//officer, err := s.rep.GetOfficersInfoByFIO(tx, strings.ToLower(fioData[0]),
	//	strings.ToLower(fioData[1]),
	//	strings.ToLower(fioData[2]))
	officer, err := s.rep.GetOfficersInfoByFIO(tx, fioData[0],
		fioData[1],
		fioData[2])
	if err != nil {
		return nil, nil, err
	}
	//crews, err := s.rep.GetOfficersCrewByOfficerId(tx, officer.Id)
	//if err != nil {
	//	return nil, nil, err
	//}
	////officer.Crew = crews

	dtps, err := s.rep.GetPoliceOfficerDtpsById(tx, officer.Id)
	if err != nil {
		return nil, nil, err
	}

	if tx.Error != nil {
		tx.Rollback()
		return nil, nil, errors.New("transaction error")
	}

	return officer, dtps, nil

}

//	func (s service) GetOfficersInfo(fio string) (*rep.PoliceOfficer, []*rep.Dtp, error) {
//		tx := s.rep.GetDb().Begin()
//		defer tx.Commit()
//		fioData := strings.Split(fio, " ")
//
//		//officer, err := s.rep.GetOfficersInfoByFIO(tx, strings.ToLower(fioData[0]),
//		//	strings.ToLower(fioData[1]),
//		//	strings.ToLower(fioData[2]))
//		officer, err := s.rep.GetOfficerByFIO(tx, fioData[0],
//			fioData[1],
//			fioData[2])
//		if err != nil {
//			return nil, nil, err
//		}
//
//		var dtps []*rep.Dtp
//		for _, crew := range officer.Crews1 {
//			for _, dtp := range crew.Dtps {
//				dtps = append(dtps, dtp)
//			}
//		}
//
//		if tx.Error != nil {
//			tx.Rollback()
//			return nil, nil, errors.New("transaction error")
//		}
//		return officer, dtps, nil
//
// }
func (s service) GetDtpsInfoNearArea(area string) ([]*rep.Dtp, error) {
	tx := s.rep.GetDb().Begin()
	defer tx.Commit()
	dtps, err := s.rep.GetDtpByArea(tx, area)
	if err != nil {
		return nil, err
	}
	if tx.Error != nil {
		tx.Rollback()
		return nil, errors.New("transaction error")
	}
	return dtps, nil
}
func (s service) GetDtpsInfoRadiusMetro() {

}
