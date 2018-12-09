package repository

import (
	"beego-api-example/models"
	"beego-api-example/pkg/database"
	mongo "beego-api-example/pkg/database/mongo"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type SeedRepository struct {
	model		database.Model
	DB 			*mongo.Database
	collection	*mgo.Collection
}

func NewSeedRepository(db *mongo.Database, collection string) *SeedRepository {
	c := db.Session.DB(db.Name).C(collection)
	return &SeedRepository {
		DB: db,
		collection: c,
	}
}

func (r *SeedRepository) FindAll() ([]*models.Seed, error) {
	var res []*models.Seed

	err := r.collection.Find(nil).Sort("name").All(&res)

	if err != nil {
		return res, err
	}

	return res, err
}

func (r *SeedRepository) Insert(seed *models.Seed) (*models.Seed, error) {

	seed.Id = bson.NewObjectId()
	err := r.collection.Insert(&seed)

	if err != nil {
		return seed, err
	}

	return seed, nil
}

func (r *SeedRepository) RemoveById(id string) (error) {
	_id := bson.ObjectIdHex(id)

	err := r.collection.Remove(bson.M{"_id": _id})

	if err != nil {
		return err
	}

	return nil
}