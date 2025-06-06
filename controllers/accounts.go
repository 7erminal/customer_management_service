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

// AccountsController operations for Accounts
type AccountsController struct {
	beego.Controller
}

// URLMapping ...
func (c *AccountsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Accounts
// @Param	body		body 	requests.CreateAccountRequestDTO	true		"body for Accounts content"
// @Success 201 {int} models.Accounts
// @Failure 403 body is empty
// @router / [post]
func (c *AccountsController) Post() {
	var q requests.CreateAccountRequestDTO
	json.Unmarshal(c.Ctx.Input.RequestBody, &q)

	if us, err := models.GetUsersById(q.UserId); err == nil {
		var v models.Accounts = models.Accounts{UserId: q.UserId, AccountNumber: us.PhoneNumber, Balance: 0, BalanceBefore: 0, DateCreated: time.Now(), DateModified: time.Now(), CreatedBy: 1, ModifiedBy: 1, Active: 1}
		if _, err := models.AddAccounts(&v); err == nil {
			resp := responses.AccountDTO{StatusCode: 200, Account: &v, StatusDesc: "Account Successfully Added"}
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = resp
		} else {
			resp := responses.AccountDTO{StatusCode: 302, Account: nil, StatusDesc: err.Error()}
			c.Data["json"] = resp
		}
	} else {
		resp := responses.AccountDTO{StatusCode: 301, Account: nil, StatusDesc: err.Error()}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Accounts by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Accounts
// @Failure 403 :id is empty
// @router /:id [get]
func (c *AccountsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetAccountsById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Accounts
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Accounts
// @Failure 403
// @router / [get]
func (c *AccountsController) GetAll() {
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

	l, err := models.GetAllAccounts(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Accounts
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Accounts	true		"body for Accounts content"
// @Success 200 {object} models.Accounts
// @Failure 403 :id is not int
// @router /:id [put]
func (c *AccountsController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Accounts{AccountId: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateAccountsById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Accounts
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *AccountsController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteAccounts(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
