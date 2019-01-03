# Site available by golang

Site available checker (test project for learn golang, very "simple" code)

Basic database structure
 
```mysql
create table sites
(
	id int(11) NOT NULL AUTO_INCREMENT,
	site VARCHAR(255) default '',
	meta_name varchar(1000) default '',
	meta_description varchar(1000) default '',
	meta_keywords varchar(1000) default '',
	active tinyint default 0,
	created_at DATETIME default CURRENT_TIMESTAMP,
	updated_at DATETIME
	
);

create table informations
(
	id int(11) NOT NULL AUTO_INCREMENT,
	site_id int(11) default 0,
	status int default 0,
	created_at DATETIME default CURRENT_TIMESTAMP
);
```
