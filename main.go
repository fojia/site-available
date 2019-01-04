package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"site-available/api"
	"site-available/models"
	"time"
)

func main() {
	models.InitDB()
	go api.Listen()

	//Run checking available site once by hour
	for t := range time.NewTicker(time.Hour).C {
		go startCheck(t)
	}
}

//Function for starting check sitese
func startCheck(t time.Time) {
	fmt.Println(t)
	for _, site := range models.GetSites() {
		checkSite(site)
	}
}



//Check site status and in future parse meta information
func checkSite(site *models.Site) {
	response, err := http.Get(site.Site)
	models.CheckErr(err)
	defer response.Body.Close()
	info := new(models.Information)
	info.SiteId = site.Id
	info.Status = response.StatusCode
	models.AddInformation(info)
}




