package controllers

import (
	"customer_management_service/controllers/functions"
	"customer_management_service/models"
	"customer_management_service/structs/responses"

	// "customer_management_service/structs/responses"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"golang.org/x/crypto/bcrypt"
)

// UsersController operations for Users
type UsersController struct {
	beego.Controller
}

// URLMapping ...
func (c *UsersController) URLMapping() {
	// c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("SignUp", c.SignUp)
	c.Mapping("SignUp2", c.SignUp2)
}

// SignUp2 ...
// @Title SignUp2
// @Description Sign up
// @Param	body		body 	models.UserCredentialsDTO	true		"body for SignUp content"
// @Success 201 {object} models.UserResponseDTO
// @Failure 403 body is empty
// @router /2/sign-up [post]
func (c *UsersController) SignUp2() {
	var v models.UserCredentialsDTO
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	logs.Info("Received ", v)

	hashedPassword, errr := bcrypt.GenerateFromPassword([]byte(v.Password), 8)

	// authorization := c.Ctx.Input.Header("Authorization")
	application := c.Ctx.Input.Header("Application")

	if errr == nil {
		logs.Debug(hashedPassword)

		v.Password = string(hashedPassword)

		logs.Debug("Sending", v.Password)

		// models.Agents{AgentName: v.AgentName, BranchId: v.BranchId, IdType: v.IdType, IdNumber: v.IdNumber, IsVerified: false, Active: 1, DateCreated: time.Now(), DateModified: time.Now(), CreatedBy: c_by, ModifiedBy: c_by}
	}

	// Convert dob string to date
	// dobm, error := time.Parse("2006-01-02 15:04:05.000", v.Dob)

	// Assign dob
	var addUserModel = models.Users{Username: v.Username, UserType: 1, Password: string(hashedPassword), DateCreated: time.Now(), DateModified: time.Now(), Active: 1, CreatedBy: 1, ModifiedBy: 1}

	if r, err := models.AddUsers(&addUserModel); err == nil {
		c.Ctx.Output.SetStatus(201)

		logs.Debug("Returned user A is", addUserModel)

		// id, _ := strconv.ParseInt(idStr, 0, 64)
		v, err := models.GetUsersById(r)

		if err != nil {
			c.Data["json"] = err.Error()

			logs.Error(err.Error())

			var resp = models.UserResponseDTO{StatusCode: 601, User: nil, StatusDesc: "Error fetching user"}
			c.Data["json"] = resp
		} else {
			logs.Debug("Returned user B is", v)

			// Shop here will be amended to cater for the shop that the customer is registering for

			var cust = models.Customers{User: v, Shop: nil, Nickname: "", DateCreated: time.Now(), DateModified: time.Now(), Active: 1, CreatedBy: 1, ModifiedBy: 1}

			if _, err := models.AddCustomers(&cust); err == nil {
				// Check application and register
				// If application is rides then create an account
				// Formulate request to send to create account
				if application == "RIDE" {
					functions.RegisterAccount(&c.Controller, addUserModel.UserId)
				}
				c.Ctx.Output.SetStatus(200)
				var resp = models.UserResponseDTO{StatusCode: 200, User: v, StatusDesc: "User created successfully"}

				c.Data["json"] = resp
			} else {
				// c.Data["json"] = err.Error()\
				logs.Error(err.Error())
				var resp = models.UserResponseDTO{StatusCode: 604, User: nil, StatusDesc: "Error adding customer"}
				c.Data["json"] = resp
			}

			// c.Data["json"] = v
		}
	} else {
		logs.Error(err.Error())

		var resp = models.UserResponseDTO{StatusCode: 606, User: nil, StatusDesc: "Error adding user"}
		c.Data["json"] = resp

		// c.Data["json"] = err.Error()
	}

	c.ServeJSON()
}

// SignUp ...
// @Title SignUp
// @Description Sign up
// @Param	body		body 	models.SignUpDTO	true		"body for SignUp content"
// @Success 201 {object} models.UserResponseDTO
// @Failure 403 body is empty
// @router /sign-up [post]
func (c *UsersController) SignUp() {
	var v models.SignUpDTO
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	logs.Info("Received ", v)

	// authorization := c.Ctx.Input.Header("Authorization")
	application := c.Ctx.Input.Header("Application")

	hashedPassword, errr := bcrypt.GenerateFromPassword([]byte(v.Password), 8)

	if errr == nil {
		logs.Debug(hashedPassword)

		v.Password = string(hashedPassword)

		logs.Debug("Sending", v.Password)

		// models.Agents{AgentName: v.AgentName, BranchId: v.BranchId, IdType: v.IdType, IdNumber: v.IdNumber, IsVerified: false, Active: 1, DateCreated: time.Now(), DateModified: time.Now(), CreatedBy: c_by, ModifiedBy: c_by}
	}

	q, err := models.GetUsersByUsername(v.Email)

	if err != nil {
		logs.Debug("About to debug")
		// Convert dob string to date
		dobm, error := time.Parse("2006-01-02", v.Dob)

		if error != nil {
			logs.Error(error)

			var resp = models.UserResponseDTO{StatusCode: 606, User: nil, StatusDesc: "Invalid date. Please enter date in the format (YYYY-MM-DD)."}
			c.Data["json"] = resp

			// c.Data["json"] = error.Error()

		} else {
			// Assign dob
			var gender string = strings.ToLower(v.Gender)

			if gender == "m" || gender == "M" || gender == "male" {
				gender = "MALE"
			}
			if gender == "f" || gender == "F" || gender == "female" {
				gender = "FEMALE"
			}
			var addUserModel = models.Users{FullName: v.Name, UserType: 1, Gender: gender, Dob: dobm, Password: string(hashedPassword), Email: v.Email, PhoneNumber: v.PhoneNumber, DateCreated: time.Now(), DateModified: time.Now(), Active: 1, CreatedBy: 1, ModifiedBy: 1}

			if _, err := models.AddUsers(&addUserModel); err == nil {
				c.Ctx.Output.SetStatus(201)

				logs.Debug("User is ", addUserModel)

				// logs.Debug("Returned user is", r)

				// id, _ := strconv.ParseInt(idStr, 0, 64)

				// logs.Debug("Returned user is", v)

				var cust = models.Customers{User: &addUserModel, Shop: nil, Nickname: "", DateCreated: time.Now(), DateModified: time.Now(), Active: 1, CreatedBy: 1, ModifiedBy: 1}

				if _, err := models.AddCustomers(&cust); err == nil {
					c.Ctx.Output.SetStatus(200)

					// Check application and register
					// If application is rides then create an account
					// Formulate request to send to create account
					if application == "RIDE" {
						functions.RegisterAccount(&c.Controller, addUserModel.UserId)
					}
					var resp = models.UserResponseDTO{StatusCode: 200, User: &addUserModel, StatusDesc: "User created successfully"}
					c.Data["json"] = resp
				} else {
					// c.Data["json"] = err.Error()
					logs.Error("Error adding customer, ", err.Error())
					var resp = models.UserResponseDTO{StatusCode: 604, User: nil, StatusDesc: "Error adding customer"}
					c.Data["json"] = resp
				}
				// c.Data["json"] = v

			} else {
				logs.Error(err.Error())

				var resp = models.UserResponseDTO{StatusCode: 606, User: nil, StatusDesc: "Error adding user"}
				c.Data["json"] = resp

				// c.Data["json"] = err.Error()
			}
		}
	} else {
		// c.Data["json"] = err.Error()
		var resp = models.UserResponseDTO{StatusCode: 604, User: q, StatusDesc: "User already exists. Username, email or mobile number already exists."}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// Verify Username ...
// @Title Verify User by username
// @Description Verify Users by username
// @Param	username		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.UserResponseDTO
// @Failure 403 :username is empty
// @router /get-user-by-username/:username [get]
func (c *UsersController) VerifyUsername() {
	username := c.Ctx.Input.Param(":username")
	v, err := models.GetUsersByUsername(username)

	if err != nil {
		logs.Error("Error::", err.Error())
		var resp = models.UserResponseDTO{StatusCode: 604, User: nil, StatusDesc: "Error getting user"}
		c.Data["json"] = resp
	} else {
		logs.Info("User found....sending user data")
		var resp = models.UserResponseDTO{StatusCode: 200, User: v, StatusDesc: "User details fetched"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// Verify User ...
// @Title Verify User by userid
// @Description Verify Users by userid
// @Param	username		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.UserResponseDTO
// @Failure 403 :username is empty
// @router /verify-user/:id [get]
func (c *UsersController) VerifyUser() {
	idStr := c.Ctx.Input.Param(":id")
	userid, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetUsersById(userid)

	if err != nil {
		logs.Error("Error::", err.Error())
		var resp = models.UserResponseDTO{StatusCode: 604, User: nil, StatusDesc: "Error getting user"}
		c.Data["json"] = resp
	} else {
		v.IsVerified = true
		if err := models.UpdateUsersById(v); err == nil {
			logs.Info("User found and verified....sending user data")
			var resp = models.UserResponseDTO{StatusCode: 200, User: v, StatusDesc: "User verified"}
			c.Data["json"] = resp
		} else {
			var resp = models.UserResponseDTO{StatusCode: 608, User: v, StatusDesc: "User not verified ::: " + err.Error()}
			c.Data["json"] = resp
		}
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Users by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Users
// @Failure 403 :id is empty
// @router /:id [get]
func (c *UsersController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetUsersById(id)
	if err != nil {
		var resp = models.UserResponseDTO{StatusCode: 604, User: nil, StatusDesc: "Error getting user ::: " + err.Error()}
		c.Data["json"] = resp
	} else {
		var resp = models.UserResponseDTO{StatusCode: 200, User: v, StatusDesc: "User details fetched"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Users
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Users
// @Failure 403
// @router / [get]
func (c *UsersController) GetAll() {
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

	l, err := models.GetAllUsers(query, fields, sortby, order, offset, limit)
	if err != nil {
		resp := responses.UsersAllCustomersDTO{StatusCode: 301, Users: nil, StatusDesc: "Fetch users failed ::: " + err.Error()}
		c.Data["json"] = resp
	} else {
		resp := responses.UsersAllCustomersDTO{StatusCode: 200, Users: &l, StatusDesc: "Users fetched successfully"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Users
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.UpdateUserRequestDTO	true		"body for Users content"
// @Success 200 {object} models.UserResponseDTO
// @Failure 403 :id is not int
// @router /:id [put]
func (c *UsersController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	var h models.UpdateUserRequestDTO

	// get the request
	json.Unmarshal(c.Ctx.Input.RequestBody, &h)

	logs.Debug("User id is ", id)

	// image of user received
	file, header, err := c.GetFile("UserImage")
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
		filePath = "/uploads/users/" + fileName // Define your file path
		err = c.SaveToFile("Image", "."+filePath)
		if err != nil {
			c.Ctx.Output.SetStatus(http.StatusInternalServerError)
			logs.Error("Error saving file", err)
			// c.Data["json"] = map[string]string{"error": "Failed to save the image file."}
			errorMessage := "Error: Failed to save the image file"

			resp := models.UserResponseDTO{StatusCode: 601, User: nil, StatusDesc: "Error updating user. " + errorMessage}

			c.Data["json"] = resp
			c.ServeJSON()
			return
		}
	}

	v, err := models.GetUsersById(id)

	logs.Debug("About to save", v)
	logs.Debug("And error is ", err)

	if err == nil {
		logs.Debug("User fetched successfully")

		logs.Debug("Marital status", h.MaritalStatus)
		logs.Debug("Full Name", h.FullName)
		logs.Debug("Gender", h.Gender)
		logs.Debug("Phone number", h.PhoneNumber)
		logs.Debug("Date of birth", h.Dob)

		// Parse request in Users object
		// v := models.Users{UserId: id, FullName: h.FullName, Gender: h.Gender, PhoneNumber: h.PhoneNumber, MaritalStatus: h.MaritalStatus, Address: h.Address}

		v.FullName = c.Ctx.Input.Query("FullName")
		v.Gender = c.Ctx.Input.Query("Gender")
		v.PhoneNumber = c.Ctx.Input.Query("PhoneNumber")
		v.MaritalStatus = c.Ctx.Input.Query("MaritalStatus")
		v.Address = c.Ctx.Input.Query("Address")
		v.ImagePath = filePath
		// Convert dob string to date
		dobm, error := time.Parse("2006-01-02 15:04:05.000", c.Ctx.Input.Query("Dob"))

		logs.Debug("Converted date", dobm)

		if error != nil {
			logs.Debug("Converted date error", error)
		} else {
			// Assign dob
			v.Dob = dobm
		}

		logs.Debug("About to save", v)
		logs.Debug("DOB", dobm)
		logs.Debug("is verified?", v.IsVerified)

		if err := models.UpdateUsersById(v); err == nil {
			v, err := models.GetUsersById(v.UserId)

			if err != nil {
				c.Data["json"] = err.Error()

				var resp = models.UserResponseDTO{StatusCode: 601, User: nil, StatusDesc: "Error fetching user"}
				c.Data["json"] = resp
			} else {
				logs.Debug("Returned user is", v)

				var resp = models.UserResponseDTO{StatusCode: 200, User: v, StatusDesc: "Profile updated successfully"}
				c.Data["json"] = resp

				// c.Data["json"] = v
			}
		} else {
			// c.Data["json"] = err.Error()
			logs.Debug("Error updating user", err.Error())
			var resp = models.UserResponseDTO{StatusCode: 200, User: nil, StatusDesc: "Error updating user"}
			c.Data["json"] = resp
		}
	} else {
		logs.Debug("Error fetching user")

		logs.Debug("Error updating user")
		var resp = models.UserResponseDTO{StatusCode: 200, User: nil, StatusDesc: "Error updating user"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Users
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *UsersController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)

	v, err := models.GetUsersById(id)
	if err == nil {
		v.Active = 6
		if err := models.UpdateUsersById(v); err == nil {
			// if err := models.DeleteUsers(id); err == nil {
			// 	c.Data["json"] = "OK"
			// } else {
			// 	c.Data["json"] = err.Error()
			// }
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
