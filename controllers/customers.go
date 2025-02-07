package controllers

import (
	"customer_management_service/models"
	"customer_management_service/structs/requests"
	"customer_management_service/structs/responses"
	"encoding/json"
	"errors"
	"net/http"
	"path/filepath"
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
	c.Mapping("UpdateCustomerImage", c.UpdateCustomerImage)
	c.Mapping("UpdateCustomerLastTxnDate", c.UpdateCustomerLastTxnDate)
	c.Mapping("GetAllByBranch", c.GetAllByBranch)
	c.Mapping("GetCustomerCount", c.GetCustomerCount)
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
// @Param	CustomerImage		formData 	file	true		"Customer Image"
// @Success 201 {object} models.CustomerResponseDTO
// @Failure 403 body is empty
// @router /add-customer [post]
func (c *CustomersController) AddCustomer() {
	// image of user received
	file, header, err := c.GetFile("CustomerImage")
	var filePath string = ""

	if err != nil {
		// c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": "Failed to get image file."}
		logs.Info("Failed to get the file ", err)
		// c.ServeJSON()
		// return
	} else {
		defer file.Close()

		// Save the uploaded file
		fileName := filepath.Base(header.Filename)
		filePath = "/uploads/customers/" + time.Now().Format("20060102150405") + fileName // Define your file path
		err = c.SaveToFile("CustomerImage", "../images/"+filePath)
		if err != nil {
			c.Ctx.Output.SetStatus(http.StatusInternalServerError)
			logs.Error("Error saving file", err)
			// c.Data["json"] = map[string]string{"error": "Failed to save the image file."}
			// errorMessage := "Error: Failed to save the image file"

			// resp := responses.UserResponseDTO{StatusCode: 601, User: nil, StatusDesc: "Error updating user. " + errorMessage}

			// c.Data["json"] = resp
			// c.ServeJSON()
			// return
		}
	}

	defaultPassword := "password1234"

	hashedPassword, errr := bcrypt.GenerateFromPassword([]byte(defaultPassword), 8)

	if errr == nil {
		logs.Debug(hashedPassword)

		// defaultPassword = string(hashedPassword)

		// logs.Debug("Category received is ", rcategory)

		// models.Agents{AgentName: v.AgentName, BranchId: v.BranchId, IdType: v.IdType, IdNumber: v.IdNumber, IsVerified: false, Active: 1, DateCreated: time.Now(), DateModified: time.Now(), CreatedBy: c_by, ModifiedBy: c_by}
	}

	rdob := c.Ctx.Input.Query("Dob")
	rcategoryid := c.Ctx.Input.Query("Category")
	remail := c.Ctx.Input.Query("Email")
	rname := c.Ctx.Input.Query("Name")
	rnickname := c.Ctx.Input.Query("Nickname")
	rphonenumber := c.Ctx.Input.Query("PhoneNumber")
	ridtype := c.Ctx.Input.Query("IdType")
	ridnumber := c.Ctx.Input.Query("IdNumber")
	rlocation := c.Ctx.Input.Query("Location")
	rbranch := c.Ctx.Input.Query("Branch")
	raddedBy := c.Ctx.Input.Query("AddedBy")

	logs.Info("Received :::", rdob, " ::: ", rcategoryid, " ::: ", remail, " ::: ", rname, " ::: ", " ::: ", rnickname)

	var proceed bool = true

	var dobm time.Time

	var allowedDateList [4]string = [4]string{"2006-01-02", "2006/01/02", "2006-01-02 15:04:05.000", "2006/01/02 15:04:05.000"}

	for _, date_ := range allowedDateList {
		logs.Debug("About to convert ", rdob)
		logs.Debug("About to convert ", c.Ctx.Input.Query("Dob"))
		// Convert dob string to date
		tdobm, error := time.Parse(date_, rdob)

		if error != nil {
			logs.Error("Error parsing date", error)
			proceed = false
		} else {
			logs.Info("Date converted to time successfully", tdobm)
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
		ccid, _ := strconv.ParseInt(rcategoryid, 0, 64)
		if cc, errr := models.GetCustomer_categoriesById(ccid); errr != nil {
			logs.Error("Customer category does not exist")
			if cc, errr := models.GetCustomer_categoriesByName(rcategoryid); errr != nil {
				logs.Error("Customer category does not exist")
			} else {
				category = *cc
			}
		} else {
			category = *cc
		}

		idType := models.Identification_types{}
		idT, _ := strconv.ParseInt(ridtype, 10, 64)
		if idtype_, err := models.GetIdentification_typesById(idT); err != nil {
			logs.Error("ID Type provided: ", ridtype, " does not exist ", err.Error())
			if idtype_, err := models.GetIdentification_typesByCode(ridtype); err != nil {
				logs.Error("ID Type provided: ", ridtype, " does not exist ", err.Error())
			} else {
				idType = *idtype_
			}
		} else {
			idType = *idtype_
		}

		branch := models.Branches{}
		idB, _ := strconv.ParseInt(rbranch, 10, 64)
		if branch_, err := models.GetBranchesById(idB); err != nil {
			logs.Error("Branch provided: ", rbranch, " does not exist ", err.Error())
			if branch_, err := models.GetBranchesByName(rbranch); err != nil {
				logs.Error("Branch provided: ", rbranch, " does not exist ", err.Error())
			} else {
				branch = *branch_
			}
		} else {
			branch = *branch_
		}

		user, _ := strconv.ParseInt(raddedBy, 10, 64)

		var cust = models.Customers{FullName: rname, Branch: &branch, ImagePath: filePath, PhoneNumber: rphonenumber, Location: rlocation, Email: remail, Dob: dobm, IdentificationType: &idType, IdentificationNumber: ridnumber, Shop: nil, Nickname: rnickname, CustomerCategory: &category, DateCreated: time.Now(), DateModified: time.Now(), Active: 1, CreatedBy: int(user), ModifiedBy: int(user)}

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
		logs.Error("Error fetching customers ", err.Error())
		resp := responses.StringResponseDTO{StatusCode: 301, Value: err.Error(), StatusDesc: "Error fetching customers"}
		c.Data["json"] = resp
	} else {
		if l == nil {
			l = []interface{}{}
		}
		resp := responses.CustomersDTO{StatusCode: 200, Customers: &l, StatusDesc: "Successfully fetched customers"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// GetAllByBranch ...
// @Title Get All By Branch
// @Description get Customers By Branch
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Customers
// @Failure 403
// @router /branch/:id [get]
func (c *CustomersController) GetAllByBranch() {
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

	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)

	if br, err := models.GetBranchesById(id); err == nil {
		l, err := models.GetAllCustomersByBranch(br, query, fields, sortby, order, offset, limit)
		if err != nil {
			logs.Error("Error fetching customers ", err.Error())
			resp := responses.StringResponseDTO{StatusCode: 301, Value: err.Error(), StatusDesc: "Error fetching customers"}
			c.Data["json"] = resp
		} else {
			if l == nil {
				l = []interface{}{}
			}
			resp := responses.CustomersDTO{StatusCode: 200, Customers: &l, StatusDesc: "Successfully fetched customers"}
			c.Data["json"] = resp
		}
	} else {
		resp := responses.CustomersDTO{StatusCode: 608, Customers: nil, StatusDesc: "Branch does not exist"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Customers
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.UpdateCustomerRequestDTO	true		"body for Customers content"
// @Param	CustomerImage		formData 	file	true		"Customer Image"
// @Success 200 {object} models.CustomerResponseDTO
// @Failure 403 :id is not int
// @router /:id [put]
func (c *CustomersController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := requests.UpdateCustomerRequestDTO{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	// image of user received
	file, header, err := c.GetFile("CustomerImage")
	var filePath string = ""

	if err != nil {
		// c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": "Failed to get image file."}
		logs.Info("Failed to get the file ", err)
		// c.ServeJSON()
		// return
	} else {
		defer file.Close()

		// Save the uploaded file
		fileName := filepath.Base(header.Filename)
		filePath = "/uploads/customers/" + time.Now().Format("20060102150405") + fileName // Define your file path
		err = c.SaveToFile("CustomerImage", "../images/"+filePath)
		if err != nil {
			c.Ctx.Output.SetStatus(http.StatusInternalServerError)
			logs.Error("Error saving file", err)
			// c.Data["json"] = map[string]string{"error": "Failed to save the image file."}
			// errorMessage := "Error: Failed to save the image file"

			// resp := responses.UserResponseDTO{StatusCode: 601, User: nil, StatusDesc: "Error updating user. " + errorMessage}

			// c.Data["json"] = resp
			// c.ServeJSON()
			// return
		}
	}

	if cust, err := models.GetCustomerById(id); err == nil {
		remail := c.Ctx.Input.Query("Email")
		rname := c.Ctx.Input.Query("Name")
		rnickname := c.Ctx.Input.Query("Nickname")
		rphonenumber := c.Ctx.Input.Query("PhoneNumber")
		ridtype := c.Ctx.Input.Query("IdType")
		ridnumber := c.Ctx.Input.Query("IdNumber")
		rlocation := c.Ctx.Input.Query("Location")
		rbranch := c.Ctx.Input.Query("Branch")
		rmodifiedby := c.Ctx.Input.Query("ModifiedBy")
		user, _ := strconv.ParseInt(rmodifiedby, 10, 64)
		if filePath == "" {
			filePath = cust.ImagePath
		}

		cust.FullName = rname
		cust.Email = remail
		cust.PhoneNumber = rphonenumber
		cust.IdentificationNumber = ridnumber
		cust.Nickname = rnickname
		cust.Location = rlocation
		cust.ImagePath = filePath
		cust.ModifiedBy = int(user)

		idT, _ := strconv.ParseInt(ridtype, 10, 64)
		if idtype, err := models.GetIdentification_typesById(idT); err != nil {

			logs.Error("ID Type provided does not exist")
			if idtype, err := models.GetIdentification_typesByCode(ridtype); err != nil {
				logs.Error("ID Type provided: ", ridtype, " does not exist ", err.Error())
			} else {
				cust.IdentificationType = idtype
			}
		} else {
			cust.IdentificationType = idtype
		}

		// category := models.Customer_categories{}
		// ccid, _ := strconv.ParseInt(rcategoryid, 0, 64)
		// if cc, errr := models.GetCustomer_categoriesById(ccid); errr != nil {
		// 	logs.Error("Customer category does not exist")
		// 	if cc, errr := models.GetCustomer_categoriesByName(rcategoryid); errr != nil {
		// 		logs.Error("Customer category does not exist")
		// 	} else {
		// 		category = *cc
		// 	}
		// } else {
		// 	category = *cc
		// }

		idB, _ := strconv.ParseInt(rbranch, 10, 64)
		if branch_, err := models.GetBranchesById(idB); err != nil {
			logs.Error("Branch provided: ", rbranch, " does not exist ", err.Error())
			if branch_, err := models.GetBranchesByName(rbranch); err != nil {
				logs.Error("Branch provided: ", rbranch, " does not exist ", err.Error())
			} else {
				cust.Branch = branch_
			}
		} else {
			cust.Branch = branch_
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

// Update Customer last Txn date ...
// @Title Update Customer's last txn date
// @Description update the Customer's least txn dat
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.UpdateCustomerLastTxnRequest	true		"body for Customers content"
// @Param	CustomerImage		formData 	file	true		"Customer Image"
// @Success 200 {object} models.CustomerResponseDTO
// @Failure 403 :id is not int
// @router /last-txn/:id [put]
func (c *CustomersController) UpdateCustomerLastTxnDate() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := requests.UpdateCustomerLastTxnRequest{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if cust, err := models.GetCustomerById(id); err == nil {
		var lastTxnDate time.Time

		var allowedDateList [4]string = [4]string{"2006-01-02", "2006/01/02", "2006-01-02 15:04:05.000", "2006/01/02 15:04:05.000"}

		for _, date_ := range allowedDateList {
			logs.Debug("About to convert ", v.TransactionDate)
			// Convert dob string to date
			tlastTxnDate, error := time.Parse(date_, v.TransactionDate)

			if error != nil {
				logs.Error("Error parsing date", error)
			} else {
				logs.Error("Date converted to time successfully", tlastTxnDate)
				lastTxnDate = tlastTxnDate

				break
			}
		}

		cust.LastTxnDate = lastTxnDate
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

// UpdateCustomerImage ...
// @Title Update customer's profile image
// @Description update the customer's profile image
// @Param	CustomerImage		formData 	file	true		"Customer Image"
// @Param	CustomerId		formData 	string	true		"Customer ID"
// @Success 200 {object} models.UserResponseDTO
// @Failure 403 body is empty
// @router /update-customer-image [post]
func (c *CustomersController) UpdateCustomerImage() {
	// image of user received
	file, header, err := c.GetFile("CustomerImage")
	var filePath string = ""

	if err != nil {
		// c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": "Failed to get image file."}
		logs.Info("Failed to get the file ", err)
		// c.ServeJSON()
		// return
	} else {
		defer file.Close()

		// Save the uploaded file
		fileName := header.Filename
		logs.Info("File name is ", fileName)
		filePath = "/uploads/customers/" + fileName // Define your file path
		logs.Info("File name is ", filePath)
		err = c.SaveToFile("UserImage", "."+filePath)
		if err != nil {
			c.Ctx.Output.SetStatus(http.StatusInternalServerError)
			logs.Error("Error saving file", err)
			// c.Data["json"] = map[string]string{"error": "Failed to save the image file."}
			errorMessage := "Error: Failed to save the image file"

			resp := models.CustomerResponseDTO{StatusCode: 601, Customer: nil, StatusDesc: "Error updating user. " + errorMessage}

			c.Data["json"] = resp
			c.ServeJSON()
			return
		}
	}

	id, _ := strconv.ParseInt(c.Ctx.Input.Query("UserId"), 10, 64)

	v, err := models.GetCustomerById(id)

	if err == nil {
		v.ImagePath = filePath

		if err := models.UpdateCustomerById(v); err == nil {

			logs.Debug("Returned customer is", v)

			var resp = models.CustomerResponseDTO{StatusCode: 200, Customer: v, StatusDesc: "Profile image updated successfully"}
			c.Data["json"] = resp
		} else {
			// c.Data["json"] = err.Error()
			logs.Debug("Error updating user", err.Error())
			var resp = models.CustomerResponseDTO{StatusCode: 602, Customer: nil, StatusDesc: "Error updating user"}
			c.Data["json"] = resp
		}
	} else {
		logs.Debug("Error fetching user")

		logs.Debug("Error updating user")
		var resp = models.CustomerResponseDTO{StatusCode: 603, Customer: nil, StatusDesc: "Error updating user"}
		c.Data["json"] = resp
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

// GetItemCount ...
// @Title Get Item Quantity
// @Description get Item_quantity by Item id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} responses.StringResponseDTO
// @Failure 403 :id is empty
// @router /count/ [get]
func (c *CustomersController) GetCustomerCount() {
	// q, err := models.GetItemsById(id)
	v, err := models.GetCustomerCount()
	count := strconv.FormatInt(v, 10)
	if err != nil {
		logs.Error("Error fetching count of customers ... ", err.Error())
		resp := responses.StringResponseDTO{StatusCode: 301, Value: "", StatusDesc: err.Error()}
		c.Data["json"] = resp
	} else {
		resp := responses.StringResponseDTO{StatusCode: 200, Value: count, StatusDesc: "Count fetched successfully"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}
