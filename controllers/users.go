package controllers

import (
	"customer_management_service/controllers/functions"
	"customer_management_service/models"
	"customer_management_service/structs/requests"
	"customer_management_service/structs/responses"
	"path/filepath"

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
	c.Mapping("VerifyUsername", c.VerifyUsername)
	c.Mapping("VerifyUser", c.VerifyUser)
	c.Mapping("VerifyInvite", c.VerifyInvite)
	c.Mapping("UpdateUserImage", c.UpdateUserImage)
	c.Mapping("UpdateUserInvite", c.UpdateUserInvite)
	c.Mapping("GetUserInvite", c.GetUserInvite)
	c.Mapping("GetUsersUnderBranch", c.GetUsersUnderBranch)
	c.Mapping("UpdateUserRole", c.UpdateUserRole)
	c.Mapping("UpdateUserBranch", c.UpdateUserBranch)
	c.Mapping("GetUserCount", c.GetUserCount)
	c.Mapping("Deactivate", c.Deactivate)
	c.Mapping("GetUserInvites", c.GetUserInvites)
	c.Mapping("InviteUser", c.InviteUser)
	c.Mapping("RevokeUserInvites", c.RevokeUserInvites)
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

			var resp = responses.UserResponseDTO{StatusCode: 601, User: nil, StatusDesc: "Error fetching user"}
			c.Data["json"] = resp
		} else {
			logs.Debug("Returned user B is", v)

			// Shop here will be amended to cater for the shop that the customer is registering for

			var userDetails = models.UserExtraDetails{User: v.UserId, Shop: nil, Nickname: "", DateCreated: time.Now(), DateModified: time.Now(), Active: 1, CreatedBy: 1, ModifiedBy: 1}

			if _, err := models.AddUserExtraDetails(&userDetails); err == nil {
				// Check application and register
				// If application is rides then create an account
				// Formulate request to send to create account
				if application == "RIDE" {
					functions.RegisterAccount(&c.Controller, addUserModel.UserId)
				}

				if err != nil {
					c.Data["json"] = err.Error()

					var resp = responses.UserResponseDTO{StatusCode: 601, User: nil, StatusDesc: "Error fetching user"}
					c.Data["json"] = resp
				} else {
					addUserModel.UserDetails = &userDetails
					if err := models.UpdateUsersById(&addUserModel); err == nil {

						// userResp := responses.UserResp{
						// 	UserId:        v.UserId,
						// 	ImagePath:     v.ImagePath,
						// 	UserType:      v.UserType,
						// 	FullName:      v.FullName,
						// 	Username:      v.Username,
						// 	Password:      v.Password,
						// 	Email:         v.Email,
						// 	PhoneNumber:   v.PhoneNumber,
						// 	Gender:        v.Gender,
						// 	Dob:           v.Dob,
						// 	Address:       v.Address,
						// 	IdType:        v.IdType,
						// 	IdNumber:      v.IdNumber,
						// 	MaritalStatus: v.MaritalStatus,
						// 	Active:        v.Active,
						// 	Role:          v.Role,
						// 	IsVerified:    v.IsVerified,
						// 	DateCreated:   v.DateCreated,
						// 	DateModified:  v.DateModified,
						// 	CreatedBy:     v.CreatedBy,
						// 	ModifiedBy:    v.ModifiedBy,
						// 	Branch:        cust.Branch,
						// }
						c.Ctx.Output.SetStatus(200)
						var resp = responses.UserResponseDTO{StatusCode: 200, User: &addUserModel, StatusDesc: "User created successfully"}

						c.Data["json"] = resp
					} else {
						logs.Error("Error updating customer ID for user ")
						var resp = responses.UserResponseDTO{StatusCode: 200, User: &addUserModel, StatusDesc: "User created successfully. Please check user"}
						c.Data["json"] = resp
					}
				}
			} else {
				// c.Data["json"] = err.Error()\
				logs.Error(err.Error())
				var resp = responses.UserResponseDTO{StatusCode: 604, User: nil, StatusDesc: "Error adding customer"}
				c.Data["json"] = resp
			}

			// c.Data["json"] = v
		}
	} else {
		logs.Error(err.Error())

		var resp = responses.UserResponseDTO{StatusCode: 606, User: nil, StatusDesc: "Error adding user"}
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

	roleId, _ := strconv.ParseInt(v.Role, 10, 64)

	role, err := models.GetRolesById(roleId)

	// logs.Info("Role in request is ", v.Role)

	var proceed bool = false

	if err != nil {
		logs.Error("Error fetching role:: ", err.Error())
		role = nil

		var resp = responses.UserResponseDTO{StatusCode: 606, User: nil, StatusDesc: "Invalid role specified. Please enter date in the format (YYYY-MM-DD)."}
		c.Data["json"] = resp

		if !v.RoleRequired {
			proceed = true
		}
	} else {
		proceed = true
	}

	logs.Info("The role fetched is ", role.Role)

	if _, err := models.GetUsersByUsername(v.Email); err != nil && proceed {
		logs.Debug("About to debug")

		proceed = false

		// Convert dob string to date
		// dobm, error := time.Parse("2006-01-02", v.Dob)
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
			var resp = responses.UserResponseDTO{StatusCode: 606, User: nil, StatusDesc: "Invalid date. Please enter date in the format (YYYY-MM-DD)."}
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
			var addUserModel = models.Users{FullName: v.Name, UserType: 1, Gender: gender, Dob: dobm, Password: string(hashedPassword), Email: v.Email, PhoneNumber: v.PhoneNumber, Role: role, DateCreated: time.Now(), DateModified: time.Now(), Active: 1, CreatedBy: 1, ModifiedBy: 1}

			if uid, err := models.AddUsers(&addUserModel); err == nil {
				c.Ctx.Output.SetStatus(201)

				logs.Debug("User is ", addUserModel)

				// logs.Debug("Returned user is", r)

				// id, _ := strconv.ParseInt(idStr, 0, 64)

				// logs.Debug("Returned user is", v)

				var userDetails = models.UserExtraDetails{User: uid, Shop: nil, Nickname: "", DateCreated: time.Now(), DateModified: time.Now(), Active: 1, CreatedBy: 1, ModifiedBy: 1}

				if _, err := models.AddUserExtraDetails(&userDetails); err == nil {
					c.Ctx.Output.SetStatus(200)

					// Check application and register
					// If application is rides then create an account
					// Formulate request to send to create account
					if application == "RIDE" {
						logs.Info("Ride application. Registering account")
						functions.RegisterAccount(&c.Controller, addUserModel.UserId)
					}

					addUserModel.UserDetails = &userDetails
					if err := models.UpdateUsersById(&addUserModel); err == nil {

						var resp = responses.UserResponseDTO{StatusCode: 200, User: &addUserModel, StatusDesc: "User created successfully"}
						c.Data["json"] = resp
					} else {
						logs.Error("Error updating customer ID for user ")
						var resp = responses.UserResponseDTO{StatusCode: 200, User: &addUserModel, StatusDesc: "User created successfully. Please check user"}
						c.Data["json"] = resp
					}

				} else {
					// c.Data["json"] = err.Error()
					logs.Error("Error adding customer, ", err.Error())
					var resp = responses.UserResponseDTO{StatusCode: 604, User: nil, StatusDesc: "Error adding customer"}
					c.Data["json"] = resp
				}
				// c.Data["json"] = v

			} else {
				logs.Error(err.Error())

				var resp = responses.UserResponseDTO{StatusCode: 606, User: nil, StatusDesc: "Error adding user"}
				c.Data["json"] = resp

				// c.Data["json"] = err.Error()
			}
		}
	} else {
		// c.Data["json"] = err.Error()
		var resp = responses.UserResponseDTO{StatusCode: 604, User: nil, StatusDesc: "User already exists. Username, email or mobile number already exists."}
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

	logs.Info("Get user request:: ", username)

	if err != nil {
		logs.Error("Error::", err.Error())
		var resp = responses.UserResponseDTO{StatusCode: 604, User: nil, StatusDesc: "Error getting user"}
		c.Data["json"] = resp
	} else {
		logs.Info("User found....sending user data")
		logs.Info("User data::: ", v)
		// cust, err := models.GetCustomersByUser(v.UserId)

		// if err != nil {
		// 	c.Data["json"] = err.Error()

		// 	var resp = responses.UserResponseDTO{StatusCode: 601, User: nil, StatusDesc: "Error fetching user"}
		// 	c.Data["json"] = resp
		// } else {

		// 	userResp := responses.UserResp{
		// 		UserId:        v.UserId,
		// 		ImagePath:     v.ImagePath,
		// 		UserType:      v.UserType,
		// 		FullName:      v.FullName,
		// 		Username:      v.Username,
		// 		Password:      v.Password,
		// 		Email:         v.Email,
		// 		PhoneNumber:   v.PhoneNumber,
		// 		Gender:        v.Gender,
		// 		Dob:           v.Dob,
		// 		Address:       v.Address,
		// 		IdType:        v.IdType,
		// 		IdNumber:      v.IdNumber,
		// 		MaritalStatus: v.MaritalStatus,
		// 		Active:        v.Active,
		// 		Role:          v.Role,
		// 		IsVerified:    v.IsVerified,
		// 		DateCreated:   v.DateCreated,
		// 		DateModified:  v.DateModified,
		// 		CreatedBy:     v.CreatedBy,
		// 		ModifiedBy:    v.ModifiedBy,
		// 		Branch:        cust.Branch,
		// 	}
		// 	var resp = responses.UserResponseDTO{StatusCode: 200, User: &userResp, StatusDesc: "User details fetched"}
		// 	c.Data["json"] = resp
		// }
		var resp = responses.UserResponseDTO{StatusCode: 200, User: v, StatusDesc: "User details fetched"}
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

	logs.Info("Get user by user ID", userid)

	if err != nil {
		logs.Error("Error::", err.Error())
		var resp = responses.UserResponseDTO{StatusCode: 604, User: nil, StatusDesc: "Error getting user"}
		c.Data["json"] = resp
	} else {
		v.IsVerified = true
		if err := models.UpdateUsersById(v); err == nil {
			logs.Info("User found and verified....sending user data")
			// cust, err := models.GetCustomersByUser(v.UserId)

			// if err != nil {
			// 	c.Data["json"] = err.Error()

			// 	var resp = responses.UserResponseDTO{StatusCode: 601, User: nil, StatusDesc: "Error fetching user"}
			// 	c.Data["json"] = resp
			// } else {
			// 	logs.Info("Customer found ", cust)

			// 	userResp := responses.UserResp{
			// 		UserId:        v.UserId,
			// 		ImagePath:     v.ImagePath,
			// 		UserType:      v.UserType,
			// 		FullName:      v.FullName,
			// 		Username:      v.Username,
			// 		Password:      v.Password,
			// 		Email:         v.Email,
			// 		PhoneNumber:   v.PhoneNumber,
			// 		Gender:        v.Gender,
			// 		Dob:           v.Dob,
			// 		Address:       v.Address,
			// 		IdType:        v.IdType,
			// 		IdNumber:      v.IdNumber,
			// 		MaritalStatus: v.MaritalStatus,
			// 		Active:        v.Active,
			// 		Role:          v.Role,
			// 		IsVerified:    v.IsVerified,
			// 		DateCreated:   v.DateCreated,
			// 		DateModified:  v.DateModified,
			// 		CreatedBy:     v.CreatedBy,
			// 		ModifiedBy:    v.ModifiedBy,
			// 		Branch:        cust.Branch,
			// 	}
			// 	var resp = responses.UserResponseDTO{StatusCode: 200, User: &userResp, StatusDesc: "User verified"}
			// 	c.Data["json"] = resp
			// }
			var resp = responses.UserResponseDTO{StatusCode: 200, User: v, StatusDesc: "User verified"}
			c.Data["json"] = resp
		} else {
			var resp = responses.UserResponseDTO{StatusCode: 608, User: nil, StatusDesc: "User not verified ::: " + err.Error()}
			c.Data["json"] = resp
		}
	}
	c.ServeJSON()
}

// SignUp ...
// @Title Invite user
// @Description Invite user using email
// @Param	body		body 	requests.RegisterInviteRequestDTO	true		"body for SignUp content"
// @Success 200 {object} responses.StringResponseDTO
// @Failure 403 body is empty
// @router /invite-user [post]
func (c *UsersController) InviteUser() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 50
	var offset int64

	var v requests.RegisterInviteRequestDTO
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	logs.Info("Received ", v)

	if _, err := models.GetUsersByUsername(v.Email); err != nil {
		logs.Info("User not found")
		proceed := false
		if uis, errr := models.GetAllUserInvitesByEmail(v.Email, query, fields, sortby, order, offset, limit); errr == nil {
			logs.Info("User invites returned ", uis)
			proceed = true
			for _, ui := range uis {
				m := ui.(models.UserInvites)
				if m.Status == "PENDING" {
					logs.Info("User invite exists but is pending")
					if m.InvitationToken.ExpiryDate.Before(time.Now()) {
						logs.Info("User invite expiry date is before now")
						status := "EXPIRED"
						m.Status = status

						_ = models.UpdateUserInvitesById(&m)
					} else {
						proceed = false
					}
				} else if m.Status == "ACCEPTED" {
					proceed = false
				} else {
					// if proceed == true{
					// 	verifyToken := functions.VerifyUserToken(&c.Controller, ui.InvitationToken.Token, ui.InvitationToken.Nonce, v.Email)
					// }
				}
			}
		} else {
			logs.Error("User not found in users. Proceed. ", errr.Error())
			proceed = true
		}

		if proceed {
			proceed = false
			message := "An error occurred"
			statCode := 608
			tokenResp := functions.GenerateToken(&c.Controller, v.Email, v.Role)

			logs.Info("Token resp is ", tokenResp.Value.Token)

			i, _ := strconv.ParseInt(v.InviteBy, 10, 64)
			status := "PENDING"

			inviteBy, err := models.GetUsersById(i)

			if err != nil {
				message = "User not found"
			}

			roleid, _ := strconv.ParseInt(v.Role, 10, 64)

			role, err := models.GetRolesById(roleid)
			if err != nil {
				logs.Error("Role provided not found for token ", tokenResp.Value.Token)
				message = "Role provided not found"
			} else {
				proceed = true
			}

			if proceed {
				proceed = false
				var userInvite models.UserInvites = models.UserInvites{InvitedBy: inviteBy, InvitationToken: tokenResp.Value.Token, Email: v.Email, Role: role, Status: status, Active: 1, DateCreated: time.Now(), DateModified: time.Now(), CreatedBy: 1, ModifiedBy: 1}

				ui, err := models.AddUserInvites(&userInvite)

				if err != nil {
					logs.Error("Error response received when adding invite ", err.Error())
					return
				} else {
					proceed = true
					statCode = 200
					message = "Sent"
				}

				if proceed {
					logs.Info("User invite added ", ui)
					logs.Info("Token to be sent is ", tokenResp.Value.Token.Token)

					link := v.Link + tokenResp.Value.Token.Token

					go functions.SendEmail(v.Email, link)

					logs.Info("Email sent")
				}
			}

			var resp = responses.StringResponseDTO{StatusCode: statCode, Value: message, StatusDesc: message}
			c.Data["json"] = resp
		} else {
			logs.Error("Unable to generate token ")
			var resp = responses.StringResponseDTO{StatusCode: 502, Value: "", StatusDesc: "Unable to generate token. Email exists. "}
			c.Data["json"] = resp
		}

	} else {
		logs.Error("User already exists. Invite token not generated")
		var resp = responses.StringResponseDTO{StatusCode: 502, Value: "", StatusDesc: "Unable to generate token. User exists."}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// VerifyInvite ...
// @Title Verify invite
// @Description Verify invite
// @Param	body		body 	requests.StringRequestDTO	true		"body for SignUp content"
// @Success 200 {object} responses.InviteDecodeResponseDTO
// @Failure 403 body is empty
// @router /verify-invite [post]
func (c *UsersController) VerifyInvite() {
	var v requests.StringRequestDTO
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	logs.Info("Received ", v)

	if token, err := models.GetUserTokensByToken(v.Value); err == nil {
		logs.Info("Token fetched from DB")
		logs.Info("Token expiry ")
		if userInvite, err := models.GetUserInvitesByToken(token); err == nil {
			if token.ExpiryDate.After(time.Now()) {
				if userInvite.Status != "PENDING" {
					var resp = responses.InviteDecodeResponseDTO{StatusCode: 608, Value: nil, StatusDesc: "Token verified failed. Token has been used already."}
					c.Data["json"] = resp
				} else {
					verifyTokenResp := functions.VerifyUserToken(&c.Controller, token.Token, token.Nonce, userInvite.InvitedBy.Email)
					if verifyTokenResp.StatusCode == 200 {
						userInvite.Status = "ACCEPTED"
						models.UpdateUserInvitesById(userInvite)
						var resp = responses.InviteDecodeResponseDTO{StatusCode: 200, Value: verifyTokenResp.Value, StatusDesc: "Token verified successfully"}
						c.Data["json"] = resp
					} else {
						var resp = responses.InviteDecodeResponseDTO{StatusCode: 501, Value: nil, StatusDesc: "Token verification failed"}
						c.Data["json"] = resp
					}
				}
			} else {
				status := "EXPIRED"
				userInvite.Status = status

				_ = models.UpdateUserInvitesById(userInvite)
				logs.Error("Token expired ")
				var resp = responses.InviteDecodeResponseDTO{StatusCode: 608, Value: nil, StatusDesc: "Token expired ::: " + err.Error()}
				c.Data["json"] = resp
			}
		} else {
			logs.Error("Unable to get specified token ", err.Error())
			var resp = responses.InviteDecodeResponseDTO{StatusCode: 608, Value: nil, StatusDesc: "Unable to get token ::: " + err.Error()}
			c.Data["json"] = resp
		}

	} else {
		logs.Error("Unable to get specified token ", err.Error())
		var resp = responses.InviteDecodeResponseDTO{StatusCode: 608, Value: nil, StatusDesc: "Unable to get token ::: " + err.Error()}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// RevokeInvite ...
// @Title Revoke invite
// @Description Revoke invite
// @Param	body		body 	requests.StringRequestDTO	true		"body for SignUp content"
// @Success 200 {object} responses.InviteDecodeResponseDTO
// @Failure 403 body is empty
// @router /revoke-invite [post]
func (c *UsersController) RevokeUserInvites() {
	var v requests.StringRequestDTO
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	logs.Info("Received ", v)

	if userInvites, err := models.GetUserInvitesByEmail(v.Value); err == nil {
		for _, invite := range userInvites {
			logs.Info("Each user invite is ", invite)
			status := "EXPIRED"
			invite.Status = status

			_ = models.UpdateUserInvitesById(invite)
			logs.Error("Token expired ")
			var resp = responses.StringResponseDTO{StatusCode: 200, Value: "DONE", StatusDesc: "Expired Invites"}
			c.Data["json"] = resp

		}
	} else {
		logs.Error("Unable to get specified token ", err.Error())
		var resp = responses.StringResponseDTO{StatusCode: 608, Value: "", StatusDesc: "Unable to get token ::: " + err.Error()}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// GetInvite ...
// @Title Get Invite
// @Description get User Invite by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Users
// @Failure 403 :id is empty
// @router /invite/:id [get]
func (c *UsersController) GetInvite() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetUserInvitesById(id)
	if err != nil {
		var resp = responses.UserInviteResponseDTO{StatusCode: 604, UserInvite: nil, StatusDesc: "Error getting user ::: " + err.Error()}
		c.Data["json"] = resp
	} else {
		logs.Info("Getting user invite ", v.InvitationToken)
		if verifyInvite := functions.VerifyUserToken(&c.Controller, v.InvitationToken.Token, v.InvitationToken.Nonce, v.Email); verifyInvite.StatusCode == 200 {
			var resp = responses.UserInviteResponseDTO{StatusCode: 200, UserInvite: v, StatusDesc: "User details fetched"}
			c.Data["json"] = resp
		} else {
			var resp = responses.UserInviteResponseDTO{StatusCode: 604, UserInvite: nil, StatusDesc: "Error verifying user"}
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
		var resp = responses.UserResponseDTO{StatusCode: 604, User: nil, StatusDesc: "Error getting user ::: " + err.Error()}
		c.Data["json"] = resp
	} else {
		logs.Info("Getting user details ", v.UserDetails)
		// cust, err := models.GetCustomersByUser(v.UserId)

		// if err != nil {
		// 	c.Data["json"] = err.Error()

		// 	var resp = responses.UserResponseDTO{StatusCode: 601, User: nil, StatusDesc: "Error fetching user"}
		// 	c.Data["json"] = resp
		// } else {
		// 	logs.Info("Getting the customer ", cust.Branch.Country.DefaultCurrency.CurrencyId)

		// 	userResp := responses.UserResp{
		// 		UserId:        v.UserId,
		// 		ImagePath:     v.ImagePath,
		// 		UserType:      v.UserType,
		// 		FullName:      v.FullName,
		// 		Username:      v.Username,
		// 		Password:      v.Password,
		// 		Email:         v.Email,
		// 		PhoneNumber:   v.PhoneNumber,
		// 		Gender:        v.Gender,
		// 		Dob:           v.Dob,
		// 		Address:       v.Address,
		// 		IdType:        v.IdType,
		// 		IdNumber:      v.IdNumber,
		// 		MaritalStatus: v.MaritalStatus,
		// 		Active:        v.Active,
		// 		Role:          v.Role,
		// 		IsVerified:    v.IsVerified,
		// 		DateCreated:   v.DateCreated,
		// 		DateModified:  v.DateModified,
		// 		CreatedBy:     v.CreatedBy,
		// 		ModifiedBy:    v.ModifiedBy,
		// 		Branch:        cust.Branch,
		// 	}
		// 	var resp = responses.UserResponseDTO{StatusCode: 200, User: &userResp, StatusDesc: "User details fetched"}
		// 	c.Data["json"] = resp
		// }
		var resp = responses.UserResponseDTO{StatusCode: 200, User: v, StatusDesc: "User details fetched"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Users
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	search	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
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
	var search = make(map[string]string)
	var limit int64 = 500
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

	// search: k:v,k:v
	if v := c.GetString("search"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid search key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			search[k] = v
		}
	}

	logs.Info("Getting all users")
	logs.Info("Query is ", query)

	l, err := models.GetAllUsers(query, fields, sortby, order, offset, limit, search)
	if err != nil {
		resp := responses.UsersAllCustomersDTO{StatusCode: 301, Users: nil, StatusDesc: "Fetch users failed ::: " + err.Error()}
		c.Data["json"] = resp
	} else {
		usersResp := []models.Users{}
		for _, urs := range l {
			m := urs.(models.Users)

			usersResp = append(usersResp, m)
		}
		resp := responses.UsersResponseDTO{StatusCode: 200, Users: &usersResp, StatusDesc: "Users fetched successfully"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// GetAllUsersWithRoleId ...
// @Title Get All Users with role Id
// @Description get Users with a role Id
// @Param	role_id		path 	string	true		"The key for staticblock"
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} responses.UsersAllCustomersDTO
// @Failure 403
// @router /role/:role_id [get]
func (c *UsersController) GetUsersWithRole() {
	role_idStr := c.Ctx.Input.Param(":role_id")
	role_id, _ := strconv.ParseInt(role_idStr, 0, 64)

	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 100
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

	if role, err := models.GetRolesById(role_id); err == nil {
		logs.Info("Role fetched ", role.Role)
		l, err := models.GetAllUsersWithRole(role, query, fields, sortby, order, offset, limit)
		if err != nil {
			resp := responses.UsersAllCustomersDTO{StatusCode: 301, Users: nil, StatusDesc: "Fetch users failed ::: " + err.Error()}
			c.Data["json"] = resp
		} else {
			logs.Info("Users fetched ", l)
			// usersResp := []models.Users{}
			// for _, urs := range l {
			// 	m := urs.(models.Users)

			// 	usersResp = append(usersResp, m)
			// }
			resp := responses.UsersAllCustomersDTO{StatusCode: 200, Users: &l, StatusDesc: "Users fetched successfully"}
			c.Data["json"] = resp
		}
	} else {
		logs.Error("Error getting role ", err.Error())
		resp := responses.UsersAllCustomersDTO{StatusCode: 301, Users: nil, StatusDesc: "Fetch users failed ::: " + err.Error()}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// GetAllUsersUnderBranch ...
// @Title Get All Users under branch
// @Description get Users under a branch
// @Param	branch_id		path 	string	true		"The key for staticblock"
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} responses.UsersAllCustomersDTO
// @Failure 403
// @router /branch/:branch_id [get]
func (c *UsersController) GetUsersUnderBranch() {
	branch_idStr := c.Ctx.Input.Param(":branch_id")
	branch_id, _ := strconv.ParseInt(branch_idStr, 0, 64)

	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 100
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

	if branch, err := models.GetBranchesById(branch_id); err == nil {
		logs.Info("Branch fetched ", branch.Branch)
		l, err := models.GetAllUsersByBranch(branch, query, fields, sortby, order, offset, limit)
		if err != nil {
			resp := responses.UsersAllCustomersDTO{StatusCode: 301, Users: nil, StatusDesc: "Fetch users failed ::: " + err.Error()}
			c.Data["json"] = resp
		} else {
			logs.Info("Users fetched ", l)
			// usersResp := []models.Users{}
			// for _, urs := range l {
			// 	m := urs.(models.Users)

			// 	usersResp = append(usersResp, m)
			// }
			resp := responses.UsersAllCustomersDTO{StatusCode: 200, Users: &l, StatusDesc: "Users fetched successfully"}
			c.Data["json"] = resp
		}
	} else {
		logs.Error("Error getting role ", err.Error())
		resp := responses.UsersAllCustomersDTO{StatusCode: 301, Users: nil, StatusDesc: "Fetch users failed ::: " + err.Error()}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// GetUserInvites ...
// @Title Get User Invites
// @Description get Users
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	search	query	string	false	"Filter. e.g. example@gmail.com ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} responses.UserInvitesResponseDTO
// @Failure 403
// @router /get-user-invites [get]
func (c *UsersController) GetUserInvites() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var search = make(map[string]string)
	var limit int64 = 100
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

		logs.Info("Query is ", query)
	}

	// search: k:v,k:v
	if v := c.GetString("search"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid search key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			search[k] = v
		}
	}

	l, err := models.GetAllUserInvites(query, fields, sortby, order, offset, limit, search)
	if err != nil {
		resp := responses.UserInvitesResponseDTO{StatusCode: 301, UserInvites: nil, StatusDesc: "Fetch user invites failed ::: " + err.Error()}
		c.Data["json"] = resp
	} else {
		// userInvites := []models.UserInvites{}
		// logs.Info("Date created for user invite is ", l[0])
		resp := responses.UserInvitesResponseDTO{StatusCode: 200, UserInvites: &l, StatusDesc: "User invites fetched successfully"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// UpdateUserImage ...
// @Title Update user's profile image
// @Description update the User's profile image
// @Param	UserImage		formData 	file	true		"User Image"
// @Param	UserId		formData 	string	true		"User ID"
// @Success 200 {object} models.UserResponseDTO
// @Failure 403 body is empty
// @router /update-user-image [post]
func (c *UsersController) UpdateUserImage() {
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
		logs.Info("File name is ", fileName)
		filePath = "/uploads/users/" + fileName // Define your file path
		logs.Info("File name is ", filePath)
		err = c.SaveToFile("UserImage", "."+filePath)
		if err != nil {
			filePath = ""
			c.Ctx.Output.SetStatus(http.StatusInternalServerError)
			logs.Error("Error saving file", err)
			// c.Data["json"] = map[string]string{"error": "Failed to save the image file."}
			errorMessage := "Error: Failed to save the image file"

			resp := responses.UserResponseDTO{StatusCode: 601, User: nil, StatusDesc: "Error updating user. " + errorMessage}

			c.Data["json"] = resp
			c.ServeJSON()
			return
		}
	}

	id, _ := strconv.ParseInt(c.Ctx.Input.Query("UserId"), 10, 64)

	v, err := models.GetUsersById(id)

	if err == nil {
		v.ImagePath = filePath

		if err := models.UpdateUsersById(v); err == nil {
			v, err := models.GetUsersById(v.UserId)

			if err != nil {
				c.Data["json"] = err.Error()

				var resp = responses.UserResponseDTO{StatusCode: 601, User: nil, StatusDesc: "Error fetching user"}
				c.Data["json"] = resp
			} else {
				logs.Debug("Returned user is", v)

				// cust, err := models.GetCustomersByUser(v)

				// if err != nil {
				// 	c.Data["json"] = err.Error()

				// 	var resp = responses.UserResponseDTO{StatusCode: 601, User: nil, StatusDesc: "Error fetching user"}
				// 	c.Data["json"] = resp
				// } else {

				// 	userResp := responses.UserResp{
				// 		UserId:        v.UserId,
				// 		ImagePath:     v.ImagePath,
				// 		UserType:      v.UserType,
				// 		FullName:      v.FullName,
				// 		Username:      v.Username,
				// 		Password:      v.Password,
				// 		Email:         v.Email,
				// 		PhoneNumber:   v.PhoneNumber,
				// 		Gender:        v.Gender,
				// 		Dob:           v.Dob,
				// 		Address:       v.Address,
				// 		IdType:        v.IdType,
				// 		IdNumber:      v.IdNumber,
				// 		MaritalStatus: v.MaritalStatus,
				// 		Active:        v.Active,
				// 		Role:          v.Role,
				// 		IsVerified:    v.IsVerified,
				// 		DateCreated:   v.DateCreated,
				// 		DateModified:  v.DateModified,
				// 		CreatedBy:     v.CreatedBy,
				// 		ModifiedBy:    v.ModifiedBy,
				// 		Branch:        cust.Branch,
				// 	}

				// 	var resp = responses.UserResponseDTO{StatusCode: 200, User: &userResp, StatusDesc: "Profile image updated successfully"}
				// 	c.Data["json"] = resp
				// }
				var resp = responses.UserResponseDTO{StatusCode: 200, User: v, StatusDesc: "Profile image updated successfully"}
				c.Data["json"] = resp
			}
		} else {
			// c.Data["json"] = err.Error()
			logs.Debug("Error updating user", err.Error())
			var resp = responses.UserResponseDTO{StatusCode: 602, User: nil, StatusDesc: "Error updating user"}
			c.Data["json"] = resp
		}
	} else {
		logs.Debug("Error fetching user")

		logs.Debug("Error updating user")
		var resp = responses.UserResponseDTO{StatusCode: 603, User: nil, StatusDesc: "Error updating user"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Users
// @Param	id		path 	string	true		"The id you want to update"
// @Param	UserImage		formData 	file	true		"User Image"
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

	logs.Info("Update user request received ", h)

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
		fileName := filepath.Base(header.Filename)
		filePath = "/uploads/users/" + time.Now().Format("20060102150405") + fileName // Define your file path
		err = c.SaveToFile("UserImage", "../images/"+filePath)
		if err != nil || header.Size < 1 {
			filePath = ""
			c.Ctx.Output.SetStatus(http.StatusInternalServerError)
			logs.Error("Error saving file", err)
			// c.Data["json"] = map[string]string{"error": "Failed to save the image file."}
			errorMessage := "Error: Failed to save the image file"

			resp := responses.UserResponseDTO{StatusCode: 601, User: nil, StatusDesc: "Error updating user. " + errorMessage}

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
		logs.Debug("Role fetched is ", h.RoleId)

		logs.Debug("File path is ", filePath, " and image path is ", v.ImagePath)
		remail := c.Ctx.Input.Query("Email")
		rname := c.Ctx.Input.Query("FullName")
		rgender := c.Ctx.Input.Query("Gender")
		rphonenumber := c.Ctx.Input.Query("PhoneNumber")
		rmaritalstatus := c.Ctx.Input.Query("MaritalStatus")
		raddress := c.Ctx.Input.Query("Address")
		rmodifiedby := c.Ctx.Input.Query("UpdatedBy")
		rroleid := c.Ctx.Input.Query("RoleId")
		rbranchid := c.Ctx.Input.Query("BranchId")
		dateofbirth := c.Ctx.Input.Query("Dob")

		if filePath == "" && v.ImagePath != "" {
			filePath = v.ImagePath
		}

		updatedBy, _ := strconv.ParseInt(rmodifiedby, 10, 64)

		if upu, err := models.GetUsersById(updatedBy); err == nil {
			updatedBy = upu.UserId
		}

		// Parse request in Users object
		// v := models.Users{UserId: id, FullName: h.FullName, Gender: h.Gender, PhoneNumber: h.PhoneNumber, MaritalStatus: h.MaritalStatus, Address: h.Address}

		v.FullName = rname
		v.Gender = rgender
		v.PhoneNumber = rphonenumber
		v.MaritalStatus = rmaritalstatus
		v.Address = raddress
		v.ImagePath = filePath
		v.Email = remail
		v.ModifiedBy = int(updatedBy)
		// Convert dob string to date
		dobm, error := time.Parse("2006-01-02", dateofbirth)

		logs.Debug("Converted date", dobm)

		if error != nil {
			logs.Debug("Converted date error", error.Error())
		} else {
			// Assign dob
			logs.Debug("DOB assigned")
			v.Dob = dobm
		}

		logs.Debug("About to save", v)
		logs.Debug("DOB", dobm, dateofbirth)
		logs.Debug("is verified?", v.IsVerified)

		// about to get role to update with
		roleid, _ := strconv.ParseInt(rroleid, 10, 64)
		if role, err := models.GetRolesById(roleid); err == nil {
			logs.Info("Role fetched for ", roleid)
			v.Role = role
		} else {
			logs.Error("There was an error getting the provided role")
		}

		if err := models.UpdateUsersById(v); err == nil {
			userDetails, err := models.GetUserExtraDetailsById(v.UserDetails.UserDetailsId)

			if err != nil {
				logs.Error("Error returned fetching customer ", err.Error())
				c.Data["json"] = err.Error()

				var resp = responses.UserResponseDTO{StatusCode: 601, User: nil, StatusDesc: "Error fetching user"}
				c.Data["json"] = resp
			} else {
				logs.Debug("Returned customer is", userDetails)
				branchid, _ := strconv.ParseInt(rbranchid, 10, 64)
				branch, err := models.GetBranchesById(branchid)

				if err != nil {
					logs.Error("Error fetching branch specified")
					branch = nil
				}

				userDetails.Branch = branch

				message := "Profile updated successfully"

				if err := models.UpdateUserExtraDetailsById(userDetails); err != nil {
					logs.Error("Failed to update customer branch")
					message = "Failed to update branch"
				}

				logs.Info("Branch saved for user is ", userDetails)
				// logs.Info("Branch saved for user is ", cust.Branch.)

				// userResp := responses.UserResp{
				// 	UserId:        v.UserId,
				// 	ImagePath:     v.ImagePath,
				// 	UserType:      v.UserType,
				// 	FullName:      v.FullName,
				// 	Username:      v.Username,
				// 	Password:      v.Password,
				// 	Email:         v.Email,
				// 	PhoneNumber:   v.PhoneNumber,
				// 	Gender:        v.Gender,
				// 	Dob:           v.Dob,
				// 	Address:       v.Address,
				// 	IdType:        v.IdType,
				// 	IdNumber:      v.IdNumber,
				// 	MaritalStatus: v.MaritalStatus,
				// 	Active:        v.Active,
				// 	Role:          v.Role,
				// 	IsVerified:    v.IsVerified,
				// 	DateCreated:   v.DateCreated,
				// 	DateModified:  v.DateModified,
				// 	CreatedBy:     v.CreatedBy,
				// 	ModifiedBy:    v.ModifiedBy,
				// 	Branch:        cust.Branch,
				// }

				var resp = responses.UserResponseDTO{StatusCode: 200, User: v, StatusDesc: message}
				c.Data["json"] = resp

				// c.Data["json"] = v
			}
		} else {
			// c.Data["json"] = err.Error()
			logs.Debug("Error updating user", err.Error())
			var resp = responses.UserResponseDTO{StatusCode: 608, User: nil, StatusDesc: "Error updating user"}
			c.Data["json"] = resp
		}
	} else {
		logs.Debug("Error fetching user")

		logs.Debug("Error updating user")
		var resp = responses.UserResponseDTO{StatusCode: 608, User: nil, StatusDesc: "Error updating user"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// Update User Role...
// @Title Update User Role
// @Description update the Users
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.UpdateUserRequestDTO	true		"body for Users content"
// @Success 200 {object} models.UserResponseDTO
// @Failure 403 :id is not int
// @router /role/:id [put]
func (c *UsersController) UpdateUserRole() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	var h models.UpdateUserRoleRequestDTO

	// get the request
	json.Unmarshal(c.Ctx.Input.RequestBody, &h)

	logs.Info("Update user request received ", h)

	logs.Debug("User id is ", id)

	v, err := models.GetUsersById(id)

	logs.Debug("About to save", v)
	logs.Debug("And error is ", err)

	if err == nil {
		logs.Debug("User fetched successfully")

		// about to get role to update with
		if role, err := models.GetRolesById(h.RoleId); err == nil {
			logs.Info("Role fetched for ", h.RoleId)
			v.Role = role

			if err := models.UpdateUsersById(v); err == nil {
				message := "Profile updated successfully"
				var resp = responses.UserResponseDTO{StatusCode: 200, User: v, StatusDesc: message}
				c.Data["json"] = resp

				// c.Data["json"] = v

			} else {
				// c.Data["json"] = err.Error()
				logs.Debug("Error updating user", err.Error())
				var resp = responses.UserResponseDTO{StatusCode: 608, User: nil, StatusDesc: "Error updating user"}
				c.Data["json"] = resp
			}
		} else {
			logs.Error("There was an error getting the provided role")

			var resp = responses.UserResponseDTO{StatusCode: 605, User: nil, StatusDesc: "Error updating user. Role not found."}
			c.Data["json"] = resp
		}

	} else {
		logs.Error("Error updating user ", err.Error())
		var resp = responses.UserResponseDTO{StatusCode: 604, User: nil, StatusDesc: "Error updating user. User not found."}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// Update User Branch...
// @Title Update User Branch
// @Description update the Users branch
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.UpdateUserBranchRequestDTO	true		"body for Users content"
// @Success 200 {object} responses.UserResponseDTO
// @Failure 403 :id is not int
// @router /branch/:id [put]
func (c *UsersController) UpdateUserBranch() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	var h models.UpdateUserBranchRequestDTO

	// get the request
	json.Unmarshal(c.Ctx.Input.RequestBody, &h)

	logs.Info("Update user request received ", h)

	logs.Debug("User id is ", id)

	v, err := models.GetUsersById(id)

	logs.Debug("About to save", v)
	logs.Debug("And error is ", err)

	if err == nil {
		logs.Debug("User fetched successfully")

		userDetails, err := models.GetUserExtraDetailsByUser(v.UserId)

		if err != nil {
			logs.Error("Error returned fetching user details ", err.Error())
			c.Data["json"] = err.Error()

			var resp = responses.UserResponseDTO{StatusCode: 601, User: nil, StatusDesc: "Error fetching user"}
			c.Data["json"] = resp
		} else {
			logs.Debug("Returned user details is", userDetails)
			branch, err := models.GetBranchesById(h.BranchId)
			message := "Profile updated successfully"

			if err != nil {
				logs.Error("Error fetching branch specified")
				message = "Branch not found"
				var resp = responses.UserResponseDTO{StatusCode: 605, User: nil, StatusDesc: message}
				c.Data["json"] = resp
			} else {
				userDetails.Branch = branch

				if err := models.UpdateUserExtraDetailsById(userDetails); err != nil {
					logs.Error("Failed to update user branch")
					message = "Failed to update branch"

					var resp = responses.UserResponseDTO{StatusCode: 608, User: v, StatusDesc: message}
					c.Data["json"] = resp
				} else {
					var resp = responses.UserResponseDTO{StatusCode: 200, User: v, StatusDesc: message}
					c.Data["json"] = resp
				}

			}

			logs.Info("Branch saved for user is ", userDetails)
			// logs.Info("Branch saved for user is ", cust.Branch.)

			// c.Data["json"] = v
		}

	} else {
		logs.Error("Error updating user ", err.Error())
		var resp = responses.UserResponseDTO{StatusCode: 604, User: nil, StatusDesc: "Error updating user. User not found."}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// GetOneUserInvite ...
// @Title Get One User Invite
// @Description get User invite by token
// @Param	token		path 	string	true		"The key for staticblock"
// @Success 200 {object} responses.UserInviteResponseDTO
// @Failure 403 :token is empty
// @router /user-invite/:token [get]
func (c *UsersController) GetUserInvite() {
	token := c.Ctx.Input.Param(":token")
	v, err := models.GetUserTokensByToken(token)
	if err != nil {
		logs.Error("Error getting user invite", err.Error())
		c.Ctx.Output.SetStatus(200)
		var resp = responses.UserInviteResponseDTO{StatusCode: 608, UserInvite: nil, StatusDesc: "Error getting user invite"}
		c.Data["json"] = resp
	} else {
		if q, err := models.GetUserInvitesByToken(v); err == nil {
			c.Ctx.Output.SetStatus(200)
			var resp = responses.UserInviteResponseDTO{StatusCode: 200, UserInvite: q, StatusDesc: "User invite fetch successful"}
			c.Data["json"] = resp
		} else {
			logs.Error("Error getting user invite")
			c.Ctx.Output.SetStatus(200)
			var resp = responses.UserInviteResponseDTO{StatusCode: 608, UserInvite: nil, StatusDesc: "Error getting user invite"}
			c.Data["json"] = resp
		}

	}
	c.ServeJSON()
}

// UpdateUserInvite ...
// @Title Put
// @Description update the Users
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	requests.UpdateUserInviteRequest	true		"body for Users content"
// @Success 200 {object} responses.UserInvitesResponseDTO
// @Failure 403 :id is not int
// @router /update-user-invite/:id [put]
func (c *UsersController) UpdateUserInvite() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	var h requests.UpdateUserInviteRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &h)

	if ui, err := models.GetUserInvitesById(id); err == nil {
		statuses := [3]string{"PENDING", "ACCEPTED", "CANCELLED"}
		proceed := false
		for _, st := range statuses {
			if h.Status == st {
				proceed = true
			}
		}

		if proceed {
			ui.Status = h.Status

			if err := models.UpdateUserInvitesById(ui); err == nil {
				logs.Debug("User invite updated successfully")
				logs.Debug(ui.InvitedBy.FullName)
				c.Ctx.Output.SetStatus(200)
				var resp = responses.UserInviteResponseDTO{StatusCode: 200, UserInvite: ui, StatusDesc: "User invite successfully updated"}
				c.Data["json"] = resp
			} else {
				logs.Debug("Error updating user invite", err.Error())
				var resp = responses.UserInviteResponseDTO{StatusCode: 608, UserInvite: nil, StatusDesc: "Error updating user invite"}
				c.Data["json"] = resp
			}
		} else {
			logs.Debug("Error updating status")
			var resp = responses.UserInviteResponseDTO{StatusCode: 503, UserInvite: nil, StatusDesc: "Please enter a valid status"}
			c.Data["json"] = resp
		}
	} else {
		logs.Debug("Error getting user invite", err.Error())
		var resp = responses.UserInviteResponseDTO{StatusCode: 608, UserInvite: nil, StatusDesc: "Error updating user invite"}
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
			resp := responses.StringResponseDTO{StatusCode: 200, Value: "OK", StatusDesc: "OK"}
			c.Data["json"] = resp
		} else {
			logs.Error("Error deleting user", err.Error())
			resp := responses.StringResponseDTO{StatusCode: 507, Value: "", StatusDesc: "ERROR " + err.Error()}
			c.Data["json"] = resp
		}
	} else {
		logs.Error("Error deleting user", err.Error())
		resp := responses.StringResponseDTO{StatusCode: 507, Value: "", StatusDesc: "ERROR " + err.Error()}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// Deactivate ...
// @Title Deactivate
// @Description delete the Users
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /deactivate/:id [delete]
func (c *UsersController) Deactivate() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)

	logs.Info("Checking user ", idStr)

	v, err := models.GetUsersById(id)
	if err == nil {
		v.Active = 4
		if err := models.UpdateUsersById(v); err == nil {
			// if err := models.DeleteUsers(id); err == nil {
			// 	c.Data["json"] = "OK"
			// } else {
			// 	c.Data["json"] = err.Error()
			// }
			resp := responses.StringResponseDTO{StatusCode: 200, Value: "OK", StatusDesc: "OK"}
			c.Data["json"] = resp
		} else {
			logs.Error("Error deactivating user", err.Error())
			resp := responses.StringResponseDTO{StatusCode: 507, Value: "", StatusDesc: "ERROR " + err.Error()}
			c.Data["json"] = resp
		}
	} else {
		logs.Error("Error deactivating user", err.Error())
		resp := responses.StringResponseDTO{StatusCode: 507, Value: "", StatusDesc: "ERROR " + err.Error()}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// GetItemCount ...
// @Title Get Item Quantity
// @Description get Item_quantity by Item id
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	search	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Success 200 {object} responses.StringResponseDTO
// @Failure 403 :id is empty
// @router /count/ [get]
func (c *UsersController) GetUserCount() {
	// q, err := models.GetItemsById(id)
	var query = make(map[string]string)
	var search = make(map[string]string)

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

	// search: k:v,k:v
	if v := c.GetString("search"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid search key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			search[k] = v
		}
	}

	v, err := models.GetUserCount(query, search)
	count := strconv.FormatInt(v, 10)
	if err != nil {
		logs.Error("Error fetching count of items ... ", err.Error())
		resp := responses.StringResponseDTO{StatusCode: 301, Value: "", StatusDesc: err.Error()}
		c.Data["json"] = resp
	} else {
		resp := responses.StringResponseDTO{StatusCode: 200, Value: count, StatusDesc: "Count fetched successfully"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}
