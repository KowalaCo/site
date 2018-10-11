package api

import (
	"flag"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//StartRouter Starts http server routing to defined routes
func StartRouter() {
	port := flag.Int("port", 1567, "The port to run the api on!")
	flag.Parse()
	log.Print("Starting api on port ", *port)
	router := mux.NewRouter().StrictSlash(true)
	//router.HandleFunc("/api/v2/user/create", CreateUserEndpoint).Methods("POST")
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), router))
}