package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	responseWriter(w, 200, struct{}{}) // use a double curly braces here
}