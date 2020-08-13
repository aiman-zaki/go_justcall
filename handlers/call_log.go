package handlers

import (
	"github.com/go-chi/chi"
)

type CallLogResources struct{}

func (rs CallLogResources) Routes() chi.Router {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
	})
	return r
}
