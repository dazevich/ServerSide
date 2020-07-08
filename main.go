package main

import (
	"net/http"

	"./api"
	"github.com/gorilla/mux"
)

func main() {

	port := ":9097"
	r := mux.NewRouter()
	r.HandleFunc("/getCourses", api.APIServer)
	r.HandleFunc("/getCrypto", api.GetCrypto)
	http.ListenAndServe(port, r)

}
