package models

import (
	"time"

	"github.com/aiman-zaki/go_justcall/services"
	"github.com/go-pg/pg/v9"
)

type Role struct {
	ID   int64  `json:"id"`
	Role string `json:"role"`

	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}

type RoleWrapper struct {
	Single Role
	Array  []Role
}

func (tw *RoleWrapper) Create() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Insert(&tw.Single)
	if err != nil {
		return err
	}
	return nil
}
func (tw *RoleWrapper) Read() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&tw.Single).Select()
	if err != nil {
		return err
	}
	return nil
}
func (tw *RoleWrapper) Update() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	_, err := db.Model(&tw.Single).WherePK().Update()
	if err != nil {
		return err
	}
	return nil
}
func (tw *RoleWrapper) Delete() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	_, err := db.Model(&tw.Single).WherePK().Delete()
	if err != nil {
		return err
	}
	return nil
}
