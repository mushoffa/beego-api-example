package controllers

import (
	// "reflect"
	"strconv"

	"beego-api-example/models"
	
	"beego-api-example/pkg/database/config"
	mongodb "beego-api-example/pkg/database/mongo"
	service "beego-api-example/pkg/database/service"
	

	"github.com/astaxie/beego"
	"github.com/json-iterator/go"
)

//  FarmController operations for Farm
type FarmController struct {
	beego.Controller
}

// URLMapping ...
func (c *FarmController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Nearby", c.Nearby)
}

func NewFarmService() *service.FarmService {
	return service.NewFarmService(
			mongodb.NewMongoRepository(mongodb.MongoSession, config.FARM),
			mongodb.MongoSession.Session,
		)
}

// Post ...
// @Title Post
// @Description create Farm
// @Security Bearer
// @Param	body		body 	models.Farm	false		"body for Farm content"
// @Success 201 {int} models.Farm
// @Failure 403 body is empty
// @router / [post]
func (c *FarmController) Post() {
	var farm models.Farm
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	
	json.Unmarshal(c.Ctx.Input.RequestBody, &farm)

	r := NewFarmService()
	defer r.CloseSession()

	_, err := r.Insert(farm)

	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = farm
	}

	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Farm by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Farm
// @Failure 403 :id is empty
// @router /:id [get]
func (c *FarmController) GetOne() {
	c.Data["json"] = "Work In Progress"
	c.ServeJSON()
}

// GetAll ...
// @Title Get list of farms from database
// @Description get Farm
// Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Farm
// @Failure 403
// @router / [get]
func (c *FarmController) GetAll() {
	repo := NewFarmService()
	defer repo.CloseSession()
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
// @Description update the Farm
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Farm	true		"body for Farm content"
// @Success 200 {object} models.Farm
// @Failure 403 :id is not int
// @router /:id [put]
func (c *FarmController) Put() {
	c.Data["json"] = "Work In Progress"
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Farm
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *FarmController) Delete() {
	c.Data["json"] = "Work In Progress"
	c.ServeJSON()
}


// @Title Nearby
// @Description Finds nearby farm location from current user location
// @Param	latitude	query	float64	true	"latitude"
// @Param	longitude	query	float64 true	"longitude"
// @Param	radius		query	number	true	"Search radius (in meter)"
// @Success 200 {object} models.Farm
// @Failure 403 no nearby farms
// @router /nearby [get]
func (c *FarmController) Nearby() {

	latitude, _ := strconv.ParseFloat(c.GetString("latitude"), 64)
	longitude, _ := strconv.ParseFloat(c.GetString("longitude"), 64)
	radius, _ := strconv.ParseFloat(c.GetString("radius"), 64)

	repo := NewFarmService()
	defer repo.CloseSession()

	res := repo.Nearby(latitude, longitude, radius)

	c.Data["json"] = res
	c.Data["json"] = "Work In Progress"
	c.ServeJSON()
}
