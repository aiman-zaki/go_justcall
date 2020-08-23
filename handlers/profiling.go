package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aiman-zaki/go_justcall/models"
	"github.com/aiman-zaki/go_justcall/wrappers"
	"github.com/go-chi/chi"
)

type ProfilingResources struct{}

func (rs ProfilingResources) Routes() chi.Router {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Get("/rate/{id}", rs.ReadRateByUserID)
		r.Post("/", rs.Create)
		r.Get("/spec-rate/{spec_id}", rs.ReadSpecRate)
		r.Get("/comments/{id}", rs.ReadComments)

	})
	return r
}

//Create : implementation taken from appInsert.php
func (rs ProfilingResources) Create(w http.ResponseWriter, r *http.Request) {
	var pw models.ProfilingWrapper
	err := wrappers.JSONDecodeWrapper(w, r, &pw.Single)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	err = pw.Create()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	json.NewEncoder(w).Encode(pw.Single)
}

func (rs ProfilingResources) ReadRateByUserID(w http.ResponseWriter, r *http.Request) {
	var pw models.ProfilingWrapper
	id := chi.URLParam(r, "id")
	parsedID, err := strconv.Atoi(id)
	if err != nil {

		http.Error(w, err.Error(), 400)
		return
	}

	err = pw.ReadRateByUserID(int64(parsedID))
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(pw.Array)

}

func (rs ProfilingResources) ReadComments(w http.ResponseWriter, r *http.Request) {
	var pw models.ProfilingWrapper
	id := chi.URLParam(r, "id")
	parsedID, err := strconv.Atoi(id)
	if err != nil {

		http.Error(w, err.Error(), 400)
		return
	}

	err = pw.ReadComments(int64(parsedID))
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(pw.Array)

}

func (rs ProfilingResources) ReadSpecRate(w http.ResponseWriter, r *http.Request) {
	var pw models.ProfilingWrapper
	id := chi.URLParam(r, "spec_id")
	parsedID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	pw.Single.SpecID = int64(parsedID)
	res, err := pw.ReadSpecRate()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(res)
}
