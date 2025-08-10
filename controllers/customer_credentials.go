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

// Customer_credentialsController operations for Customer_credentials
type Customer_credentialsController struct {
	beego.Controller
}

// URLMapping ...
func (c *Customer_credentialsController) URLMapping() {
	c.Mapping("AddCustomerCredential", c.AddCustomerCredential)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// AddCustomerCredential ...
// @Title Post
// @Description create Customer_credentials
// @Param	body		body 	requests.CustomerCredentialRequestDTO	true		"body for Customer_credentials content"
// @Success 201 {int} models.Customer_credentials
// @Failure 403 body is empty
// @router /add-customer-credential [post]
func (c *Customer_credentialsController) AddCustomerCredential() {
	var v requests.CustomerCredentialRequestDTO
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	logs.Info("Adding customer credential for customer ID: ", v.CustomerId)

	statusCode := 401
	statusDesc := "Unauthorized"

	customer := models.Customers{CustomerId: v.CustomerId}

	hashedPassword, errr := bcrypt.GenerateFromPassword([]byte(v.Password), 8)
	if errr != nil {
		logs.Error("Error hashing password: %v", errr)
		c.Data["json"] = &responses.CustomerCredentialsResponseDTO{
			StatusCode: 500,
			StatusDesc: "Internal Server Error",
			Result:     "Error hashing password",
		}
		c.Ctx.Output.SetStatus(500)
		return
	}

	// hashedPin, errr := bcrypt.GenerateFromPassword([]byte(v.Pin), 8)
	// if errr != nil {
	// 	logs.Error("Error hashing password: %v", errr)
	// 	c.Data["json"] = &responses.CustomerCredentialsResponseDTO{
	// 		StatusCode: 500,
	// 		StatusDesc: "Internal Server Error",
	// 		Result:     "Error hashing password",
	// 	}
	// 	c.Ctx.Output.SetStatus(500)
	// 	return
	// }

	ccredential := models.Customer_credentials{
		Customer:     &customer,
		Username:     v.Username,
		Password:     string(hashedPassword),
		Pin:          v.Pin,
		DateCreated:  time.Now(),
		DateModified: time.Now(),
		CreatedBy:    1,
		ModifiedBy:   1,
		Active:       1,
	}

	if _, err := models.AddCustomer_credentials(&ccredential); err == nil {
		c.Ctx.Output.SetStatus(200)

		statusCode = 200
		statusDesc = "Customer credential added successfully"

		logs.Info("Customer credential added successfully: ", ccredential)
		c.Data["json"] = &responses.CustomerCredentialsResponseDTO{
			StatusCode: statusCode,
			StatusDesc: statusDesc,
			Result:     "Customer credential added successfully",
		}
		// c.Data["json"] = v
	} else {
		c.Data["json"] = err.Error()

		logs.Error("Error adding customer credential: %v", err)
		c.Ctx.Output.SetStatus(200)

		c.Data["json"] = &responses.CustomerCredentialsResponseDTO{
			StatusCode: statusCode,
			StatusDesc: statusDesc,
			Result:     err.Error(),
		}
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Customer_credentials by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Customer_credentials
// @Failure 403 :id is empty
// @router /:id [get]
func (c *Customer_credentialsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetCustomer_credentialsById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Customer_credentials
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Customer_credentials
// @Failure 403
// @router / [get]
func (c *Customer_credentialsController) GetAll() {
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

	l, err := models.GetAllCustomer_credentials(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Customer_credentials
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	requests.CustomerCredentialUpdateRequestDTO	true		"body for Customer_credentials content"
// @Success 200 {object} models.Customer_credentials
// @Failure 403 :customerId is not int
// @router /update-customer-credential/:customerId [put]
func (c *Customer_credentialsController) Put() {
	idStr := c.Ctx.Input.Param(":customerId")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	// v := models.Customer_credentials{Id: id}

	var v requests.CustomerCredentialUpdateRequestDTO
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	logs.Info("Adding customer credential for customer ID: ", idStr)

	statusCode := 401
	statusDesc := "Unauthorized"

	customer := models.Customers{CustomerId: id}

	if customerC, err := models.GetCustomer_credentialsByCustomerId(customer); err == nil {

		hashedPassword, errr := bcrypt.GenerateFromPassword([]byte(v.Password), 8)
		if errr != nil {
			logs.Error("Error hashing password: %v", errr)
			c.Data["json"] = &responses.CustomerCredentialsResponseDTO{
				StatusCode: 500,
				StatusDesc: "Internal Server Error",
				Result:     "Error hashing password",
			}
			c.Ctx.Output.SetStatus(500)
			return
		}

		// hashedPin, errr := bcrypt.GenerateFromPassword([]byte(v.Pin), 8)
		// if errr != nil {
		// 	logs.Error("Error hashing password: %v", errr)
		// 	c.Data["json"] = &responses.CustomerCredentialsResponseDTO{
		// 		StatusCode: 500,
		// 		StatusDesc: "Internal Server Error",
		// 		Result:     "Error hashing password",
		// 	}
		// 	c.Ctx.Output.SetStatus(500)
		// 	return
		// }

		ccredential := models.Customer_credentials{
			Id:           customerC.Id,
			Customer:     &customer,
			Username:     v.Username,
			Password:     string(hashedPassword),
			Pin:          v.Pin,
			DateCreated:  time.Now(),
			DateModified: time.Now(),
			CreatedBy:    1,
			ModifiedBy:   1,
			Active:       1,
		}

		// json.Unmarshal(c.Ctx.Input.RequestBody, &v)
		if err := models.UpdateCustomer_credentialsById(&ccredential); err == nil {
			c.Data["json"] = "OK"

			c.Ctx.Output.SetStatus(200)
			statusCode = 200
			statusDesc = "Customer credential updated successfully"
			logs.Info("Customer credential updated successfully: ", ccredential)
			c.Data["json"] = &responses.CustomerCredentialsResponseDTO{
				StatusCode: statusCode,
				StatusDesc: statusDesc,
				Result:     "Customer credential updated successfully",
			}
		} else {
			// c.Data["json"] = err.Error()
			logs.Error("Error updating customer credential: %v", err)
			c.Ctx.Output.SetStatus(200)
			c.Data["json"] = &responses.CustomerCredentialsResponseDTO{
				StatusCode: statusCode,
				StatusDesc: statusDesc,
				Result:     err.Error(),
			}
		}
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Customer_credentials
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *Customer_credentialsController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteCustomer_credentials(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
