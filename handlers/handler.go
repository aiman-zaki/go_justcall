package handlers

import "github.com/go-chi/chi"

type JwtResources struct {
}

func (rs JwtResources) Routes() chi.Router {
	r := chi.NewRouter()

	return r
}
