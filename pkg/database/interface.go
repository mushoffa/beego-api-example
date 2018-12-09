package database

import (
	"beego-api-example/models"
)

type Model struct {
	Seed	models.Seed
	User	models.User
}

type RepositoryHandler interface {
	FindAll() ([]*Model, error)
	Insert(model Model) (*Model, error)
	RemoveById(id string) (error)
}