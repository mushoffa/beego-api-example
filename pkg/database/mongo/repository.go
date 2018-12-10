package mongo

import (
	"errors"

	"gopkg.in/mgo.v2/bson"
	mgo "gopkg.in/mgo.v2"
)

type MongoRepository struct {
	DB			*Database
	collection 	*mgo.Collection
	session		*mgo.Session
}

func NewMongoRepository(db *Database, collection string) *MongoRepository {
	c := db.Session.DB(db.Name).C(collection)
	return &MongoRepository {
		DB: db,
		collection: c,
	}
}

func (r *MongoRepository) FindAll() ([]interface{}, error) {

	res := []interface{}{}
	
	err := r.collection.Find(nil).Sort("name").All(&res)

	switch err {
	case nil:
		return res, nil
	case mgo.ErrNotFound:
		return res, errors.New("Error, not found")
	default:
		return res, err
	}
}

func (r *MongoRepository) FindById(id string) (interface{}, error) {
	var res interface{}

	_id := bson.ObjectIdHex(id)

	err := r.collection.Find(bson.M{"_id": _id}).One(&res)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *MongoRepository) Insert(model interface{}) (interface{}, error) {

	err := r.collection.Insert(&model)

	if err != nil {
		return model, err
	}

	return model, nil
}

func (r *MongoRepository) RemoveById(id string) (error) {
	_id := bson.ObjectIdHex(id)

	err := r.collection.Remove(bson.M{"_id": _id})

	if err != nil {
		return err
	}

	return nil
}

