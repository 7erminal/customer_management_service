package controllers

import (
	"customer_management_service/models"
	"customer_management_service/structs/requests"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/beego/beego/v2/core/logs"

	beego "github.com/beego/beego/v2/server/web"
)

// Customer_emergency_contactsController operations for Customer_emergency_contacts
type Customer_emergency_contactsController struct {
	beego.Controller
}

// URLMapping ...
func (c *Customer_emergency_contactsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Customer_emergency_contacts
// @Param	body		body 	requests.AddCustomerEmergencyContactRequestDTO	true		"body for Customer_emergency_contacts content"
// @Success 201 {int} models.CustomerEmergencyContactResponseDTO
// @Failure 403 body is empty
// @router / [post]
func (c *Customer_emergency_contactsController) Post() {
	var req requests.AddCustomerEmergencyContactRequestDTO
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	logs.Info("Request body: ", req)
	logs.Info("Request body: ", c.Ctx.Input.RequestBody)
	logs.Info("Customer ID: ", req.CustomerId)
	// customerId := strconv.FormatInt(req.CustomerId, 64)
	customer := models.Customers{CustomerId: req.CustomerId}
	var v models.Customer_emergency_contacts = models.Customer_emergency_contacts{Name: req.Name, Contact: req.PhoneNumber, Customer: &customer, DateCreated: time.Now(), DateModified: time.Now(), CreatedBy: 1, ModifiedBy: 1}

	if _, err := models.AddCustomer_emergency_contacts(&v); err == nil {
		c.Ctx.Output.SetStatus(200)
		var resp = models.CustomerEmergencyContactResponseDTO{StatusCode: 200, CustomerEmergencyContact: &v, StatusDesc: "Customer emergency contact created successfully"}
		c.Data["json"] = resp
	} else {
		logs.Info("Error adding customer emergency contact: ", err)
		// c.Ctx.Output.SetStatus(604)
		var resp = models.CustomerEmergencyContactResponseDTO{StatusCode: 604, CustomerEmergencyContact: nil, StatusDesc: "Error adding customer emergency contact"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Customer_emergency_contacts by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Customer_emergency_contacts
// @Failure 403 :id is empty
// @router /:id [get]
func (c *Customer_emergency_contactsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetCustomer_emergency_contactsById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Customer_emergency_contacts
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Customer_emergency_contacts
// @Failure 403
// @router / [get]
func (c *Customer_emergency_contactsController) GetAll() {
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

	l, err := models.GetAllCustomer_emergency_contacts(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Customer_emergency_contacts
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	requests.AddCustomerEmergencyContactRequestDTO	true		"body for Customer_emergency_contacts content"
// @Success 200 {object} models.CustomerEmergencyContactResponseDTO
// @Failure 403 :id is not int
// @router /:id [put]
func (c *Customer_emergency_contactsController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	var req requests.EditCustomerEmergencyContactRequestDTO
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	var v models.Customer_emergency_contacts = models.Customer_emergency_contacts{CustomerEmergencyContactId: id, Name: req.Name, Contact: req.PhoneNumber, DateModified: time.Now(), ModifiedBy: 1}
	// v := models.Customer_emergency_contacts{CustomerEmergencyContactId: id}
	if err := models.UpdateCustomer_emergency_contactsById(&v); err == nil {
		var resp = models.CustomerEmergencyContactResponseDTO{StatusCode: 200, CustomerEmergencyContact: &v, StatusDesc: "Customer emergency contact updated successfully"}
		c.Data["json"] = resp
	} else {
		// c.Data["json"] = err.Error()
		logs.Info("Error updating customer emergency contact: ", err)
		var resp = models.CustomerEmergencyContactResponseDTO{StatusCode: 604, CustomerEmergencyContact: nil, StatusDesc: "Error updating customer emergency contact"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Customer_emergency_contacts
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *Customer_emergency_contactsController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteCustomer_emergency_contacts(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
