package models

type Information struct {
	Id     int
	SiteId int
	Status int
}

//Add information to databse
func AddInformation(information *Information) {
	_, err := DB.Exec("INSERT INTO `informations` (`site_id`, `status`) VALUES (?,?)", information.SiteId, information.Status)
	CheckErr(err)

}

//Get all rows with informations
func GetAllInformations(siteId int) []*Information {
	rows, err := DB.Query("SELECT `id`, `site_id`, `status` FROM `informations` WHERE `site_id`=? ORDER BY `created_at` DESC", siteId)
	CheckErr(err)
	infos := make([]*Information, 0)

	for rows.Next() {
		info := new(Information)
		err := rows.Scan(&info.Id, &info.SiteId, &info.Status)
		CheckErr(err)
		infos = append(infos, info)
	}
	
	return infos
}

//Get last information row
func GetLastInformation() *Information {
	row, err := DB.Query("SELECT `id` ,`site_id`, `status` FROM `informations` ORDER BY `created_at` DESC LIMIT 1 ")
	CheckErr(err)
	info := new(Information)
	err = row.Scan(&info.Id, &info.SiteId, &info.Status)
	CheckErr(err)

	return info
}
