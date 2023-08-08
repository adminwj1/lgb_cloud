package models

import "gorm.io/gorm"

type Catalogue struct {
	gorm.Model
	Diskname   string `gorm:"unique"`
	Userid     int64
	Bucketname string
	Bucketid   int64
}
