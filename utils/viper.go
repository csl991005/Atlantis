package utils

import "github.com/spf13/viper"

var (
	// Gin
	Mode string
	Port string

	// Gorm
	DbUser string
	DbPassword string
	DbHost string
	DbPort string
	DbName string
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("Config file not found")
		} else {
			panic("Config file was found but another error was produced")
		}
	}

	GetGin()
	GetGorm()
}


func GetGin()  {
	Mode = viper.GetString("Gin.Mode")
	Port = viper.GetString("Gin.Port")
}

func GetGorm(){
	DbUser =viper.GetString("Gorm.DbUser")
	DbPassword =viper.GetString("Gorm.DbPassword")
	DbHost =viper.GetString("Gorm.DbHost")
	DbPort =viper.GetString("Gorm.DbPort")
	DbName =viper.GetString("Gorm.DbName")
}