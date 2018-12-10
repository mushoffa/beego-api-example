package controllers

import (
	_ "beego-api-example/models"

	"github.com/astaxie/beego"
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

// Post ...
// @Title Post
// @Description create Farm
// @Param	body		body 	models.Farm	true		"body for Farm content"
// @Success 201 {int} models.Farm
// @Failure 403 body is empty
// @router / [post]
func (c *FarmController) Post() {

}

// GetOne ...
// @Title Get One
// @Description get Farm by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Farm
// @Failure 403 :id is empty
// @router /:id [get]
func (c *FarmController) GetOne() {

}

// GetAll ...
// @Title Get All
// @Description get Farm
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Farm
// @Failure 403
// @router / [get]
func (c *FarmController) GetAll() {
	
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

}

// Delete ...
// @Title Delete
// @Description delete the Farm
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *FarmController) Delete() {

}


// @Title Nearby
// @Description Finds nearby farm location from current user location
// @Param	latitude	path	float64	true	"latitude"
// @Param	longitude	path	float64 true	"longitude"
// @Param	radius		path	number	true	"Search radius (in meter)"
// @Success 200 {object} models.Farm
// @Failure 403 no nearby farms
// @router /nearby [post]
func (c *FarmController) Nearby() {
	c.Data["json"] = "OK"
	c.ServeJSON()
}
