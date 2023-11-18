package service

import (
	psql "TgDbMai/internal/repository"
	"github.com/umahmood/haversine"
	"sort"
	"strconv"
	"strings"
)

type Crew interface {
	FindClosedCrews(coords string) ([]*psql.Crew, error)
	SetDutyCrews(duty bool, crew_id ...int) ([]*psql.Crew, error)
}

type _toSort struct {
	i, diff int
}

func (s service) FindClosedCrews(coords string) ([]*psql.Crew, error) {
	splited := strings.Split(coords, ",")
	lat, err := strconv.ParseFloat(splited[0], 64)
	lon, err := strconv.ParseFloat(splited[1], 64)

	if err != nil {

		return nil, err
	}
	tx := s.rep.GetDb()

	defer tx.Commit()
	crews, err := s.rep.GetAllCrews(tx)
	if err != nil || tx.Error != nil {
		tx.Rollback()
		return nil, err
	}

	dtpCoord := haversine.Coord{Lat: lat, Lon: lon} // Oxford, UK
	var toSort []_toSort
	var result []*psql.Crew
	for i, crew := range crews {
		if crew.Duty {
			continue
		}
		splited := strings.Split(crew.Gai.Coords, ",")
		lat, err := strconv.ParseFloat(splited[0], 64)
		lon, err := strconv.ParseFloat(splited[1], 64)

		if err != nil {
			return nil, err
		}
		crewCoord := haversine.Coord{Lat: lat, Lon: lon} // Turin, Italy
		_, km := haversine.Distance(dtpCoord, crewCoord)
		toSort = append(toSort, _toSort{i, int(km)})
		result = append(result, crew)
	}
	sort.Slice(result, func(i, j int) bool {
		return toSort[i].diff < toSort[j].diff
	})

	return result, nil
}
func (s service) SetDutyCrews(duty bool, crewIds ...int) ([]*psql.Crew, error) {
	tx := s.rep.GetDb().Begin()
	defer tx.Commit()
	crews, err := s.rep.SetDutyCrews(tx, duty, crewIds...)
	if err != nil || tx.Error != nil {
		tx.Rollback()
		return nil, err
	}

	return crews, nil
}
