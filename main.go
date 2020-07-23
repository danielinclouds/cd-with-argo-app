package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/status/{status:[0-9]+}", status).Methods("GET", "HEAD")
	rtr.HandleFunc("/version", version).Methods("GET", "HEAD")

	http.Handle("/", rtr)

	log.Println("Listening on port 8080")
	log.Println(os.Getenv("APP_VERSION"))
	http.ListenAndServe(":8080", nil)
}

func status(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	status, _ := strconv.Atoi(params["status"])
	w.WriteHeader(status)
	w.Write([]byte(params["status"]))
}

func version(w http.ResponseWriter, r *http.Request) {
	version := os.Getenv("APP_VERSION")
	w.WriteHeader(200)
	w.Write([]byte(version))
}
