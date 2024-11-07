package service

import (
	core "api-rest/internal/core"
	r "api-rest/internal/repository"
	"fmt"
)

type Service struct {
	db r.Repository
}

func NewService(db r.Repository) Service {
	return Service{db: db}
}

func (s *Service) Save(NovaFlor core.FlorComCultura) (core.FlorComCultura, error) {

	exists, err := s.checkFlorExists(NovaFlor)
	if err != nil {
		return core.FlorComCultura{}, err
	}

	if exists {

		florEncontrada, err := s.db.Get(int(NovaFlor.Id))
		if err != nil {
			return core.FlorComCultura{}, err
		}

		florEncontrada.Cuidados = NovaFlor.Cuidados
		florEncontrada.Descricao = NovaFlor.Descricao

		err = s.db.Create(florEncontrada)
		if err != nil {
			return core.FlorComCultura{}, err
		}
	}

	err = s.db.Create(NovaFlor)
	if err != nil {
		return core.FlorComCultura{}, err
	}

	return NovaFlor, nil
}

func (s *Service) checkFlorExists(NovaFlor core.FlorComCultura) (bool, error) {

	todasFlores, err := s.db.GetAll()
	if err != nil {
		fmt.Println("Nenhuma flor encontrada")
	}

	for _, flor := range todasFlores {
		if flor.Id == NovaFlor.Id || flor.Nome == NovaFlor.Nome {
			return true, nil
		}
	}

	return false, err
}

// pega todos os produtos cadastrados no banco

func Get(s *Service) ([]core.FlorComCultura, error) {

	todasFlores, err := s.db.GetAll()
	if err != nil {
		return []core.FlorComCultura{}, nil
	}
	return todasFlores, nil

}

//func Insert(db*gorm.DB) error{
//if core.Id == nil {

//}
//}
