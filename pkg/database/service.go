package database

import (
	"beego-api-example/pkg/database/mongo"
)

type Service struct {
	r			RepositoryHandler
	mongo.Session
}

func NewService(r RepositoryHandler) *Service {
	return &Service{
			r:r,
		}
}

func (s *Service) Find(model interface{}) (interface{}) {
	return s.r.Find(model)
}

func (s *Service) FindAll() (interface{}, error) {
	return s.r.FindAll() 
}

func (s *Service) FindById(id string) (interface{}, error) {
	return s.r.FindById(id)
}

func (s *Service) Insert(model interface{}) (interface{}, error) {

	return s.r.Insert(model)
}

func (s *Service) RemoveById(id string) (error) {
	return s.r.RemoveById(id)
}

func (s *Service) CloseSession() {
	s.Session.Close()
}

