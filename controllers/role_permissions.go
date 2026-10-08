package controllers

import (
	"customer_management_service/models"
	"customer_management_service/structs/requests"
	"customer_management_service/structs/responses"
	"encoding/json"
	"errors"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// Role_permissionsController operations for Role_permissions
type Role_permissionsController struct {
	beego.Controller
}

// URLMapping ...
func (c *Role_permissionsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Role_permissions
// @Param	body		body 	requests.RolePermissionRequest	true		"body for Role_permissions content"
// @Success 200 {int} responses.RolePermissionResponseDTO
// @Failure 403 body is empty
// @router / [post]
func (c *Role_permissionsController) Post() {
	var v requests.RolePermissionRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	roleInt, _ := strconv.ParseInt(v.Role, 0, 64)
	if role, err := models.GetRolesById(roleInt); err == nil {
		if permission, err := models.GetPermissionsByCode(v.PermissionCode); err == nil {
			if action, err := models.GetActionsByName(v.Action); err == nil {
				if rolePermission, err := models.GetRolePermissionByRoleActionPermission(roleInt, v.Action, v.PermissionCode); err == nil {
					if rolePermission != nil {
						var resp = responses.RoleResponseDTO{StatusCode: 604, Role: nil, StatusDesc: "Role permission already exists"}
						c.Data["json"] = resp
						c.ServeJSON()
						return
					}
				}
				var rolePermission models.Role_permissions = models.Role_permissions{Role: role, Permission: permission, Action: action, DateCreated: time.Now(), DateModified: time.Now(), Active: 1, CreatedBy: v.AddedBy, ModifiedBy: v.AddedBy}
				if _, err := models.AddRole_permissions(&rolePermission); err == nil {
					logs.Info("Role permission added successfully")
					if role, err = models.GetRolesById(roleInt); err != nil {
						logs.Error("Unable to fetch role using id ", roleInt)
						var resp = responses.RoleResponseDTO{StatusCode: 604, Role: nil, StatusDesc: "Error adding role permission ::: " + err.Error()}
						c.Data["json"] = resp
						c.ServeJSON()
						return
					}
					c.Ctx.Output.SetStatus(200)
					respRole := responses.Roles{
						RoleId:       role.RoleId,
						Role:         role.Role,
						Description:  role.Description,
						DateCreated:  role.DateCreated,
						DateModified: role.DateModified,
						Active:       role.Active,
						CreatedBy:    role.CreatedBy,
						ModifiedBy:   role.ModifiedBy,
					}
					var resp = responses.RoleResponseDTO{StatusCode: 200, Role: &respRole, StatusDesc: "Role Permission added"}
					c.Data["json"] = resp
				} else {
					var resp = responses.RoleResponseDTO{StatusCode: 604, Role: nil, StatusDesc: "Error adding permission ::: " + err.Error()}
					c.Data["json"] = resp
				}
			} else {
				logs.Error("Unable to fetch action using name ", v.Action)
				var resp = responses.RoleResponseDTO{StatusCode: 604, Role: nil, StatusDesc: "Error adding role permission ::: " + err.Error()}
				c.Data["json"] = resp
			}
		} else {
			logs.Error("Unable to fetch permission using code ", v.PermissionCode)
			var resp = responses.RoleResponseDTO{StatusCode: 604, Role: nil, StatusDesc: "Error adding role permission ::: " + err.Error()}
			c.Data["json"] = resp
		}
	} else {
		logs.Error("Unable to fetch role using id ", v.Role)
		var resp = responses.RoleResponseDTO{StatusCode: 604, Role: nil, StatusDesc: "Error adding role permission ::: " + err.Error()}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Role_permissions by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Role_permissions
// @Failure 403 :id is empty
// @router /:id [get]
func (c *Role_permissionsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetRole_permissionsById(id)
	if err != nil {
		var resp = responses.RolePermissionResponseDTO{StatusCode: 604, RolePermission: nil, StatusDesc: "Error adding permission ::: " + err.Error()}
		c.Data["json"] = resp
	} else {
		var rolePermResp *responses.Role_permissions
		if v.Role != nil && v.Permission != nil && v.Action != nil {
			roleResp := responses.Roles{
				RoleId:       v.Role.RoleId,
				Role:         v.Role.Role,
				Description:  v.Role.Description,
				DateCreated:  v.Role.DateCreated,
				DateModified: v.Role.DateModified,
				CreatedBy:    v.Role.CreatedBy,
				ModifiedBy:   v.Role.ModifiedBy,
				Active:       v.Role.Active,
			}
			permResp := responses.Permissions{
				PermissionId:          strconv.FormatInt(v.Permission.PermissionId, 10),
				Permission:            v.Permission.Permission,
				PermissionCode:        v.Permission.PermissionCode,
				PermissionDescription: v.Permission.PermissionDescription,
				DateCreated:           v.Permission.DateCreated,
				DateModified:          v.Permission.DateModified,
				CreatedBy:             v.Permission.CreatedBy,
				ModifiedBy:            v.Permission.ModifiedBy,
				Active:                v.Permission.Active,
			}
			actionResp := responses.Actions{
				ActionId:     strconv.FormatInt(v.Action.ActionId, 10),
				Action:       v.Action.Action,
				Description:  v.Action.Description,
				DateCreated:  v.Action.DateCreated,
				DateModified: v.Action.DateModified,
				CreatedBy:    v.Action.CreatedBy,
				ModifiedBy:   v.Action.ModifiedBy,
				Active:       v.Action.Active,
			}
			rolePermResp = &responses.Role_permissions{
				RolePermissionId: strconv.FormatInt(v.RolePermissionId, 10),
				Role:             &roleResp,
				Permission:       &permResp,
				Action:           &actionResp,
				DateCreated:      v.DateCreated,
				DateModified:     v.DateModified,
				CreatedBy:        v.CreatedBy,
				ModifiedBy:       v.ModifiedBy,
				Active:           v.Active,
			}
		}
		var resp = responses.RolePermissionResponseDTO{StatusCode: 200, RolePermission: rolePermResp, StatusDesc: "Role Permission added"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Role_permissions
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Role_permissions
// @Failure 403
// @router / [get]
func (c *Role_permissionsController) GetAll() {
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

	l, err := models.GetAllRole_permissions(query, fields, sortby, order, offset, limit)
	if err != nil {
		var resp = responses.RolePermissionsResponseDTO{StatusCode: 604, RolePermissions: nil, StatusDesc: "Error getting permission ::: " + err.Error()}
		c.Data["json"] = resp
	} else {
		if l == nil {
			l = []interface{}{}
		}
		rolePermissions := []responses.Role_permissions{}
		for _, urs := range l {
			m := urs.(models.Role_permissions)

			roleResp := responses.Roles{
				RoleId:       m.Role.RoleId,
				Role:         m.Role.Role,
				Description:  m.Role.Description,
				DateCreated:  m.Role.DateCreated,
				DateModified: m.Role.DateModified,
				CreatedBy:    m.Role.CreatedBy,
				ModifiedBy:   m.Role.ModifiedBy,
				Active:       m.Role.Active,
			}
			permResp := responses.Permissions{
				PermissionId:          strconv.FormatInt(m.Permission.PermissionId, 10),
				Permission:            m.Permission.Permission,
				PermissionCode:        m.Permission.PermissionCode,
				PermissionDescription: m.Permission.PermissionDescription,
				DateCreated:           m.Permission.DateCreated,
				DateModified:          m.Permission.DateModified,
				CreatedBy:             m.Permission.CreatedBy,
				ModifiedBy:            m.Permission.ModifiedBy,
				Active:                m.Permission.Active,
			}
			actionResp := responses.Actions{
				ActionId:     strconv.FormatInt(m.Action.ActionId, 10),
				Action:       m.Action.Action,
				Description:  m.Action.Description,
				DateCreated:  m.Action.DateCreated,
				DateModified: m.Action.DateModified,
				CreatedBy:    m.Action.CreatedBy,
				ModifiedBy:   m.Action.ModifiedBy,
				Active:       m.Action.Active,
			}

			rolePermissions = append(rolePermissions, responses.Role_permissions{
				RolePermissionId: strconv.FormatInt(m.RolePermissionId, 10),
				Role:             &roleResp,
				Permission:       &permResp,
				Action:           &actionResp,
				DateCreated:      m.DateCreated,
				DateModified:     m.DateModified,
				CreatedBy:        m.CreatedBy,
				ModifiedBy:       m.ModifiedBy,
				Active:           m.Active,
			})
		}
		var resp = responses.RolePermissionsResponseDTO{StatusCode: 200, RolePermissions: &rolePermissions, StatusDesc: "Role Permissions fetched"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Role_permissions
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Role_permissions	true		"body for Role_permissions content"
// @Success 200 {object} models.Role_permissions
// @Failure 403 :id is not int
// @router /:id [put]
func (c *Role_permissionsController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Role_permissions{RolePermissionId: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateRole_permissionsById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Role_permissions
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *Role_permissionsController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	_, _ = strconv.ParseInt(idStr, 0, 64)

	v := requests.RemoveRolePermissionRequest{
		Role:           c.GetString("Role"),
		PermissionCode: c.GetString("PermissionCode"),
		Action:         c.GetString("Action"),
		RemovedBy:      c.GetString("RemovedBy"),
	}

	// Some clients send DELETE params as x-www-form-urlencoded body.
	if (v.Role == "" || v.PermissionCode == "" || v.Action == "" || v.RemovedBy == "") && len(c.Ctx.Input.RequestBody) > 0 {
		if formValues, err := url.ParseQuery(string(c.Ctx.Input.RequestBody)); err == nil {
			if v.Role == "" {
				v.Role = formValues.Get("Role")
			}
			if v.PermissionCode == "" {
				v.PermissionCode = formValues.Get("PermissionCode")
			}
			if v.Action == "" {
				v.Action = formValues.Get("Action")
			}
			if v.RemovedBy == "" {
				v.RemovedBy = formValues.Get("RemovedBy")
			}
		}
	}

	if v.Role == "" || v.PermissionCode == "" || v.Action == "" || v.RemovedBy == "" {
		json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	}

	logs.Info("Request json is ")
	logs.Info("Request json is ", string(c.Ctx.Input.RequestBody))
	logs.Info("Request body: ", string(c.Ctx.Input.RequestBody), " Role: ", c.GetString("Role"), " PermissionCode: ", c.GetString("PermissionCode"), " Action: ", c.GetString("Action"), " RemovedBy: ", c.GetString("RemovedBy"))
	logs.Info("Parsed request: ", v)

	statusCode := 400
	statusMessage := "Role permission not found"
	var roleResp *responses.Roles

	idStrr := v.Role
	roleId, _ := strconv.ParseInt(idStrr, 0, 64)
	logs.Info("Sending action: ", v.Action, " for role: ", v.Role, " and permission: ", v.PermissionCode)
	if roleP, err := models.GetRolePermissionByRoleActionPermission(roleId, v.Action, v.PermissionCode); err == nil {
		if roleP != nil {
			if err := models.DeleteRole_permissions(roleP.RolePermissionId); err == nil {
				// c.Data["json"] = "OK"
				logs.Info("Role permission deleted successfully")
				statusCode = 200
				statusMessage = "Role permission deleted successfully"

			} else {
				logs.Error("Error deleting role permission: %v", err)
				statusCode = 500
				statusMessage = "Error deleting role permission: " + err.Error()
			}
		}
	} else {
		logs.Error("Role permission not found")
		statusCode = 404
		statusMessage = "Role permission not found"
	}

	if statusCode == 200 {
		logs.Info("Role permission deleted successfully")
		if role, err := models.GetRolesById(roleId); err == nil && role != nil {
			roleResp = &responses.Roles{
				RoleId:       role.RoleId,
				Role:         role.Role,
				Description:  role.Description,
				DateCreated:  role.DateCreated,
				DateModified: role.DateModified,
				Active:       role.Active,
				CreatedBy:    role.CreatedBy,
				ModifiedBy:   role.ModifiedBy,
			}
		}
	}
	resp := responses.RoleResponseDTO{StatusCode: statusCode, Role: roleResp, StatusDesc: statusMessage}
	c.Data["json"] = resp
	c.ServeJSON()
}
