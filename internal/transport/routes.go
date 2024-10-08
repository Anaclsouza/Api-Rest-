package transport

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"

)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", Home)
	r.HandleFunc("/api/historias/{id}", GetHistory).Methods("Get")
	r.HandleFunc("/api/historias", GetAllHistories).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))
}