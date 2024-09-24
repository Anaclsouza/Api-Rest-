package internal

import (
	"fmt"

	db "api-rest/internal/infraestucture"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	TimeZone = "America/Sao_Paulo"
)

func ConectionDb(config *db.Config) (*gorm.DB, error) {

	conection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", config.Host,config.User, config.Password, config.DbName, config.Port, config.Sslmode, TimeZone)
	db, err := gorm.Open(postgres.Open(conection))
	if err != nil {
		fmt.Println(err, "database not connected")
	}
	return db, err

}
