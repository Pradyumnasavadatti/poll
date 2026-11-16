package main

import (
	"net/http"
)

func (app * application) healthBeat(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("Dhuk dhuk"))
}