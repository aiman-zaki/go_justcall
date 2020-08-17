package models

import (
	"time"

	"github.com/aiman-zaki/go_justcall/services"
	"github.com/go-pg/pg/v9"
)

type Spec struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`

	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}

type SpecRate struct {
	SpecID    int64   `json:"spec_id"`
	SumRate   float64 `json:"sum_rate"`
	CountRate int64   `json:"count_rate"`
	CallLogID int64   `json:"call_log_id"`
	UserID    int64   `json:"user_id"`
	CallID    int64   `json:"call_id"`
}

type SpecWrapper struct {
	Single Spec
	Array  []Spec
}

func (tw *SpecWrapper) Create() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Insert(&tw.Single)
	if err != nil {
		return err
	}
	return nil
}
func (tw *SpecWrapper) Read() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&tw.Single).Select()
	if err != nil {
		return err
	}
	return nil
}
func (tw *SpecWrapper) Update() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	_, err := db.Model(&tw.Single).WherePK().Update()
	if err != nil {
		return err
	}
	return nil
}
func (tw *SpecWrapper) Delete() error {
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	_, err := db.Model(&tw.Single).WherePK().Delete()
	if err != nil {
		return err
	}
	return nil
}
