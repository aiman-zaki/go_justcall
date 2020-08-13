package models

import (
	"time"

	"github.com/aiman-zaki/go_justcall/services"
	"github.com/go-pg/pg/v9"
)

type Profiling struct {
	ID        int64   `json:"id"`
	Rate      float64 `json:"rate"`
	Shape     int64   `json:"shape"`
	Help      int16   `json:"help"`
	Comment   string  `json:"comment"`
	CallLogID int64   `json:"call_log_id"`
	CallLog   CallLog `pg:"fk:call_log_id" json:"call_log"`
	SpecID    int64   `json:"spec_id"`
	Spec      Spec    `pg:"fk:spec_id" json:"spec"`

	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}

type ProfilingWrapper struct {
	Single Profiling
	Array  []Profiling
}

func (tw *ProfilingWrapper) Create() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Insert(tw.Single)
	if err != nil {
		return err
	}
	return nil
}

func (tw *ProfilingWrapper) Update() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	_, err := db.Model(&tw.Single).WherePK().Update()
	if err != nil {
		return err
	}
	return nil
}
func (tw *ProfilingWrapper) Delete() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	_, err := db.Model(&tw.Single).WherePK().Delete()
	if err != nil {
		return err
	}
	return nil
}

func (tw *ProfilingWrapper) ReadRateByUserID(id int64) error {
	db := pg.Connect(services.PgOptions())
	db.AddQueryHook(services.DbLogger{})
	defer db.Close()
	err := db.Model(&tw.Array).Relation("CallLog").Where(`"call_log"."user_id" = ?`, id).Select()
	if err != nil {
		return err
	}
	return nil
}
