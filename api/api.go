package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"site-available/models"
)

type JSON interface {
}

//Listen and wait queries
func Listen() {
	router := mux.NewRouter()
	//Sites
	router.HandleFunc("/api/sites", getSites).Methods("GET")
	router.HandleFunc("/api/sites/{id}", getSite).Methods("GET")
	router.HandleFunc("/api/sites", createSite).Methods("POST")
	router.HandleFunc("/api/sites/{id}", deleteSite).Methods("DELETE")

	router.HandleFunc("/api/informaitons/{siteId}", getInformations).Methods("GET")
	http.ListenAndServe(":8080", router)
}

//Get inforations by id
func getInformations(w http.ResponseWriter, r *http.Request) {
	var infos []JSON
	params := mux.Vars(r)

	for _, value := range models.GetAllInformations(models.StringtoInt(params["siteId"])) {
		infos = append(infos, value)
	}
	responseJSON(w, infos)
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
	params := mux.Vars(r)
	models.DeleteSite(models.StringtoInt(params["id"]))
}

//Add site
func createSite(w http.ResponseWriter, r *http.Request) {

	site := new(models.Site)
	site.Site = r.FormValue("site")
	models.AddSite(site)
}

//Output json
func responseJSON(w http.ResponseWriter, body JSON) {
	json, err := json.Marshal(body)
	models.CheckErr(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

//Response error with 404 status
func ResponseError(w http.ResponseWriter, body JSON) {
	json, err := json.Marshal(body)
	models.CheckErr(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)
	w.Write(json)
}
