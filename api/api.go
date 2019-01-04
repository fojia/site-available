package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"site-available/models"
)

type JSON interface {
}

type Site struct {
	id   int    `json:"id,omitempty"`
	site string `json:"site,omitempty"`
}

//Listen and wait queries
func Listen() {
	router := mux.NewRouter()
	router.HandleFunc("/api/sites", getSites).Methods("GET")
	router.HandleFunc("/api/sites/create", createSite).Methods("POST")
	router.HandleFunc("/api/sites/{id}", getSite).Methods("GET")
	router.HandleFunc("/api/sites/{id}", deleteSite).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}

//Get all sites
func getSites(w http.ResponseWriter, r *http.Request) {
	var sites []JSON
	for _, value := range models.GetSites() {
		sites = append(sites, value)

	}
	responseJSON(w, sites)
}

//Get site by id
func getSite(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	s := models.GetSite(params["id"])
	responseJSON(w, s)
}

//Delete site
func deleteSite(w http.ResponseWriter, r *http.Request) {

}

//Add site
func createSite(w http.ResponseWriter, r *http.Request) {

}

//Output json
func responseJSON(w http.ResponseWriter, body JSON) {
	json, err := json.Marshal(body)
	models.CheckErr(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
