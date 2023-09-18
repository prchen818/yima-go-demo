package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
)

func DatabaseInit() {
	d, err := gorm.Open(mysql.Open("root:mysql.peifang.para.party@(43.248.96.203:30306)/yima_user_db?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err.Error())
	}
	Db = d
}
