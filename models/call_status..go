package models

import (
	"time"

	"github.com/aiman-zaki/go_justcall/services"
	"github.com/go-pg/pg/v9"
)

type CallStatus struct {
	ID     int64  `json:"id"`
	Status string `json:"status"`

	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}

type CallStatusWrapper struct {
	Single CallStatus
	Array  []CallStatus
}

func (tw *CallStatusWrapper) Create() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Insert(&tw.Single)
	if err != nil {
		return err
	}
	return nil
}
func (tw *CallStatusWrapper) Read() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&tw.Single).Select()
	if err != nil {
		return err
	}
	return nil
}
func (tw *CallStatusWrapper) Update() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	_, err := db.Model(&tw.Single).WherePK().Update()
	if err != nil {
		return err
	}
	return nil
}
func (tw *CallStatusWrapper) Delete() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	_, err := db.Model(&tw.Single).WherePK().Delete()
	if err != nil {
		return err
	}
	return nil
}
