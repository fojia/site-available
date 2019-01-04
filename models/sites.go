package models

type Site struct {
	Id              int
	Site            string
	MetaTitle       string
	MetaDescription string
	MetaKeywords    string
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

//Get site by id
func GetSite(Id string) (*Site) {
	row := DB.QueryRow("SELECT `id`, `site` FROM `sites` WHERE `active`=1 and `id`=?", Id)
	site := new(Site)
	err := row.Scan(&site.Id, &site.Site)
	CheckErr(err)

	return site
}

//Add site
func AddSite(site *Site) {
	sth, err := DB.Prepare("INSERT INTO 'sites' (`site`) VALUES (?) ")
	CheckErr(err)
	sth.Exec(site.Site)
}

//Delete site
func DeleteSite(Id int) {
	_, err := DB.Exec("DELETE FROM `sites` WHERE `id`=?", Id)
	CheckErr(err)

}
