package controllers

import (
	"customer_management_service/models"
	"customer_management_service/structs/responses"
	"errors"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// BranchController operations for Branch
type BranchController struct {
	beego.Controller
}

// URLMapping ...
func (c *BranchController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Branch
// @Param	body		body 	models.Branch	true		"body for Branch content"
// @Success 201 {object} models.Branch
// @Failure 403 body is empty
// @router / [post]
func (c *BranchController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Branch by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Branch
// @Failure 403 :id is empty
// @router /:id [get]
func (c *BranchController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetCustomer_categoriesById(id)
	if err != nil {
		logs.Error("Error fetching branch details for ", id, " is ", err.Error())
		var resp = responses.BranchResponseDTO{StatusCode: 301, Result: nil, StatusDesc: "Error fetching branch details for " + strconv.FormatInt(id, 10) + " is " + err.Error()}
		c.Data["json"] = resp
	} else {
		logs.Info("Successfully fetched branch details for ", id)
		var resp = responses.BranchResponseDTO{StatusCode: 200, Result: &v, StatusDesc: "Successfully fetched branch details for " + strconv.FormatInt(id, 10)}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description get Branch
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Branch
// @Failure 403
// @router / [get]
func (c *BranchController) GetAll() {
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
		resp := responses.BranchesResponseDTO{StatusCode: 301, Result: nil, StatusDesc: "Error fetching category details"}
		c.Data["json"] = resp
	} else {
		branchesResp := []responses.BranchResp{}
		for _, br := range l {
			m := br.(models.Branches)
			branchesResp = append(branchesResp, responses.BranchResp{
				BranchId:     m.BranchId,
				BranchName:   m.Branch,
				Location:     m.Location,
				Active:       m.Active,
				DateCreated:  m.DateCreated,
				DateModified: m.DateModified,
				CreatedBy:    m.CreatedBy,
				ModifiedBy:   m.ModifiedBy,
			})
		}
		resp := responses.BranchesResponseDTO{StatusCode: 200, Result: &branchesResp, StatusDesc: "Successfully fetched categories"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Branch
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Branch	true		"body for Branch content"
// @Success 200 {object} models.Branch
// @Failure 403 :id is not int
// @router /:id [put]
func (c *BranchController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Branch
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *BranchController) Delete() {

}
