package controllers

import (
	"customer_management_service/models"
	"customer_management_service/structs/requests"
	"customer_management_service/structs/responses"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/beego/beego/v2/core/logs"

	beego "github.com/beego/beego/v2/server/web"
)

// Newsletter_customersController operations for Newsletter_customers
type Newsletter_customersController struct {
	beego.Controller
}

// URLMapping ...
func (c *Newsletter_customersController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	// c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Newsletter_customers
// @Param	body		body 	requests.NewsletterRequestDTO	true		"body for Newsletter_customers content"
// @Success 201 {int} models.Newsletter_customers
// @Failure 403 body is empty
// @router / [post]
func (c *Newsletter_customersController) Post() {
	var v requests.NewsletterRequestDTO
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	var nl models.Newsletter_customers = models.Newsletter_customers{FirstName: v.FirstName, LastName: v.LastName, Email: v.Email, DateCreated: time.Now(), DateModified: time.Now(), Active: 1, CreatedBy: 1, ModifiedBy: 1}
	if _, err := models.AddNewsletter_customers(&nl); err == nil {
		resp := responses.StringResponseDTO{StatusCode: 200, Value: "SUCCESSFUL", StatusDesc: "Successfully added customer to newsletter"}
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = resp
	} else {
		resp := responses.StringResponseDTO{StatusCode: 200, Value: "ERROR", StatusDesc: "An error occurred when adding customer to newsletter"}
		c.Data["json"] = err.Error()
		logs.Debug("An error occurred when adding customer to newsletter ", err.Error())
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Newsletter_customers by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Newsletter_customers
// @Failure 403 :id is empty
// @router /:id [get]
func (c *Newsletter_customersController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetNewsletter_customersById(id)
	if err != nil {
		var resp responses.NewsLetterCustomerDTO = responses.NewsLetterCustomerDTO{StatusCode: 301, Customer: nil, StatusDesc: "Error fetching customer " + err.Error()}
		c.Data["json"] = resp
	} else {
		var resp responses.NewsLetterCustomerDTO = responses.NewsLetterCustomerDTO{StatusCode: 200, Customer: v, StatusDesc: "Successfully fetched customer"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Newsletter_customers
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Newsletter_customers
// @Failure 403
// @router / [get]
func (c *Newsletter_customersController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
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

	l, err := models.GetAllNewsletter_customers(query, fields, sortby, order, offset, limit)
	if err != nil {
		var resp responses.NewsLetterAllCustomersDTO = responses.NewsLetterAllCustomersDTO{StatusCode: 301, Customers: nil, StatusDesc: "Error fetching customers " + err.Error()}
		c.Data["json"] = resp
	} else {
		var resp responses.NewsLetterAllCustomersDTO = responses.NewsLetterAllCustomersDTO{StatusCode: 200, Customers: &l, StatusDesc: "Successfully fetched customers "}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Newsletter_customers
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Newsletter_customers	true		"body for Newsletter_customers content"
// @Success 200 {object} models.Newsletter_customers
// @Failure 403 :id is not int
// @router /:id [put]
func (c *Newsletter_customersController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Newsletter_customers{CustomerId: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateNewsletter_customersById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Newsletter_customers
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *Newsletter_customersController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteNewsletter_customers(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
