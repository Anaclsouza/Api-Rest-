package service

import (
	core"api-rest/internal/core"
	"fmt"

	"gorm.io/gorm"
)

// pega todos os produtos cadastrados no banco
func GetAll(db *gorm.DB) ([]core.Produtos, error) {

	var produtos []core.Produtos

	if err := db.Table("produtos").Find(&produtos).Error; err != nil {
		fmt.Println("Erro ao buscar os produtos:", err)

	}

	for _, produto := range produtos {
		fmt.Println(produto)
	}
	return produtos, nil
}

//func Insert(db*gorm.DB) error{
	//if core.Id == nil {

	//}
	//}
