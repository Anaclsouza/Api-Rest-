package internal

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConectionDb() (*gorm.DB, error){

conection:= "host=localhost user=postgres password=Cas1974c#2024 dbname=postgres port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
db, err := gorm.Open(postgres.Open(conection))
if err != nil {
	fmt.Println(err)
}
return db, err

}


