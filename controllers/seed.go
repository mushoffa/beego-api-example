package controllers

import (
	"encoding/json"
	_ "reflect"
	"github.com/astaxie/beego"
	mgo "gopkg.in/mgo.v2"

	mongodb "beego-api-example/pkg/database/mongo"
	"beego-api-example/models"
	repository "beego-api-example/pkg/database/repository"
)

// SeedController operations for Seed
type SeedController struct {
	beego.Controller
}

// URLMapping ...
func (c *SeedController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

func SeedRepository() *repository.SeedRepository {
	mongodb_host := beego.AppConfig.String("mongodb.url")
	//session, _ := mongodb.NewMongoDbSession(mongodb_host)
	session, _ := mgo.Dial(mongodb_host)

	db := &mongodb.Database{
		Name: "farmingo",
		Session: session,
	}

	return repository.NewSeedRepository(
		db,
		"seed",
	)
}

// func NewSeedDbService() *database.DatabaseService {
// 	return database.NewDatabaseService(
// 		mongodb.NewMongoRepository(
// 			mongodb.MongoSession.Copy(),
// 			"farmingo",
// 			"seed",
// 		),
// 	)
// }

// Post ...
// @Title Create
// @Description create Seed
// body	body	models.Seed 	true	""
// @Param	body	body	models.Seed	true	"name of seed"
// @Success 201 {object} models.Seed
// @Failure 403 body is empty
// @router / [post]
func (c *SeedController) Post() {
	var seed models.Seed
	// var err error
	
	json.Unmarshal(c.Ctx.Input.RequestBody, &seed)
	repo := SeedRepository()
	defer repo.DB.Session.Close()

	_, err := repo.Insert(&seed)

	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = seed
	}

	// if err != nil {
	// 	c.Data["json"] = err.Error()
	// } else {
	// 	seed.Id = bson.NewObjectId()
	// 	session := database.MgoSession.Copy()
	// 	defer session.Close()
	// 	coll := session.DB("farmingo").C("seed")
	// 	err = coll.Insert(&seed)

	// 	if err != nil {
	// 		c.Data["json"]  = err.Error()
	// 	}
	// }

	c.ServeJSON()
}

// GetOne ...
// @Title GetOne
// @Description get Seed by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Seed
// @Failure 403 :id is empty
// @router /:id [get]
func (c *SeedController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Seed
// 	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// 	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// 	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// 	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// 	limit	query	string	false	"Limit the size of result set. Must be an integer"
// 	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Seed
// @Failure 403
// @router / [get]
func (c *SeedController) GetAll() {

	//var res []*models.Seed
	// var err error
	// session := database.MgoSession.Copy()
	// defer session.Close()
	// coll := session.DB("farmingo").C("seed")
	// coll.Find(nil).Sort("name").All(&seeds)
	// s := NewSeedDbService()
	// seeds := reflect.TypeOf(([]models.Seed)(nil))
	// seeds := []models.Seed{}
	// data, err := s.FindAll()
	repo := SeedRepository()
	defer repo.DB.Session.Close()
	res, err := repo.FindAll()

	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = res	
	}


	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Seed
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Seed	true		"body for Seed content"
// @Success 200 {object} models.Seed
// @Failure 403 :id is not int
// @router /:id [put]
func (c *SeedController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Seed
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *SeedController) Delete() {
	seedId := c.Ctx.Input.Param(":id")
	repo := SeedRepository()
	defer repo.DB.Session.Close()

	err := repo.RemoveById(seedId)

	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = "OK"
	}

	c.ServeJSON()
}
