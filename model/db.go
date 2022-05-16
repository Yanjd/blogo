package model

import (
	"blogo/utils"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var (
	db  *gorm.DB
	err error
)

func InitDb() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4", utils.DbUser, utils.DbPassword,
		utils.DbHost, utils.DbPort, utils.DbName)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})

	if err != nil {
		fmt.Println("gorm open err: ", err)
		return
	}

	sqlDb, err := db.DB()
	if err != nil {
		fmt.Println("get sqlDb err: ", err)
		return
	}
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Second * 10)
	err = db.AutoMigrate(&User{}, &Category{}, &Article{})
	if err != nil {
		fmt.Println("auto migrate err: ", err)
		return
	}
}
