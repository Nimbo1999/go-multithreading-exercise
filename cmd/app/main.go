package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nimbo1999/go-multithreading-exercise/internal/infra/webserver/handler"
)

func main() {
	CepHandler := handler.CepHandler{}

	r := chi.NewRouter()
	r.Use(middleware.SetHeader("Content-Type", "application/json"))
	r.Use(middleware.Logger)
	r.Get("/{cep}", CepHandler.GetCep)
	http.ListenAndServe(":8000", r)
}
