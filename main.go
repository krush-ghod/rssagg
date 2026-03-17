package main

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"net/http"
)

func main() {

	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("You gotta do something")
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "UPDATE", "PUT", "OPTIONS", "DELETE"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	routerV2 := chi.NewRouter()

	routerV2.Get("/health", handlerReadiness) // to handle only get request, use the get method instead of the handleFunc method
	routerV2.Get("/err", handlerErrors)
	router.Mount("/v1", routerV2)

	svr := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}

	log.Printf("You got the server up at %v", portString)
	err := svr.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	// lets try to print the port
	fmt.Println("Port:", portString)
}