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

	beego "github.com/beego/beego/v2/server/web"
)

// PermissionsController operations for Permissions
type PermissionsController struct {
	beego.Controller
}

// URLMapping ...
func (c *PermissionsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Permissions
// @Param	body		body 	requests.PermissionRequest	true		"body for Permissions content"
// @Success 200 {int} responses.PermissionResponseDTO
// @Failure 403 body is empty
// @router / [post]
func (c *PermissionsController) Post() {
	var v requests.PermissionRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if action, err := models.GetActionsByName(v.Action); err == nil {
		permissionCode := strings.ToUpper(v.Permission) + "_" + strings.ToUpper(v.Action)
		permission := models.Permissions{Permission: v.Permission, PermissionDescription: v.Description, PermissionCode: permissionCode, Action: action, DateCreated: time.Now(), DateModified: time.Now(), Active: 1, CreatedBy: 1, ModifiedBy: 1}
		if _, err := models.AddPermissions(&permission); err == nil {
			c.Ctx.Output.SetStatus(200)
			var resp = responses.PermissionResponseDTO{StatusCode: 200, Permission: &permission, StatusDesc: "Permission added"}
			c.Data["json"] = resp
		} else {
			var resp = responses.PermissionResponseDTO{StatusCode: 604, Permission: nil, StatusDesc: "Error adding permission ::: " + err.Error()}
			c.Data["json"] = resp
		}
	} else {
		var resp = responses.PermissionResponseDTO{StatusCode: 604, Permission: nil, StatusDesc: "Error adding permission ::: " + err.Error()}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Permissions by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} responses.PermissionResponseDTO
// @Failure 403 :id is empty
// @router /:id [get]
func (c *PermissionsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetPermissionsById(id)
	if err != nil {
		var resp = responses.PermissionResponseDTO{StatusCode: 604, Permission: nil, StatusDesc: "Error getting permission ::: " + err.Error()}
		c.Data["json"] = resp
	} else {
		var resp = responses.PermissionResponseDTO{StatusCode: 200, Permission: v, StatusDesc: "Permission fetched"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Permissions
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} responses.PermissionsAllResponseDTO
// @Failure 403
// @router / [get]
func (c *PermissionsController) GetAll() {
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

	l, err := models.GetAllPermissions(query, fields, sortby, order, offset, limit)
	if err != nil {
		var resp = responses.PermissionsAllResponseDTO{StatusCode: 604, Permissions: nil, StatusDesc: "Error getting permissions ::: " + err.Error()}
		c.Data["json"] = resp
	} else {
		var resp = responses.PermissionsAllResponseDTO{StatusCode: 200, Permissions: &l, StatusDesc: "Permissions fetched"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Permissions
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Permissions	true		"body for Permissions content"
// @Success 200 {object} models.Permissions
// @Failure 403 :id is not int
// @router /:id [put]
func (c *PermissionsController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Permissions{PermissionId: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdatePermissionsById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Permissions
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *PermissionsController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeletePermissions(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
