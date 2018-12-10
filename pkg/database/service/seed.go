package service

import (
	//"reflect"
	// "beego-api-example/models"
	"beego-api-example/pkg/database"
	// log "github.com/sirupsen/logrus"
)

type SeedService struct {
	r 	database.RepositoryHandler
}

func NewSeedService(repository database.RepositoryHandler) *SeedService {
	return &SeedService{
		r: repository,
	}
}

func (s *SeedService) FindAll() ([]interface{}, error) {
	// res, err := s.r.FindAll()
	// log.Info(res)
	// return res,err
	return s.r.FindAll()
}

func (s *SeedService) Insert(model interface{}) (interface{}, error) {
	return s.r.Insert(model)
}

func (s *SeedService) RemoveById(id string) (error) {
	return s.r.RemoveById(id)
}