package controllers

import (
	"customer_management_service/models"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"

	beego "github.com/beego/beego/v2/server/web"
)

// Customer_categoriesController operations for Customer_categories
type Customer_categoriesController struct {
	beego.Controller
}

// URLMapping ...
func (c *Customer_categoriesController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Customer_categories
// @Param	body		body 	models.CustomerCategoriesRequestDTO	true		"body for Customer_categories content"
// @Success 200 {int} models.CustomerCategoriesResponseDTO
// @Failure 403 body is empty
// @router / [post]
func (c *Customer_categoriesController) Post() {
	var v models.CustomerCategoriesRequestDTO
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	var custCat = models.Customer_categories{Category: v.Category, Description: v.Description, CreatedBy: v.CreatedBy, ModifiedBy: v.CreatedBy, Active: 1, DateCreated: time.Now(), DateModified: time.Now()}
	if _, err := models.AddCustomer_categories(&custCat); err == nil {
		c.Ctx.Output.SetStatus(200)
		var resp = models.CustomerCategoriesResponseDTO{StatusCode: 200, Category: &custCat, StatusDesc: "Customer Category Added Successfully"}
		c.Data["json"] = resp
	} else {
		var resp = models.CustomerCategoriesResponseDTO{StatusCode: 301, Category: nil, StatusDesc: err.Error()}

		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Customer_categories by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Customer_categories
// @Failure 403 :id is empty
// @router /:id [get]
func (c *Customer_categoriesController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetCustomer_categoriesById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Customer_categories
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Customer_categories
// @Failure 403
// @router / [get]
func (c *Customer_categoriesController) GetAll() {
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

	l, err := models.GetAllCustomer_categories(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Customer_categories
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Customer_categories	true		"body for Customer_categories content"
// @Success 200 {object} models.Customer_categories
// @Failure 403 :id is not int
// @router /:id [put]
func (c *Customer_categoriesController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Customer_categories{CustomerCategoryId: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateCustomer_categoriesById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Customer_categories
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *Customer_categoriesController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteCustomer_categories(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
