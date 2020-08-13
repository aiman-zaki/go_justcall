package main

import (
	"log"
	"net/http"

	"github.com/aiman-zaki/go_justcall/handlers"
	"github.com/aiman-zaki/go_justcall/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog"
)

func main() {
	models.InitialDb()

	r := chi.NewRouter()
	logger := httplog.NewLogger("httplog", httplog.Options{
		JSON: true,
	})
	r.Use(httplog.RequestLogger(logger))
	r.Use(middleware.Heartbeat("/ping"))
	r.Mount("/api/auth", handlers.AuthResources.Routes(handlers.AuthResources{}))
	r.Mount("/api/user", handlers.UserResources.Routes(handlers.UserResources{}))
	r.Mount("/api/profilling", handlers.ProfilingResources.Routes(handlers.ProfilingResources{}))
	r.Mount("/api/call-log", handlers.CallLogResources.Routes(handlers.CallLogResources{}))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	})
	log.Fatal(http.ListenAndServe(":8181", c.Handler(r)))
}
