package controllers

import (
	"customer_management_service/models"
	"customer_management_service/structs/responses"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// CustomersController operations for Customers
type UserExtraDetailsController struct {
	beego.Controller
}

// URLMapping ...
func (c *UserExtraDetailsController) URLMapping() {
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Customers
// @Param	body		body 	models.Customers	true		"body for Customers content"
// @Success 201 {int} models.Customers
// @Failure 403 body is empty
// @router / [post]
// func (c *CustomersController) Post() {
// 	var v models.Customers
// 	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
// 	if _, err := models.AddCustomers(&v); err == nil {
// 		c.Ctx.Output.SetStatus(201)
// 		c.Data["json"] = v
// 	} else {
// 		c.Data["json"] = err.Error()
// 	}
// 	c.ServeJSON()
// }

// GetOne ...
// @Title Get One
// @Description get Customers by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Customers
// @Failure 403 :id is empty
// @router /:id [get]
func (c *UserExtraDetailsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetUserExtraDetailsById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Customers
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	search	query	string	false	"Filter. e.g. example@gmail.com ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Customers
// @Failure 403
// @router / [get]
func (c *UserExtraDetailsController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var search = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	// search: k:v,k:v
	if v := c.GetString("search"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid search key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			search[k] = v
		}
	}

	l, err := models.GetAllCustomers(query, fields, sortby, order, offset, limit, search)
	if err != nil {
		resp := responses.StringResponseDTO{StatusCode: 301, Value: err.Error(), StatusDesc: "Error fetching customers"}
		c.Data["json"] = resp
	} else {
		resp := responses.CustomersDTO{StatusCode: 200, Customers: &l, StatusDesc: "Successfully fetched categories"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Customers
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Customers	true		"body for Customers content"
// @Success 200 {object} models.CustomerResponseDTO
// @Failure 403 :id is not int
// @router /:id [put]
func (c *UserExtraDetailsController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.UserExtraDetails{UserDetailsId: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateUserExtraDetailsById(&v); err == nil {
		c.Ctx.Output.SetStatus(200)
		var resp = responses.UserExtraDetailsResponseDTO{StatusCode: 200, UserDetails: &v, StatusDesc: "Customer updated successfully"}
		c.Data["json"] = resp
	} else {
		logs.Error("Customer update failed ", err.Error())
		var resp = responses.UserExtraDetailsResponseDTO{StatusCode: 608, UserDetails: &v, StatusDesc: "Customer update failed"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Customers
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *UserExtraDetailsController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteUserExtraDetails(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
