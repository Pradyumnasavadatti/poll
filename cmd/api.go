package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Config struct{
	Addr string
	DBAddr string
}

type application struct{
	Addr string
	config *Config
}

func (app *application) mux() chi.Router{
	mux := chi.NewMux();

	mux.Route("/",func(r chi.Router){
		r.Get("/", app.healthBeat)
	})

	mux.Route("/v1",func(r chi.Router){
		r.Post("/createPoll",app.createPoll)
		r.Get("/getPoll",app.getPoll)
		r.Post("/votePoll",app.votePoll)
	})

	return mux;
}

func (app *application) run(mux chi.Router) error{
	server := &http.Server{
		Addr: app.Addr,
		Handler: mux,
	}

	fmt.Printf("ON: %s",app.Addr)
	err := server.ListenAndServe()

	return err
}