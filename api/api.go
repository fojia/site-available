package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"site-available/models"
)

type Site struct {
	id   int    `json:"id"`
	site string `json:"site"`
}

//Listen and wait queries
func Listen() {
	router := mux.NewRouter()
	router.HandleFunc("/sites", getSites).Methods("GET")
	http.ListenAndServe(":8080", router)
}

func getSites(w http.ResponseWriter, r *http.Request) {
	var sites []Site
	for _, value := range models.GetSites() {
		sites = append(sites, Site{value.Id, value.Site})
	}
	json.NewEncoder(w).Encode(sites)
}
