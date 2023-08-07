package models

import "gorm.io/gorm"

type Disk struct {
	gorm.Model
	Diskname   string
	CreateAt   string
	UpdateAt   string
	Userid     int64
	Bucketname string
	Bucketid   int64
}
