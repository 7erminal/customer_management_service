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
	"golang.org/x/crypto/bcrypt"
)

// CustomersController operations for Customers
type CustomersController struct {
	beego.Controller
}

// URLMapping ...
func (c *CustomersController) URLMapping() {
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("AddCustomer", c.AddCustomer)
}

// Post ...
// @Title Post
// @Description create Customers
// @Param	body		body 	models.Customers	true		"body for Customers content"
// @Success 201 {int} models.Customers
// @Failure 403 body is empty
// @router / [post]
// func (c *CustomersController) Post() {
// 	var v models.Customers
// 	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
// 	if _, err := models.AddCustomers(&v); err == nil {
// 		c.Ctx.Output.SetStatus(201)
// 		c.Data["json"] = v
// 	} else {
// 		c.Data["json"] = err.Error()
// 	}
// 	c.ServeJSON()
// }

// AddCustomer ...
// @Title AddCustomer
// @Description Add customer
// @Param	body		body 	requests.AddCustomerRequestDTO	true		"body for add customer content"
// @Success 201 {object} models.CustomerResponseDTO
// @Failure 403 body is empty
// @router /add-customer [post]
func (c *CustomersController) AddCustomer() {
	var v requests.AddCustomerRequestDTO
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	logs.Info("Received ", v)

	defaultPassword := "password1234"

	hashedPassword, errr := bcrypt.GenerateFromPassword([]byte(defaultPassword), 8)

	if errr == nil {
		logs.Debug(hashedPassword)

		defaultPassword = string(hashedPassword)

		logs.Debug("Category received is ", c.Ctx.Input.Query("Category"))

		// models.Agents{AgentName: v.AgentName, BranchId: v.BranchId, IdType: v.IdType, IdNumber: v.IdNumber, IsVerified: false, Active: 1, DateCreated: time.Now(), DateModified: time.Now(), CreatedBy: c_by, ModifiedBy: c_by}
	}

	logs.Info("Unmarshalled already:::", v.Dob, " ::: ", v.Category, " ::: ", v.Email, " ::: ", v.Name, " ::: ", " ::: ", v.Nickname)

	addedBy, _ := strconv.Atoi(v.AddedBy)

	var proceed bool = true

	var dobm time.Time

	var allowedDateList [4]string = [4]string{"2006-01-02", "2006/01/02", "2006-01-02 15:04:05.000", "2006/01/02 15:04:05.000"}

	for _, date_ := range allowedDateList {
		logs.Debug("About to convert ", v.Dob)
		logs.Debug("About to convert ", c.Ctx.Input.Query("Dob"))
		// Convert dob string to date
		tdobm, error := time.Parse(date_, v.Dob)

		if error != nil {
			logs.Error("Error parsing date", error)
			proceed = false
		} else {
			logs.Error("Date converted to time successfully", tdobm)
			dobm = tdobm
			proceed = true

			break
		}
	}

	if !proceed {

		var resp = models.CustomerResponseDTO{StatusCode: 606, Customer: nil, StatusDesc: "Error adding user"}
		c.Data["json"] = resp

		// c.Data["json"] = error.Error()

	} else {
		// Assign dob
		category := models.Customer_categories{}
		ccid, _ := strconv.ParseInt(v.Category, 0, 64)
		if cc, errr := models.GetCustomer_categoriesById(ccid); errr != nil {
			logs.Error("Customer category does not exist")
		} else {
			category = *cc
		}

		idType := models.Identification_types{}
		idT, _ := strconv.ParseInt(v.IdType, 10, 64)
		if idtype, err := models.GetIdentification_typesById(idT); err != nil {
			logs.Error("ID Type provided does not exist")
		} else {
			idType = *idtype
		}
		var cust = models.Customers{FullName: v.Name, PhoneNumber: v.PhoneNumber, Email: v.Email, Dob: dobm, IdentificationType: &idType, IdentificationNumber: v.IdNumber, Shop: nil, Nickname: v.Nickname, CustomerCategory: &category, DateCreated: time.Now(), DateModified: time.Now(), Active: 1, CreatedBy: addedBy, ModifiedBy: addedBy}

		if _, err := models.AddCustomer(&cust); err == nil {
			c.Ctx.Output.SetStatus(200)
			var resp = models.CustomerResponseDTO{StatusCode: 200, Customer: &cust, StatusDesc: "User created successfully"}
			c.Data["json"] = resp
		} else {
			// c.Data["json"] = err.Error()
			var resp = models.CustomerResponseDTO{StatusCode: 604, Customer: nil, StatusDesc: "Error adding customer"}
			c.Data["json"] = resp
		}

		// c.Data["json"] = v

	}

	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Customers by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Customers
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CustomersController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetCustomerById(id)
	if err != nil {
		logs.Error("An error occurred fetching customer")
		var resp = models.CustomerResponseDTO{StatusCode: 608, Customer: nil, StatusDesc: "Error fetching customer " + err.Error()}
		c.Data["json"] = resp
	} else {
		var resp = models.CustomerResponseDTO{StatusCode: 200, Customer: v, StatusDesc: "User created successfully"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Customers
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Customers
// @Failure 403
// @router / [get]
func (c *CustomersController) GetAll() {
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

	l, err := models.GetAllCustomers(query, fields, sortby, order, offset, limit)
	if err != nil {
		resp := responses.StringResponseDTO{StatusCode: 301, Value: err.Error(), StatusDesc: "Error fetching customers"}
		c.Data["json"] = resp
	} else {
		resp := responses.CustomersDTO{StatusCode: 200, Customers: &l, StatusDesc: "Successfully fetched categories"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Customers
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Customers	true		"body for Customers content"
// @Success 200 {object} models.CustomerResponseDTO
// @Failure 403 :id is not int
// @router /:id [put]
func (c *CustomersController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := requests.UpdateCustomerRequestDTO{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if cust, err := models.GetCustomerById(id); err == nil {
		cust.FullName = v.Name
		cust.Email = v.Email
		cust.PhoneNumber = v.PhoneNumber
		cust.IdentificationNumber = v.IdNumber
		cust.Nickname = v.Nickname
		idT, _ := strconv.ParseInt(v.IdType, 10, 64)
		if idtype, err := models.GetIdentification_typesById(idT); err != nil {

			logs.Error("ID Type provided does not exist")
		} else {
			cust.IdentificationType = idtype
		}
		if err := models.UpdateCustomerById(cust); err == nil {
			c.Ctx.Output.SetStatus(200)
			var resp = models.CustomerResponseDTO{StatusCode: 200, Customer: cust, StatusDesc: "Customer updated successfully"}
			c.Data["json"] = resp
		} else {
			logs.Error("Customer update failed ", err.Error())
			var resp = models.CustomerResponseDTO{StatusCode: 608, Customer: nil, StatusDesc: "Customer update failed"}
			c.Data["json"] = resp
		}
	}

	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Customers
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *CustomersController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteCustomer(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
