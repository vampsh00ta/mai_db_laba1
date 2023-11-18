package psql

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type DescriptionRepository interface {
	GetLastUpdateByDtpId(tx *gorm.DB, dtpId int) (*DtpDescription, error)
	AddDescriptionToDtp(tx *gorm.DB, dtpId int, text string) (*DtpDescription, error)
}

const (
	CallCops         = "Зарегестрировано"
	TestOfAlcDrugs   = "Проверили  на алкоголь/наркотики"
	DoneProtocol     = "Написали протоколы"
	HandOverToPolice = "Передали дело сотрудникам полиции"
	ClosedDtp        = "Завершено"
)

func (db Db) AddDescriptionToDtp(tx *gorm.DB, dtpId int, text string) (*DtpDescription, error) {
	var dtp Dtp
	err := tx.Model(&dtp).
		Where("id = ?", dtpId).
		Find(&dtp).Error
	if err != nil {
		description := fmt.Sprintf("repository:AddDtpDescription: %s", err.Error())
		return nil, errors.New(description)
	}
	var dtpDescription DtpDescription
	dtpDescription = DtpDescription{
		Time:  time.Now(),
		Text:  text,
		DtpId: dtpId,
	}
	fmt.Println(dtpDescription)
	err = tx.Model(&dtpDescription).Create(&dtpDescription).Error
	if err != nil {
		description := fmt.Sprintf("repository:AddDtpDescription: %s", err.Error())
		return nil, errors.New(description)
	}
	return &dtpDescription, nil
}
func (db Db) GetLastUpdateByDtpId(tx *gorm.DB, dtpId int) (*DtpDescription, error) {
	var desc *DtpDescription

	res := tx.Raw(`
	select * from dtp_description where dtp_id = ? order by time desc
`, dtpId).Find(&desc)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	err := res.Error
	if err != nil {
		description := fmt.Sprintf("repository:GetDtpById: %s", err.Error())
		return nil, errors.New(description)
	}
	return desc, nil
}
