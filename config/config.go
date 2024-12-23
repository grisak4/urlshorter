package config

import (
	"log"

	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("[ERROR] %s", err.Error())
	}
}

type DBConf struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func GetDatabaseConf() DBConf {
	return DBConf{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetInt("database.port"),
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.pass"),
		DBName:   viper.GetString("database.dbname"),
	}
}
