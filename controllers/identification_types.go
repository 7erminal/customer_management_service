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

// Identification_typesController operations for Identification_types
type Identification_typesController struct {
	beego.Controller
}

// URLMapping ...
func (c *Identification_typesController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Identification_types
// @Param	body		body 	models.Identification_types	true		"body for Identification_types content"
// @Success 201 {int} models.Identification_types
// @Failure 403 body is empty
// @router / [post]
func (c *Identification_typesController) Post() {
	var v models.Identification_types
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if _, err := models.AddIdentification_types(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = v
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Identification_types by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Identification_types
// @Failure 403 :id is empty
// @router /:id [get]
func (c *Identification_typesController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetIdentification_typesById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Identification_types
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Identification_types
// @Failure 403
// @router / [get]
func (c *Identification_typesController) GetAll() {
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

	l, err := models.GetAllIdentification_types(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error("Error fetching ID Types ", err.Error())
		resp := responses.IDTypesResponseDTO{StatusCode: 301, IdTypes: nil, StatusDesc: "ID Type not fetched"}
		c.Data["json"] = resp
	} else {
		// idResp := []responses.IDTypeResponse{}
		// for _, urs := range l {
		// 	m := urs.(models.Identification_types)

		// 	idResp = append(idResp, m)
		// }
		if l == nil {
			l = []interface{}{}
		}

		resp := responses.IDTypesResponseDTO{StatusCode: 200, IdTypes: &l, StatusDesc: "Users fetched successfully"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Identification_types
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Identification_types	true		"body for Identification_types content"
// @Success 200 {object} models.Identification_types
// @Failure 403 :id is not int
// @router /:id [put]
func (c *Identification_typesController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Identification_types{IdentificationTypeId: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateIdentification_typesById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Identification_types
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *Identification_typesController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteIdentification_types(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
