package database

import (

)

type RepositoryHandler interface {
	Find(interface{}) (interface{})
	FindAll() ([]interface{}, error)
	FindById(id string) (interface{}, error)
	Insert(model interface{}) (interface{}, error)
	RemoveById(id string) (error)
}