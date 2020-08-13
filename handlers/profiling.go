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

	})
	return r
}

//Create : implementation taken from appInsert.php
func (rs ProfilingResources) Create(w http.ResponseWriter, r *http.Request) {
	var pw models.ProfilingWrapper
	err := wrappers.JSONDecodeWrapper(w, r, &pw.Single)
	if err != nil {
		return
	}
	err = pw.Create()
	if err != nil {
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
