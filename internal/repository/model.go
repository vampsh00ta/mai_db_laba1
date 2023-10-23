package psql

import "time"

//	type Gai struct {
//		Id         int    `json:"id" db:"id,omitempty"`
//		Name       string `json:"name" db:"name,"`
//		License    string `json:"license" db:"license"`
//		Department string `json:"department" db:"Department"`
//	}

type ParticipantOfDtp struct {
	Id          int       `json:"id" db:"id,omitempty" gorm:"primaryKey"`
	DtpId       int       `json:"dtp_id" db:"dtp_id" gorm:"index"`
	PersonId    int       `json:"person_id" db:"person_id" gorm:"index"`
	Person      Person    `gorm:"foreignKey:id;references:person_id"`
	VehicleId   int       `json:"vehicle_id" db:"vehicle_id" gorm:"index"`
	Vehicle     Vehicle   `gorm:"foreignKey:id;references:vehicle_id"`
	ViolationId int       `json:"violation_id" db:"violation_id" gorm:"index"`
	Violation   Violation `gorm:"foreignKey:id;references:violation_id"`
	Role        string    `json:"role" db:"role"`
	Dtp         *Dtp
}

func (ParticipantOfDtp) TableName() string {
	return "participant_of_dtp"
}

type DtpDescription struct {
	Id    int       `json:"id" db:"id,omitempty" gorm:"primaryKey"`
	Text  string    `json:"text" db:"text"`
	Time  time.Time `json:"time" db:"time"`
	DtpId int       `json:"dtp_id" db:"dtp_id" gorm:"index"`
}

func (DtpDescription) TableName() string {
	return "dtp_description"
}

type Crew struct {
	Id               int            `json:"id" db:"id,omitempty" gorm:"primaryKey"`
	PoliceOfficerId1 int            `json:"p_officer_id_1" db:"p_officer_id_1" gorm:"index;column:p_officer_id_1" `
	PoliceOfficerId2 int            `json:"p_officer_id_2" db:"p_officer_id_2" gorm:"index;column:p_officer_id_2"`
	PoliceOfficer1   *PoliceOfficer `gorm:"foreignKey:PoliceOfficerId1;references:id"`
	PoliceOfficer2   *PoliceOfficer `gorm:"foreignKey:PoliceOfficerId2;references:id"`
	Gai_id           int            `json:"gai_id" db:"gai_id" gorm:"index;column:gai_id"`
	Gai              *Gai           `gorm:"foreignKey:gai_id;references:id"`
	Dtps             []*Dtp         `gorm:"many2many:crew_dtp;"`
	Time             time.Time      `gorm:"many2many:crew_dtp;default:current_timestamp"`
	Duty             bool
}

func (Crew) TableName() string {
	return "crew"
}

type Gai struct {
	Id    int    `json:"id" db:"id,omitempty" gorm:"primaryKey"`
	Area  string `json:"area" db:"area"`
	Metro string `json:"metro" db:"metro"`
}
type Dtp struct {
	Id           int                 `json:"id" db:"id,omitempty" gorm:"primaryKey"`
	Coords       string              `json:"coords" db:"coords"`
	Date         time.Time           `json:"date" db:"date" gorm:"column:date"`
	Street       string              `json:"street" db:"street"`
	Area         string              `json:"area" db:"area"`
	Metro        string              `json:"metro" db:"metro"`
	Category     string              `json:"category" db:"category"`
	Crews        []*Crew             `gorm:"many2many:crew_dtp"`
	Participants []*ParticipantOfDtp `gorm:"foreignKey:DtpId;references:Id"`
}

func (Dtp) TableName() string {
	return "dtp"
}

type Violation struct {
	Id        int    `json:"id" db:"id,omitempty" gorm:"primaryKey"`
	LawNumber int    `json:"law_number" db:"law_number"`
	Law       string `json:"law" db:"law"`
}

func (Violation) TableName() string {
	return "violation"
}

type Vehicle struct {
	Id       int       `json:"id" db:"id,omitempty" gorm:"primaryKey"`
	Pts      string    `json:"pts" db:"pts"`
	Model    string    `json:"model" db:"model"`
	Category string    `json:"category" db:"category"`
	Persons  []*Person `gorm:"many2many:person_vehicle;"`
}

func (Vehicle) TableName() string {
	return "vehicle"
}

type PoliceOfficer struct {
	Id       int     `json:"id" db:"id,omitempty" gorm:"primaryKey"`
	PersonId int     `gorm:"index"`
	Person   *Person `gorm:"foreignKey:PersonId;references:id"`
	Rank     string  `json:"rank" db:"rank"`
	GaiName  string  `json:"gai_name" db:"gai_name"`
}

func (PoliceOfficer) TableName() string {
	return "police_officer"
}

type Person struct {
	Id                int        `json:"id" db:"id,omitempty" gorm:"primaryKey"`
	Name              string     `json:"name" db:"name" `
	Surname           string     `json:"surname" db:"surname" `
	Patronymic        string     `json:"patronymic" db:"patronymic" `
	Birthday          time.Time  `json:"birthday" db:"birthday"`
	Passport          int        `json:"passport" db:"passport"`
	Citizenship       string     `json:"citizenship" db:"citizenship"`
	Vehicles          []*Vehicle `gorm:"many2many:person_vehicle;"`
	ParticipantsOfDtp []*ParticipantOfDtp
}

func (Person) TableName() string {
	return "person"
}

//type ParticipantOfDtp struct {
//	Id        int    `json:"id" db:"id,omitempty"`
//	PersonId  int    `json:"person_id" db:"person_id"`
//	VehicleId int    `json:"vehicle_id" db:"vehicle_id"`
//	DtpId     int    `json:"dtp_id" db:"dtp_id"`
//	Role      string `json:"role" db:"role"`
//}
