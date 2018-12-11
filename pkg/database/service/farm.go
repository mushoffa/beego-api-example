package service

import (
	// "fmt"
	"beego-api-example/models"
	"beego-api-example/pkg/database"
	// "beego-api-example/pkg/database/mongo"
	log "github.com/sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type FarmService struct {
	r	database.RepositoryHandler
	// session   mongo.Session
	session *mgo.Session
}

func NewFarmService(repository database.RepositoryHandler, s *mgo.Session) *FarmService {
	return &FarmService {
		r: repository,
		session:s,
	}
}

func (s *FarmService) FindAll() (interface{}, error) {
	return s.r.FindAll()
}

func (s *FarmService) Insert(farm interface{}) (interface{}, error) {
	
	// Type assertion example for overriding abstract struct field
	// var i interface{}

	// i = farm
	// t := i.(models.Farm)
	// fmt.Println(t)
	// t.ID = bson.NewObjectId()

	// i.ID = bson.NewObjectId()

	return s.r.Insert(farm)
}

func (s *FarmService) Nearby(lat float64, lng float64, radius float64) ([]models.Farm) {

	// var res models.Farms
	res := []models.Farm{}

	// session := s.session.Copy()

	c := s.session.DB("farmingo").C("farm")

	err := c.Find(bson.M{
						"location": bson.M{
							// "$nearSphere": []float64{lat,lng},
							"$near": bson.M{
								// "lat": lat,
								// "lng": lng,
								"$geometry": bson.M{
									"type": "Point",
									"coordinates": []float64{lng,lat},
								},
							"$maxDistance": radius,
							},
						},
					}).All(&res)
	// .Sort("name").All(&res)

	log.Info("Nearby-Error: ", err)
	log.Info("Res: ", res)
	return res
}

func (s *FarmService) CloseSession() {
	// s.Session.Close()
	s.session.Close()
}