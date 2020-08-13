package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/aiman-zaki/go_justcall/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

type UserResources struct{}

func (rs UserResources) Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(jwtauth.Verifier(models.TokenSetting()))
	r.Use(jwtauth.Authenticator)
	r.Route("/", func(r chi.Router) {
		r.Get("/profile", rs.ReadProfile)
	})
	return r
}

func (rs UserResources) ReadProfile(w http.ResponseWriter, r *http.Request) {
	jwtToken := r.Header.Get("Authorization")
	split := strings.Split(jwtToken, " ")
	tokenAuth := models.TokenSetting()
	t, err := tokenAuth.Decode(split[1])
	if err != nil {
		return
	}
	id := t.Claims.(jwt.MapClaims)["user"]
	fmt.Println(id)

	var uw models.UserWrapper
	uw.Single.ID = int64(id.(float64))

	err = uw.ReadProfile()
	if err != nil {
		return
	}
	fmt.Println(uw.Single)
	json.NewEncoder(w).Encode(uw.Single)
}
