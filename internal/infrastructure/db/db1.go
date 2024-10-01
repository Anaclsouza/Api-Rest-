package db

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	TimeZone = "America/Sao_Paulo"
)

type DB struct {
	Host    string
	User    string
	Pass    string
	Dbname  string
	Port    string
	Sslmode string
}

func configDB() *DB {
	config := &DB{
		Host:    viper.GetString("host"),
		User:    viper.GetString("user"),
		Pass:    viper.GetString("password"),
		Dbname:  viper.GetString("dbname"),
		Port:    viper.GetString("port"),
		Sslmode: viper.GetString("sslmode"),
	}

	return config
}

func NewFlorECulturaDB() (*gorm.DB, error) {

	config := configDB()

	conection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.Host,
		config.User,
		config.Pass,
		config.Dbname,
		config.Port,
		config.Sslmode,
		TimeZone)

	db, err := gorm.Open(postgres.Open(conection))
	if err != nil {
		fmt.Println(err, "database not connected")
	}
	return db, err

}
