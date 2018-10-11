package api

import (
	"flag"
	"fmt"
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
	router.HandleFunc("/api/v1/user/create", CreateUserEndpoint).Methods("POST")
	router.HandleFunc("/api/v1/user/authenticate", AuthenticateUserEndpoint).Methods("POST")
	router.HandleFunc("/api/v1/user/me", MeUserEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), router))
}

func RespondError(w http.ResponseWriter, err error) {
	w.WriteHeader(400) // 400 is pretty generic but meh
	fmt.Fprintf(w, err.Error())
}
