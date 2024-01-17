package model

import (
	"fmt"

	"github.com/csl991005/Atlantis/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDb() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassword,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})
}
