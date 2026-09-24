package controllers

import (
	"customer_management_service/controllers/functions"
	"customer_management_service/models"
	"customer_management_service/structs/requests"
	"customer_management_service/structs/responses"
	"encoding/json"
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
	var branchRequestDTO requests.BranchRequestDTO
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &branchRequestDTO); err != nil {
		logs.Error("Error parsing request body: ", err.Error())
		var resp = responses.BranchResponseDTO{StatusCode: 400, Result: nil, StatusDesc: "Error parsing request body: " + err.Error()}
		c.Data["json"] = resp
		c.ServeJSON()
		return
	}

	if country, err := functions.GetCountryWithCode(&c.Controller, branchRequestDTO.CountryCode); err != nil {
		logs.Error("Error fetching country details for ", branchRequestDTO.CountryCode, " is ", err.Error())
		var resp = responses.BranchResponseDTO{StatusCode: 400, Result: nil, StatusDesc: "Error fetching country details for " + branchRequestDTO.CountryCode + " is " + err.Error()}
		c.Data["json"] = resp
		c.ServeJSON()
		return
	} else {

		branchActive := 0
		addedByInt, _ := strconv.ParseInt(branchRequestDTO.AddedBy, 10, 64)
		if branchRequestDTO.Active == "true" {
			branchActive = 1
		} else {
			branchActive = 0
		}
		branchModel := models.Branches{
			Branch:      branchRequestDTO.Branch,
			Country:     country.Result.CountryId,
			PhoneNumber: branchRequestDTO.PhoneNumber,
			Location:    branchRequestDTO.Location,
			Active:      branchActive,
			CreatedBy:   int(addedByInt),
			ModifiedBy:  int(addedByInt),
		}

		// Call the model function to create the branch
		_, err := models.AddBranches(&branchModel)
		if err != nil {
			logs.Error("Error creating branch: ", err.Error())
			var resp = responses.BranchResponseDTO{StatusCode: 500, Result: nil, StatusDesc: "Error creating branch: " + err.Error()}
			c.Data["json"] = resp
		} else {
			logs.Info("Successfully created branch with ID ", branchModel.BranchId)
			countryResp := responses.CountryResp{
				CountryId:   country.Result.CountryId,
				Country:     country.Result.Country,
				CountryCode: country.Result.CountryCode,
				Currency:    nil,
			}
			branchResp := responses.BranchResp{
				BranchId:     branchModel.BranchId,
				BranchName:   branchModel.Branch,
				Country:      &countryResp,
				Location:     branchModel.Location,
				CreatedBy:    branchModel.CreatedBy,
				ModifiedBy:   branchModel.ModifiedBy,
				DateCreated:  branchModel.DateCreated,
				DateModified: branchModel.DateModified,
			}
			var resp = responses.BranchResponseDTO{StatusCode: 200, Result: &branchResp, StatusDesc: "Successfully created branch"}
			c.Data["json"] = resp
		}
	}
	c.ServeJSON()
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
	v, err := models.GetBranchesById(id)
	if err != nil {
		logs.Error("Error fetching branch details for ", id, " is ", err.Error())
		var resp = responses.BranchResponseDTO{StatusCode: 301, Result: nil, StatusDesc: "Error fetching branch details for " + strconv.FormatInt(id, 10) + " is " + err.Error()}
		c.Data["json"] = resp
	} else {
		logs.Info("Successfully fetched branch details for ", id)

		// Log branch details in json
		branchDetails, _ := json.Marshal(v)
		logs.Info("Branch details are ", string(branchDetails))

		countryM := responses.CountryResp{}
		currencyM := responses.CurrencyResp{}
		countryResp, err := functions.GetCountryWithId(&c.Controller, strconv.FormatInt(v.Country, 10))
		if err != nil {
			logs.Error("Error fetching country details for ", v.Country, " is ", err.Error())
			var resp = responses.BranchResponseDTO{StatusCode: 301, Result: nil, StatusDesc: "Error fetching country details for " + strconv.FormatInt(v.Country, 10) + " is " + err.Error()}
			c.Data["json"] = resp
		} else {
			logs.Info("Successfully fetched country details for ", v.Country)
			currencyM = responses.CurrencyResp{
				CurrencyId: countryResp.Result.DefaultCurrency.CurrencyId,
				Currency:   countryResp.Result.DefaultCurrency.Currency,
				Symbol:     countryResp.Result.DefaultCurrency.Symbol,
			}
			countryM = responses.CountryResp{
				CountryId:   countryResp.Result.CountryId,
				Country:     countryResp.Result.Country,
				CountryCode: countryResp.Result.CountryCode,
				Currency:    &currencyM,
			}

		}
		branchResp := responses.BranchResp{
			BranchId:     v.BranchId,
			BranchName:   v.Branch,
			Location:     v.Location,
			Country:      &countryM,
			Active:       v.Active,
			DateCreated:  v.DateCreated,
			DateModified: v.DateModified,
			CreatedBy:    v.CreatedBy,
			ModifiedBy:   v.ModifiedBy,
		}
		var resp = responses.BranchResponseDTO{StatusCode: 200, Result: &branchResp, StatusDesc: "Successfully fetched branch details for " + strconv.FormatInt(id, 10)}
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

	statusCode := 404
	logs.Info("Fetching all branches")
	message := "Fetching all branches"

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
	query_ := c.GetString("query")
	logs.Info("Trimming space for ", query_)
	query_ = strings.TrimSpace(query_) + ",Active:1"
	logs.Info("Trimmed query string: ", query_)
	if v := query_; v != "" {
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

	logs.Info("About to run query with parameters: ", query)
	l, err := models.GetAllBranches(query, fields, sortby, order, offset, limit)
	logs.Info("Query executed, checking for errors")
	if err != nil {
		statusCode = 301
		message = "Error fetching branch details"
		resp := responses.BranchesResponseDTO{StatusCode: statusCode, Result: nil, StatusDesc: message}
		c.Data["json"] = resp
	} else {
		logs.Info("Number of branches fetched: ", len(l))
		branchesResp := []responses.BranchResp{}
		for _, br := range l {
			m := br.(models.Branches)

			branchesResp = append(branchesResp, responses.BranchResp{
				BranchId:   m.BranchId,
				BranchName: m.Branch,
				Location:   m.Location,
				// Country:      &countryM,
				Active:       m.Active,
				DateCreated:  m.DateCreated,
				DateModified: m.DateModified,
				CreatedBy:    m.CreatedBy,
				ModifiedBy:   m.ModifiedBy,
			})
		}
		statusCode = 200
		message = "Successfully fetched branches"
		resp := responses.BranchesResponseDTO{StatusCode: statusCode, Result: &branchesResp, StatusDesc: message}
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
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if branch, err := models.GetBranchesById(id); err == nil {
		branch.Active = 6
		if err := models.UpdateBranchesById(branch); err == nil {
			statusCode := 200
			message := "OK"
			logs.Info("Successfully updated branch with ID ", branch.BranchId)
			branchResp := responses.BranchResp{
				BranchId:     branch.BranchId,
				BranchName:   branch.Branch,
				Location:     branch.Location,
				Active:       branch.Active,
				DateCreated:  branch.DateCreated,
				DateModified: branch.DateModified,
				CreatedBy:    branch.CreatedBy,
				ModifiedBy:   branch.ModifiedBy,
			}
			resp := responses.BranchResponseDTO{
				StatusCode: statusCode,
				StatusDesc: message,
				Result:     &branchResp,
			}
			c.Data["json"] = resp
		} else {
			resp := responses.BranchResponseDTO{
				StatusCode: 500,
				StatusDesc: err.Error(),
				Result:     &responses.BranchResp{},
			}
			c.Data["json"] = resp
		}
	} else {
		resp := responses.BranchResponseDTO{
			StatusCode: 500,
			StatusDesc: err.Error(),
			Result:     &responses.BranchResp{},
		}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}
