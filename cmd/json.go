package main

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, body any) error{
	w.Header().Set("Content-Type","application/json")
	return json.NewEncoder(w).Encode(body)
}

func ReadJson(w http.ResponseWriter, r *http.Request, body any) error {
	maxBytes := 1_048_578
	r.Body = http.MaxBytesReader(w,r.Body,int64(maxBytes))
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(body)
}