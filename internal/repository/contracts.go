package repository

import "api-rest/internal/core"

type FlorCulturaRepository interface {
	Save(core.FlorComCultura) error
	Get(id int)(core.FlorComCultura, error)
	GetAll()([]core.FlorComCultura,error) 
}