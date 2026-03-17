package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func errorResponse(w http.ResponseWriter, code int, msg string){
	if code > 499 {
		log.Printf("Server had some server error with code: %v and message: %v", code, msg)
	}
	type errResponse struct{
		Error string `json:"error"`
	}
	responseWriter(w, code, errResponse{
		Error:msg,
	})
}

func responseWriter(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed: Could'nt marshal %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}