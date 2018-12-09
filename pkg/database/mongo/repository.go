package mongo

import (
	"errors"
	mgo "gopkg.in/mgo.v2"

	_ "beego-api-example/pkg/database"
	"beego-api-example/models"
)

type MongoDbRepository struct {
	DB			*Database
	collection 	*mgo.Collection
	session		*mgo.Session
}

// func NewMongoRepository(db *Database,s *mgo.Session, c string) *MongoDbRepository {
// 	coll := s.DB(db).C(c)
// 	return &MongoDbRepository {
// 		db: db,
// 		collection: coll,
// 		session: s,
// 	}
// }

func (r *MongoDbRepository) FindAll() ([]*models.Seed, error) {
	var res []*models.Seed
	//coll := r.session.DB(r.db).C(r.collection)
	err := r.collection.Find(nil).Sort("name").All(&res)

	switch err {
	case nil:
		return res, nil
	case mgo.ErrNotFound:
		return nil, errors.New("Error, not found")
	default:
		return nil, err
	}
}

func (r *MongoDbRepository) Insert(b *struct{}) (*struct{}, error) {

	// coll := r.session.DB(r.db).C(r.collection)
	err := r.collection.Insert(b)

	if err != nil {
		return b, err
	}

	return b, nil
}

