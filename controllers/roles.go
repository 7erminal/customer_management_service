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

// RolesController operations for Roles
type RolesController struct {
	beego.Controller
}

// URLMapping ...
func (c *RolesController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("GetOneByName", c.GetOneByName)
}

// Post ...
// @Title Post
// @Description create Roles
// @Param	body		body 	requests.RolesRequest	true		"body for Roles content"
// @Success 200 {int} responses.RoleResponseDTO
// @Failure 403 body is empty
// @router / [post]
func (c *RolesController) Post() {
	var v requests.RolesRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	var role models.Roles = models.Roles{Role: v.Role, Description: v.Description, DateCreated: time.Now(), DateModified: time.Now(), Active: 1, CreatedBy: v.AddedBy, ModifiedBy: v.AddedBy}
	if _, err := models.AddRoles(&role); err == nil {
		c.Ctx.Output.SetStatus(200)
		var resp = responses.RoleResponseDTO{StatusCode: 200, Role: &role, StatusDesc: "Role added"}
		c.Data["json"] = resp
	} else {
		var resp = responses.RoleResponseDTO{StatusCode: 604, Role: nil, StatusDesc: "Error adding role ::: " + err.Error()}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Roles by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} responses.RoleResponseDTO
// @Failure 403 :id is empty
// @router /:id [get]
func (c *RolesController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	logs.Info("Role received is ", idStr)
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetRolesById(id)
	if err != nil {
		logs.Error("Unable to fetch role ", err.Error())
		var resp = responses.RoleResponseDTO{StatusCode: 604, Role: nil, StatusDesc: "Error getting user ::: " + err.Error()}
		c.Data["json"] = resp
	} else {
		logs.Info("Role fetched")
		var resp = responses.RoleResponseDTO{StatusCode: 200, Role: v, StatusDesc: "Role fetched"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// GetOneByName ...
// @Title Get One By Role
// @Description get Roles by name
// @Param	role		path 	string	true		"The key for staticblock"
// @Success 200 {object} responses.RoleResponseDTO
// @Failure 403 :name is empty
// @router /role/:role [get]
func (c *RolesController) GetOneByName() {
	role := c.Ctx.Input.Param(":role")
	v, err := models.GetRolesByName(role)
	if err != nil {
		var resp = responses.RoleResponseDTO{StatusCode: 604, Role: nil, StatusDesc: "Error getting user ::: " + err.Error()}
		c.Data["json"] = resp
	} else {
		var resp = responses.RoleResponseDTO{StatusCode: 200, Role: v, StatusDesc: "Role fetched"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Roles
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} responses.RolesAllResponseDTO
// @Failure 403
// @router / [get]
func (c *RolesController) GetAll() {
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

	l, err := models.GetAllRoles(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error("An error occurred while fetching roles: ", err)
		var resp = responses.RolesAllResponseDTO{StatusCode: 604, Roles: &[]models.Roles{}, StatusDesc: "Error getting roles ::: " + err.Error()}
		c.Data["json"] = resp
	} else {
		rolePermissions := []models.Roles{}
		for _, urs := range l {
			m := urs.(models.Roles)

			rolePermissions = append(rolePermissions, m)
		}
		var resp = responses.RolesAllResponseDTO{StatusCode: 200, Roles: &rolePermissions, StatusDesc: "Roles fetched"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Roles
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Roles	true		"body for Roles content"
// @Success 200 {object} models.Roles
// @Failure 403 :id is not int
// @router /:id [put]
func (c *RolesController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Roles{RoleId: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateRolesById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Roles
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *RolesController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	logs.Info("Deleting role with ID: ", id)
	statusCode := 400
	statusMessage := "Deleting role"

	roleId := idStr
	if role, err := models.GetRolesById(id); err != nil {
		logs.Error("Error getting role by ID: ", err)
		statusCode = 404
		statusMessage = "Role not found"
		if role, err = models.GetRolesByName(idStr); err != nil {
			logs.Error("Error getting role by name: ", err)
			statusCode = 404
			statusMessage = "Role not found by name"
		} else {
			roleId = strconv.FormatInt(role.RoleId, 10)
		}
	} else {
		logs.Info("Role found: ", role)
		roleId = strconv.FormatInt(role.RoleId, 10)
	}

	roleIdInt, _ := strconv.ParseInt(roleId, 0, 64)

	query := map[string]string{"RoleId": roleId}
	fields := []string{}
	sortby := []string{}
	order := []string{}
	offset := int64(0)
	limit := int64(100)
	l, err := models.GetAllRole_permissions(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error("Error getting role permissions: ", err)
	} else {
		logs.Info("Role permissions found: ", l)

		for _, urs := range l {
			m := urs.(models.Role_permissions)

			if err := models.DeleteRole_permissions(m.RolePermissionId); err == nil {
				// c.Data["json"] = "OK"
				logs.Info("Role permission deleted successfully")
			} else {
				logs.Error("Error deleting role permission: %v", err)
			}
		}
	}
	if err := models.DeleteRoles(roleIdInt); err == nil {
		statusCode = 200
		statusMessage = "Role deleted successfully"
		resp := responses.RoleResponseDTO{StatusCode: statusCode, StatusDesc: statusMessage}
		c.Data["json"] = resp
	} else {
		logs.Error("Error deleting role: ", err)
		statusCode = 400
		statusMessage = "Error deleting role"
		resp := responses.RoleResponseDTO{StatusCode: statusCode, StatusDesc: statusMessage + " ::: " + err.Error()}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}
