package models

import (
	// "errors"
	// "fmt"
	// "reflect"
	// "strings"

	// "github.com/astaxie/beego/orm"
	"gopkg.in/mgo.v2/bson"
)

type Farm struct {
	ID   		bson.ObjectId 	`json:"id,omitempty" bson:"_id,omitempty" mapstructure:"id"`
	Name 		string 			`json:"name" bson:"name" mapstructure:"name"`
	Location 	Location 		`json:"location" bson:"location" mapstructure:"location"`
	// Location 	struct {
	// 	Latitude	float64		`json:"latitude"		bson:"latitude"			mapstructure:"latitude"`
	// 	Longitude	float64		`json:"longitude"		bson:"longitude"		mapstructure:"longitude"`
	// } 							`json:"location"		bson:"location"			mapstructure:"location"`
}

type Location struct {
	Latitude	float64			`json:"latitude" bson:"latitude" mapstructure:"latitude"`
	Longitude	float64			`json:"longitude" bson:"longitude" mapstructure:"longitude"`
}

type Farms struct {
	Seeds	[]Farm
}

// func init() {
// 	orm.RegisterModel(new(Farm))
// }

// AddFarm insert a new Farm into database and returns
// last inserted Id on success.
// func AddFarm(m *Farm) (id int64, err error) {
// 	o := orm.NewOrm()
// 	id, err = o.Insert(m)
// 	return
// }

// // GetFarmById retrieves Farm by Id. Returns error if
// // Id doesn't exist
// func GetFarmById(id int64) (v *Farm, err error) {
// 	o := orm.NewOrm()
// 	v = &Farm{Id: id}
// 	if err = o.QueryTable(new(Farm)).Filter("Id", id).RelatedSel().One(v); err == nil {
// 		return v, nil
// 	}
// 	return nil, err
// }

// // GetAllFarm retrieves all Farm matches certain condition. Returns empty list if
// // no records exist
// func GetAllFarm(query map[string]string, fields []string, sortby []string, order []string,
// 	offset int64, limit int64) (ml []interface{}, err error) {
// 	o := orm.NewOrm()
// 	qs := o.QueryTable(new(Farm))
// 	// query k=v
// 	for k, v := range query {
// 		// rewrite dot-notation to Object__Attribute
// 		k = strings.Replace(k, ".", "__", -1)
// 		qs = qs.Filter(k, v)
// 	}
// 	// order by:
// 	var sortFields []string
// 	if len(sortby) != 0 {
// 		if len(sortby) == len(order) {
// 			// 1) for each sort field, there is an associated order
// 			for i, v := range sortby {
// 				orderby := ""
// 				if order[i] == "desc" {
// 					orderby = "-" + v
// 				} else if order[i] == "asc" {
// 					orderby = v
// 				} else {
// 					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
// 				}
// 				sortFields = append(sortFields, orderby)
// 			}
// 			qs = qs.OrderBy(sortFields...)
// 		} else if len(sortby) != len(order) && len(order) == 1 {
// 			// 2) there is exactly one order, all the sorted fields will be sorted by this order
// 			for _, v := range sortby {
// 				orderby := ""
// 				if order[0] == "desc" {
// 					orderby = "-" + v
// 				} else if order[0] == "asc" {
// 					orderby = v
// 				} else {
// 					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
// 				}
// 				sortFields = append(sortFields, orderby)
// 			}
// 		} else if len(sortby) != len(order) && len(order) != 1 {
// 			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
// 		}
// 	} else {
// 		if len(order) != 0 {
// 			return nil, errors.New("Error: unused 'order' fields")
// 		}
// 	}

// 	var l []Farm
// 	qs = qs.OrderBy(sortFields...).RelatedSel()
// 	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
// 		if len(fields) == 0 {
// 			for _, v := range l {
// 				ml = append(ml, v)
// 			}
// 		} else {
// 			// trim unused fields
// 			for _, v := range l {
// 				m := make(map[string]interface{})
// 				val := reflect.ValueOf(v)
// 				for _, fname := range fields {
// 					m[fname] = val.FieldByName(fname).Interface()
// 				}
// 				ml = append(ml, m)
// 			}
// 		}
// 		return ml, nil
// 	}
// 	return nil, err
// }

// // UpdateFarm updates Farm by Id and returns error if
// // the record to be updated doesn't exist
// func UpdateFarmById(m *Farm) (err error) {
// 	o := orm.NewOrm()
// 	v := Farm{Id: m.Id}
// 	// ascertain id exists in the database
// 	if err = o.Read(&v); err == nil {
// 		var num int64
// 		if num, err = o.Update(m); err == nil {
// 			fmt.Println("Number of records updated in database:", num)
// 		}
// 	}
// 	return
// }

// // DeleteFarm deletes Farm by Id and returns error if
// // the record to be deleted doesn't exist
// func DeleteFarm(id int64) (err error) {
// 	o := orm.NewOrm()
// 	v := Farm{Id: id}
// 	// ascertain id exists in the database
// 	if err = o.Read(&v); err == nil {
// 		var num int64
// 		if num, err = o.Delete(&Farm{Id: id}); err == nil {
// 			fmt.Println("Number of records deleted in database:", num)
// 		}
// 	}
// 	return
// }
