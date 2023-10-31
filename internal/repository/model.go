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
	Dtp   *Dtp
}

func (DtpDescription) TableName() string {
	return "dtp_description"
}

type Crew struct {
	Id             int              `json:"id" db:"id,omitempty" gorm:"primaryKey"`
	PoliceOfficers []*PoliceOfficer `gorm:"many2many:crew_po;joinForeignKey:crew_id;joinReferences:po_id;"`
	Gai_id         int              `json:"gai_id" db:"gai_id" gorm:"index;column:gai_id;"`
	Gai            *Gai
	Dtps           []*Dtp    `gorm:"many2many:crew_dtp;joinForeignKey:crew_id;joinReferences:dtp_id"`
	Time           time.Time `gorm:"default:current_timestamp;"`
	Duty           bool
}

func (Crew) TableName() string {
	return "crew"
}

type Gai struct {
	Id     int     `json:"id" db:"id,omitempty" gorm:"primaryKey"`
	Area   string  `json:"area" db:"area"`
	Metro  string  `json:"metro" db:"metro"`
	Coords string  `json:"coords" db:"coords"`
	Crews  []*Crew `gorm:"foreignKey:gai_id;references:id"`
}

func (Gai) TableName() string {
	return "gai"
}

type Dtp struct {
	Id              int                 `json:"id" db:"id,omitempty" gorm:"primaryKey"`
	Coords          string              `json:"coords" db:"coords"`
	Date            time.Time           `json:"date" db:"date" gorm:"column:date"`
	Street          string              `json:"street" db:"street"`
	Area            string              `json:"area" db:"area"`
	Metro           string              `json:"metro" db:"metro"`
	Category        string              `json:"category" db:"category"`
	Crews           []*Crew             `gorm:"many2many:crew_dtp;joinForeignKey:dtp_id;joinReferences:crew_id"`
	Participants    []*ParticipantOfDtp `gorm:"foreignKey:DtpId;references:Id"`
	DtpDescriptions []*DtpDescription   `gorm:"foreignKey:DtpId;references:Id"`
}

func (Dtp) TableName() string {
	return "dtp"
}

type Violation struct {
	Id        int    `json:"id" db:"id,omitempty" gorm:"primaryKey"`
	LawNumber string `json:"law_number" db:"law_number"`
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
	Person   *Person `gorm:"foreignKey:PersonId"`
	Rank     string  `json:"rank" db:"rank"`
	GaiName  string  `json:"gai_name" db:"gai_name"`
	Crews    []*Crew `gorm:"many2many:crew_po;joinForeignKey:po_id;joinReferences:crew_id"`
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
	Fine              []*Fine        `gorm:"foreignKey:PersonId;references:Id"`
	PoliceOfficer     *PoliceOfficer `gorm:"foreignKey:PersonId;references:id"`
}

func (Person) TableName() string {
	return "person"
}

type Fine struct {
	Id          int `json:"id" db:"id,omitempty" gorm:"primaryKey"`
	PersonId    int `json:"personId" db:"personId" `
	Person      *Person
	Date        time.Time `json:"date" db:"date" `
	PaymentTime time.Time `json:"payment_time" db:"payment_time"`
	Amount      int       `json:"amount" db:"amount" `
	Reason      string    `json:"reason" db:"reason"`
}

func (Fine) TableName() string {
	return "fine"
}

//type ParticipantOfDtp struct {
//	Id        int    `json:"id" db:"id,omitempty"`
//	PersonId  int    `json:"person_id" db:"person_id"`
//	VehicleId int    `json:"vehicle_id" db:"vehicle_id"`
//	DtpId     int    `json:"dtp_id" db:"dtp_id"`
//	Role      string `json:"role" db:"role"`
//}
