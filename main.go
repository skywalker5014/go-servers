package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/go-chi/cors"
)

func main() {
	//read the .env file in hte root dir and load all the values inside it to the server at runtime
	godotenv.Load(".env")
	//create a new http router using chi package
    router := chi.NewRouter()
	//read the "PORT" value from the .env file passed to the server
	portString := os.Getenv("PORT")
    //cors implementation
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET","POST","PUT","DELETE","OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))
	//creating another router 
	v1Router := chi.NewRouter()
	//
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handleErr)
	router.Mount("/v1", v1Router)
	//creating the server instance 
	srv := &http.Server{
		Handler: router,
		Addr: ":"+portString,
	}

	log.Printf("Server starting at port %v", portString)
	//starting the server instance while handling error
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}