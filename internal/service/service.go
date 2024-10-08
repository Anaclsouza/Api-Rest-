package service

import (
	core"api-rest/internal/core"
	"fmt"

	"gorm.io/gorm"
)

// pega todos os produtos cadastrados no banco
func GetAll(db *gorm.DB) ([]core.HistoriaFlor, error) {

	var historias []core.HistoriaFlor

	if err := db.Table("produtos").Find(&historias).Error; err != nil {
		fmt.Println("Erro ao buscar os produtos:", err)

	}

	for _, historia := range historias {
		fmt.Println(historia)
	}
	return historias, nil
}

//func Insert(db*gorm.DB) error{
	//if core.Id == nil {

	//}
	//}
