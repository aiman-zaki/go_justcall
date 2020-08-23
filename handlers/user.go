package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/aiman-zaki/go_justcall/models"
	"github.com/aiman-zaki/go_justcall/wrappers"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
)

type UserResources struct{}

func (rs UserResources) Routes() chi.Router {
	r := chi.NewRouter()
	//r.Use(jwtauth.Verifier(models.TokenSetting()))
	//r.Use(jwtauth.Authenticator)
	r.Route("/", func(r chi.Router) {
		r.Put("/profile", rs.UpdateProfile)
		r.Get("/profile", rs.ReadProfile)
		r.Post("/profile/picture/{id}", rs.UpdateProfilePicture)
	})
	return r
}

func (rs UserResources) ReadPublicProfileByID(w http.ResponseWriter, r *http.Request) {
	var uw models.UserWrapper
	id := chi.URLParam(r, "id")
	parsedID, err := strconv.Atoi(id)
	if err != nil {

		http.Error(w, err.Error(), 400)
		return
	}
	uw.Single.ID = int64(parsedID)

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

func (rs UserResources) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var uw models.UserWrapper

	jwtToken := r.Header.Get("Authorization")
	split := strings.Split(jwtToken, " ")
	tokenAuth := models.TokenSetting()
	t, err := tokenAuth.Decode(split[1])
	if err != nil {
		return
	}
	id := t.Claims.(jwt.MapClaims)["user"]
	fmt.Println(id)
	err = wrappers.JSONDecodeWrapper(w, r, &uw.Single)

	uw.Single.ID = int64(id.(float64))
	fmt.Println(uw.Single)
	err = uw.Update()
	if err != nil {
		return
	}
	fmt.Println(uw.Single)
	json.NewEncoder(w).Encode(uw.Single)
}

func CreateFolder(id string) {
	_, err := os.Stat("./static/users/" + id)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll("./static/users/"+id, 0755)
		if errDir != nil {
			log.Fatal(err)
		}

	}
}

func (rs UserResources) UpdateProfilePicture(w http.ResponseWriter, r *http.Request) {
	var uw models.UserWrapper
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("image")
	id := chi.URLParam(r, "id")
	CreateFolder(id)
	parsedID, _ := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Fprintf(w, "%v", handler.Header)
	f, err := os.OpenFile("./assets/users/"+id+"/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	uw.Single.ID = int64(parsedID)
	err = uw.Read()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	uw.Single.Photo = handler.Filename
	err = uw.Update()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}
