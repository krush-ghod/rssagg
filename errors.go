package main

import "net/http"

func handlerErrors(w http.ResponseWriter, r *http.Request) {
	errorResponse(w, 400, "You crewed up!")
}