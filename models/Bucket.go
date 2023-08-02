package models

import "gorm.io/gorm"

type Storage struct {
	gorm.Model
	Accesskey  string
	Alias      string
	Secretkey  string
	Bucketname string
	Zone       string
	Userid     int64
}
