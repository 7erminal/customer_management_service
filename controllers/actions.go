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

// ActionsController operations for Actions
type ActionsController struct {
	beego.Controller
}

// URLMapping ...
func (c *ActionsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Actions
// @Param	body		body 	models.Actions	true		"body for Actions content"
// @Success 201 {int} requests.ActionRequest
// @Failure 403 body is empty
// @router / [post]
func (c *ActionsController) Post() {

	var v requests.ActionRequest

	statusCode := 400
	statusMessage := "Failed to create action"

	acModel := models.Actions{
		Action:       v.Action,
		Description:  v.Description,
		DateCreated:  time.Now(),
		DateModified: time.Now(),
		CreatedBy:    1,
		ModifiedBy:   1,
		Active:       1,
	}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if _, err := models.AddActions(&acModel); err == nil {
		c.Ctx.Output.SetStatus(200)
		statusCode = 200
		statusMessage = "Action created successfully"

	} else {
		c.Data["json"] = err.Error()
	}
	resp := responses.ActionResponseDTO{
		StatusCode: statusCode,
		StatusDesc: statusMessage,
		Action:     &acModel,
	}
	c.Data["json"] = resp
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Actions by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Actions
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ActionsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetActionsById(id)

	statusCode := 400
	statusMessage := "Failed to get action"
	data := models.Actions{ActionId: id}
	if err != nil {
		logs.Error(err)
		statusCode = 400
		statusMessage = "Failed to get action"
	} else {
		data = *v
		statusCode = 200
		statusMessage = "Action retrieved successfully"
	}
	c.Data["json"] = responses.ActionResponseDTO{
		StatusCode: statusCode,
		StatusDesc: statusMessage,
		Action:     &data,
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Actions
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Actions
// @Failure 403
// @router / [get]
func (c *ActionsController) GetAll() {
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

	statusCode := 400
	statusMessage := "Failed to get actions"
	data := []models.Actions{}

	l, err := models.GetAllActions(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error("An error occurred while fetching actions ", err)
		statusCode = 500
		statusMessage = "An error occurred while fetching actions"
	} else {
		for _, action := range l {
			m := action.(models.Actions)
			data = append(data, m)
		}
		statusCode = 200
		statusMessage = "Actions retrieved successfully"
		c.Data["json"] = data
	}
	resp := responses.ActionsResponseDTO{
		StatusCode: statusCode,
		StatusDesc: statusMessage,
		Actions:    &data,
	}
	c.Data["json"] = resp
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Actions
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Actions	true		"body for Actions content"
// @Success 200 {object} models.Actions
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ActionsController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Actions{ActionId: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateActionsById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Actions
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ActionsController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteActions(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
