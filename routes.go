package main

import (
	"net/http"

	"github.com/ck46/ebus/api"
	"github.com/ck46/ebus/utils"

	"github.com/gorilla/mux"
)

func router() *mux.Router {
	r := mux.NewRouter()

	// serving static files
	staticFileDirectory := http.Dir(utils.AppFilePath("assets/"))

	staticFileHandler := http.StripPrefix("/assets/",
		http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	// serve a file
	r.HandleFunc("/favicon.ico", api.ServeFileHandler)
	r.HandleFunc("/robots.txt", api.ServeFileHandler)

	// Api Requests
	r.HandleFunc("/api/login", api.Login).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/signup", api.Signup).Methods("POST", "OPTIONS")

	return r
}
