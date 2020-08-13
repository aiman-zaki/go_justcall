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
		(*Profiling)(nil),
	}

	for _, model := range models {
		services.CreateTable(db, model)
	}

	var rw RoleWrapper
	rw.Single = Role{0, "ADMIN", time.Now(), time.Now()}
	rw.Create()
	rw.Single = Role{0, "USER", time.Now(), time.Now()}
	rw.Create()

}
