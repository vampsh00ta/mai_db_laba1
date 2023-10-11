package controllers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func s() {
	//vehicle, err := rep.GetVehicleByPts("1488")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(vehicle.Model)
	//err = rep.RegisterVehicle(11111, vehicle)
	//fmt.Println(vehicle, err)
	//fmt.Println(vehicle.Model)
	//
	//persons, err := rep.GetVehicleOwners("2281488")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(persons[0].Name)
	//var dtp psql.Dtp
	//
	//err = db.Model(&psql.Dtp{}).
	//	Preload("Participants.Person.Vehicles").
	//	Preload("Crews.PoliceOfficer1.Person.Vehicles").
	//	Preload("Crews.PoliceOfficer2.Person.Vehicles").
	//	Preload("Participants.Vehicle").
	//	Preload("Participants.Violation").Find(&dtp).Error
	//if err != nil {
	//	fmt.Println(err)
	//}
	//for _, v := range dtp.Participants {
	//	fmt.Println(v.Person, v.Vehicle)
	//}
	//fmt.Println(dtp.Crews[0])
	//var crew psql.Crew
	//err = db.Model(&psql.Crew{}).Joins("join police_officer on police_officer.id = crew.p_officer_id_1 or police_officer.id = crew.p_officer_id_2").
	//	Joins("join person on person.id = police_officer.person_id").
	//	Joins("join person_vehicle on person_vehicle.person_id = person.id").
	//	Joins("join vehicle on person_vehicle.vehicle_id = vehicle.id").
	//	Preload("PoliceOfficer1").
	//	Where("crew.id= ?", 2).
	//	Scan(&crew).
	//	Error
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(crew)

	//db.Joins("PoliceOfficer", db.Where("id=?", 1)).
	//	Joins("JOIN emails ON emails.user_id = users.id AND emails.email = ?")

}

var dispacherActions = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Вызвать патруль", "call_patrul"),
	),
)
