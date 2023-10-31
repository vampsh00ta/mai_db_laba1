package psql

import (
	"fmt"
	"gorm.io/gorm"
)

type CrewRepository interface {
	GetAllCrews(tx *gorm.DB) ([]*Crew, error)
	SetDutyCrews(tx *gorm.DB, duty bool, crew_id ...int) ([]*Crew, error)
	AddCrewsToDtp(tx *gorm.DB, duty int, crewIds ...int) ([]*Crew, error)
}

func (db Db) GetAllCrews(tx *gorm.DB) ([]*Crew, error) {
	var crews []*Crew
	//var a []*PoliceOfficer

	err := tx.Model(&Crew{}).
		Preload("PoliceOfficers.Person").
		Preload("Gai").
		Find(&crews).Error
	//err = tx.Model(&PoliceOfficer{}).
	//	Preload("Crews").
	//	Find(&a).Error
	if err != nil {
		return nil, Error("GetPersonInfoByPassport", err)

	}
	return crews, nil
}
func (db Db) SetDutyCrews(tx *gorm.DB, duty bool, crewIds ...int) ([]*Crew, error) {
	var crews []*Crew
	//var a []*PoliceOfficer
	vars := "("
	if len(crewIds) == 0 {
		return nil, nil
	}
	var input = []any{any(duty)}
	for _, id := range crewIds {
		input = append(input, id)
		vars += "?,"
	}
	vars = vars[0:len(vars)-1] + ")"
	err := tx.Raw(`
	update crew set duty = ? where id in `+vars+`
`, input...).
		Find(&crews).Error

	if err != nil {
		return nil, Error("GetPersonInfoByPassport", err)

	}
	return crews, nil
}

func (db Db) AddCrewsToDtp(tx *gorm.DB, dtpId int, crewIds ...int) ([]*Crew, error) {
	var crews []*Crew
	vars := ""
	if len(crewIds) == 0 {
		return nil, nil
	}
	input := []any{}
	for _, id := range crewIds {
		input = append(input, id)
		vars += fmt.Sprintf("(%d,?),", dtpId)
	}
	vars = vars[0 : len(vars)-1]
	fmt.Println(vars)
	err := tx.Raw(`
		insert into crew_dtp (dtp_id,crew_id) values `+vars+`
		`, input...).
		Find(&crews).Error
	//err = tx.Model(&PoliceOfficer{}).
	//	Preload("Crews").
	//	Find(&a).Error
	if err != nil {
		return nil, Error("GetPersonInfoByPassport", err)

	}
	return crews, nil
}
