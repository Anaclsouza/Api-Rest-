package service

import (
	"api-rest/internal/core"

	"gorm.io/gorm"
)

//pega todos os produtos cadastrados no banco
func GetAll(db *gorm.DB) ([]core.Produtos, error){
	var produto []core.Produtos
	produtoToGet := db.Table("produtos").Find(&produto)

	return produto, produtoToGet.Error

}