package models

import (
	"time"

	"github.com/aiman-zaki/go_justcall/services"
	"github.com/go-pg/pg/v9"
)

func InitialDb() {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	models := []interface{}{
		(*Role)(nil),
		(*User)(nil),
		(*CallType)(nil),
		(*CallStatus)(nil),
		(*CallLog)(nil),
		(*Spec)(nil),
		(*Profilling)(nil),
	}

	for _, model := range models {
		services.CreateTable(db, model)
	}

	var rw RoleWrapper
	rw.Single = Role{0, "ADMIN", time.Now(), time.Now()}
	rw.Create()
	rw.Single = Role{0, "USER", time.Now(), time.Now()}
	rw.Create()

	var uw UserWrapper
	uw.Single = User{0, "AIMAN", "AIMAN", "aiman@test.com", "123456789", "01298422142", "", 1, &Role{}, time.Now(), time.Now(), "", ""}
	uw.Register()
	uw.Single = User{0, "AIMAN", "AIMAN", "aiman1@test.com", "123456789", "01298422142", "", 1, &Role{}, time.Now(), time.Now(), "", ""}
	uw.Register()

	var csw CallStatusWrapper

	csw.Single = CallStatus{0, "CALLER", time.Now(), time.Now()}
	csw.Create()

	csw.Single = CallStatus{0, "CALLEE", time.Now(), time.Now()}
	csw.Create()

	var ctw CallTypeWrapper

	ctw.Single = CallType{0, "MISSED", time.Now(), time.Now()}
	ctw.Create()
	ctw.Single = CallType{0, "ANSWERED", time.Now(), time.Now()}
	ctw.Create()

	var sw SpecWrapper

	sw.Single = Spec{0, "Abuse", "", time.Now(), time.Now()}
	sw.Create()
	sw.Single = Spec{0, "Conflict", "", time.Now(), time.Now()}
	sw.Create()
	sw.Single = Spec{0, "Death or Loss", "", time.Now(), time.Now()}
	sw.Create()
	sw.Single = Spec{0, "Serious Ilness", "", time.Now(), time.Now()}
	sw.Create()
	sw.Single = Spec{0, "Serious Problem", "", time.Now(), time.Now()}
	sw.Create()
	sw.Single = Spec{0, "Others", "", time.Now(), time.Now()}
	sw.Create()

}
