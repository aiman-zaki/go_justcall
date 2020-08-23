package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aiman-zaki/go_justcall/models"
	"github.com/aiman-zaki/go_justcall/wrappers"
	"github.com/go-chi/chi"
)

type AuthResources struct{}

func (rs AuthResources) Routes() chi.Router {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Post("/register", rs.Register)
		r.Post("/login", rs.Login)
		r.Get("/jwt/{refreshToken}", rs.RefreshToken)
	})
	return r
}

func (rs AuthResources) Register(w http.ResponseWriter, r *http.Request) {
	var uw models.UserWrapper
	wrappers.JSONDecodeWrapper(w, r, &uw.Single)
	fmt.Println(uw.Single)
	err := uw.Register()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	json.NewEncoder(w).Encode(uw.Single)
}

func (rs AuthResources) Login(w http.ResponseWriter, r *http.Request) {
	var aw models.UserWrapper
	wrappers.JSONDecodeWrapper(w, r, &aw.Single)
	fmt.Println(aw.Single)
	err, httpError := aw.Login()
	if err != nil {
		if httpError == 401 {
			http.Error(w, err.Error(), 401)
			return
		}
		http.Error(w, err.Error(), 400)
		return
	}
	json.NewEncoder(w).Encode(aw.Single)
}

func (rs AuthResources) RefreshToken(w http.ResponseWriter, r *http.Request) {
	refreshToken := chi.URLParam(r, "refreshToken")
	var aw models.UserWrapper
	//wrappers.JSONDecodeWrapper(w, r, &aw.Auth)
	aw.Single.RefreshToken = refreshToken
	err := aw.RefreshToken()
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(aw.Single)

}
