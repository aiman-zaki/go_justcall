package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aiman-zaki/go_justcall/models"
	"github.com/go-chi/chi"
)

type CallLogResources struct{}

func (rs CallLogResources) Routes() chi.Router {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Get("/{id}", rs.ReadWithUserID)
	})
	return r
}

func (rs CallLogResources) ReadWithUserID(w http.ResponseWriter, r *http.Request) {
	var pw models.CallLogWrapper
	id := chi.URLParam(r, "id")
	parsedID, err := strconv.Atoi(id)
	if err != nil {

		http.Error(w, err.Error(), 400)
		return
	}
	pw.Single.UserID = int64(parsedID)

	err = pw.ReadByUserID()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(pw.Array)

}
