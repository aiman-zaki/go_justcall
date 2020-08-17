package models

import (
	"time"

	"github.com/aiman-zaki/go_justcall/services"
	"github.com/go-pg/pg/v9"
)

type CallLog struct {
	ID           int64     `json:"id"`
	UserID       int64     `json:"user_id"`
	User         *User     `pg:"fk:user_id" json:"user"`
	CallID       int64     `json:"call_id"`
	TimeCalled   time.Time `json:"time_called"`
	CallDuration string    `json:"call_duration"`

	CallStatusID int64       `json:"call_status_id"`
	CallStatus   *CallStatus `pg:"fk:call_type_id" json:"call_status"`

	CallTypeID int64    `json:"call_type_id"`
	CallType   CallType `pg:"fk:call_type_id" json:"call_type"`

	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}

type CallLogWrapper struct {
	Single CallLog
	Array  []CallLog
}

func (tw *CallLogWrapper) Create() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Insert(tw.Single)
	if err != nil {
		return err
	}
	return nil
}
func (tw *CallLogWrapper) Read() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&tw.Single).Select()
	if err != nil {
		return err
	}
	return nil
}

func (tw *CallLogWrapper) ReadByUserID() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&tw.Array).Where(`user_id = ?`, tw.Single.UserID).Relation("User").Select()
	if err != nil {
		return err
	}
	return nil
}

func (tw *CallLogWrapper) Update() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	_, err := db.Model(&tw.Single).WherePK().Update()
	if err != nil {
		return err
	}
	return nil
}
func (tw *CallLogWrapper) Delete() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	_, err := db.Model(&tw.Single).WherePK().Delete()
	if err != nil {
		return err
	}
	return nil
}
