package transport

import (
	"api-rest/internal/core"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func GetAllHistories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(core.FlorComCulturaList)
}

func GetHistory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, historia := range core.FlorComCulturaList {
		if strconv.Itoa(int(historia.Id)) == id {
			json.NewEncoder(w).Encode(historia)
		}
	}
}
