package service

import (
	rep "TgDbMai/internal/repository"
	"errors"
	"fmt"
	"github.com/umahmood/haversine"
	"strconv"
	"strings"
)

type Dtp interface {
	RegDtp(dtp *rep.Dtp, officerCount int) (*rep.Dtp, error)
	GetDtpsInfoNearArea(area string) ([]*rep.Dtp, error)
	VehicleDpts(pts string) ([]*rep.Dtp, error)
	AddDescriptionToDtp(dtpId int, text string) (*rep.DtpDescription, error)
	GetCurrentDtp(id int) (*rep.Dtp, error)
	GetDtpsInfoRadius(radius int, coords string) ([]*rep.Dtp, error)
	CloseDtp(id int) error
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
func (s service) GetDtpsInfoRadius(radius int, coords string) ([]*rep.Dtp, error) {
	tx := s.rep.GetDb().Begin()
	defer tx.Commit()
	dtps, err := s.rep.GetAllDtps(tx)

	if err != nil {
		return nil, err
	}

	splited := strings.Split(coords, ",")
	lat, err := strconv.ParseFloat(splited[0], 64)
	lon, err := strconv.ParseFloat(splited[1], 64)
	searchCoord := haversine.Coord{Lat: lat, Lon: lon} // Oxford, UK
	result := make([]*rep.Dtp, 0)
	for _, dtp := range dtps {

		splited := strings.Split(dtp.Coords, ",")
		lat, err := strconv.ParseFloat(splited[0], 64)
		lon, err := strconv.ParseFloat(splited[1], 64)

		if err != nil {
			return nil, err
		}
		crewCoord := haversine.Coord{Lat: lat, Lon: lon} // Turin, Italy
		_, km := haversine.Distance(searchCoord, crewCoord)
		if int(km) < radius {
			result = append(result, dtp)

		}
	}
	if tx.Error != nil {
		tx.Rollback()
		return nil, errors.New("transaction error")
	}
	return result, nil
}
