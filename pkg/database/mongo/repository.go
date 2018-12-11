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
	session := db.Session.Copy()
	c := session.DB(db.Name).C(collection)
	return &MongoRepository {
		DB: db,
		collection: c,
	}
}

func (r *MongoRepository) Find(model interface{}) (interface{}) {
	res := r.collection.Find(&model)
	return res
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


	// structValue := reflect.ValueOf(model).Elem()
	// structFieldValue := structValue.FieldByName("ID")

	// if !structFieldValue.IsValid() {
	// 	return fmt.Errorf("No such field: %s in obj", name)
	// }

	// if !structFieldValue.CanSet() {
	// 	return fmt.Errorf("Cannot set %s field value", name)
	// }

	// structFieldType := structFieldValue.Type()
	// val := reflect.ValueOf(value)
	// if structFieldType != val.Type() {
	// 	invalidTypeError := errors.New("Provided value type didn't match obj field type")
	// 	return invalidTypeError
	// }

	// val := reflect.ValueOf(bson.NewObjectId())
	// // structFieldValue.Set(val)
	// var i reflect.Type
	// log.Info("Type: ", reflect.TypeOf(model))
	// i = reflect.TypeOf(model)
	// log.Info("i: ", i)
	// t := model.(i)


	err := r.collection.Insert(&model)

	// if err != nil {
	// 	return model, err
	// }

	return model, err
}

func (r *MongoRepository) RemoveById(id string) (error) {
	_id := bson.ObjectIdHex(id)

	err := r.collection.Remove(bson.M{"_id": _id})

	if err != nil {
		return err
	}

	return nil
}

