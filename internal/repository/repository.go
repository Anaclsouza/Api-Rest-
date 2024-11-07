package repository

import (
	"api-rest/internal/core"

	"gorm.io/gorm"
)

const (
	TabelaDB = "produtos"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{db: db}
}

func (r *Repository) Create(NovaFlor core.FlorComCultura) error {
	return r.db.Table(TabelaDB).Create(NovaFlor).Error

}


func (r *Repository) Get(FlorID int) (core.FlorComCultura, error) {
	var Flor core.FlorComCultura

	r.db.Table(TabelaDB).First(&Flor, FlorID)
	return Flor, nil

}

func (r *Repository) GetAll() ([]core.FlorComCultura, error) {
	var Flores []core.FlorComCultura

	r.db.Table(TabelaDB).Find(&Flores)

	return Flores, nil

}
