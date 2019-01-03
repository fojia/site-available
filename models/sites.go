package models


type Site struct {
	Id               int
	Site             string
	MetaTitle       string
	MetaDescription string
	MetaKeywords    string
}
type Information struct {
	SiteId int
	Status  int
}

//Get sites from database for checking
func GetSites() ([]*Site) {

	rows, err := DB.Query("SELECT `id`, `site` FROM `sites` WHERE `active`=1")
	CheckErr(err)

	defer rows.Close();
	sites := make([]*Site, 0)

	for rows.Next() {
		site := new(Site)
		err := rows.Scan(&site.Id, &site.Site)
		CheckErr(err)
		sites = append(sites, site)
	}
	return sites
}
//Add information to databse
func AddInformation(information *Information) {
	sth, err := DB.Prepare("INSERT INTO `informations` (`site_id`, `status`) values (?,?)")
	CheckErr(err)
	sth.Exec(information.SiteId, information.Status)

}