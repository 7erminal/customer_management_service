package controllers

import (
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

// ShopsController operations for Shops
type ShopsController struct {
	beego.Controller
}

// URLMapping ...
func (c *ShopsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Shops
// @Param	body		body 	requests.ShopRequest	true		"body for Shops content"
// @Success 201 {int} models.Shops
// @Failure 403 body is empty
// @router / [post]
func (c *ShopsController) Post() {
	var req requests.ShopRequest
	statusCode := 200
	message := ""
	resp := responses.ShopResp{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	v := models.Shops{
		ShopName:            req.Name,
		ShopLocation:        req.Location,
		ShopDescription:     req.Description,
		PhoneNumber:         req.PhoneNumber,
		Email:               req.Email,
		Image:               req.ImageUrl,
		ShopAssistantName:   req.AssistantName,
		ShopAssistantNumber: req.AssistantNumber,
	}
	if _, err := models.AddShops(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		message = "Created"
		resp = responses.ShopResp{
			ShopId:              strconv.FormatInt(v.ShopId, 10),
			ShopName:            v.ShopName,
			ShopDescription:     v.ShopDescription,
			ShopAssistantName:   v.ShopAssistantName,
			ShopAssistantNumber: v.ShopAssistantNumber,
			PhoneNumber:         v.PhoneNumber,
			Email:               v.Email,
			Image:               v.Image,
			ShopLocation:        v.ShopLocation,
			DateCreated:         v.DateCreated,
			DateModified:        v.DateModified,
			CreatedBy:           v.CreatedBy,
			ModifiedBy:          v.ModifiedBy,
			Active:              v.Active,
		}
		apiResp := responses.ShopResponse{
			StatusCode: statusCode,
			StatusDesc: message,
			Result:     resp,
		}
		c.Data["json"] = apiResp
	} else {
		statusCode = 500
		message = "Internal Server Error"
		apiResp := responses.ShopResponse{
			StatusCode: statusCode,
			StatusDesc: message,
			Result:     resp,
		}
		c.Data["json"] = apiResp
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Shops by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Shops
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ShopsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	statusCode := 200
	message := ""
	resp := responses.ShopResp{}
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetShopsById(id)
	if err != nil {
		statusCode = 404
		message = "Not Found"
		apiResp := responses.ShopResponse{
			StatusCode: statusCode,
			StatusDesc: message,
			Result:     resp,
		}
		c.Data["json"] = apiResp
	} else {
		resp = responses.ShopResp{
			ShopId:              strconv.FormatInt(v.ShopId, 10),
			ShopName:            v.ShopName,
			ShopDescription:     v.ShopDescription,
			ShopAssistantName:   v.ShopAssistantName,
			ShopAssistantNumber: v.ShopAssistantNumber,
			PhoneNumber:         v.PhoneNumber,
			Email:               v.Email,
			Image:               v.Image,
			ShopLocation:        v.ShopLocation,
			DateCreated:         v.DateCreated,
			DateModified:        v.DateModified,
			CreatedBy:           v.CreatedBy,
			ModifiedBy:          v.ModifiedBy,
			Active:              v.Active,
		}
		apiResp := responses.ShopResponse{
			StatusCode: 200,
			StatusDesc: "OK",
			Result:     resp,
		}
		c.Data["json"] = apiResp
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Shops
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Shops
// @Failure 403
// @router / [get]
func (c *ShopsController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	statusCode := 200
	message := "OK"

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

	l, err := models.GetAllShops(query, fields, sortby, order, offset, limit)
	if err != nil {
		statusCode = 500
		message = "Internal Server Error"
		apiResp := responses.ShopsResponse{
			StatusCode: statusCode,
			StatusDesc: message,
			Result:     nil,
		}
		c.Data["json"] = apiResp
	} else {
		respList := make([]responses.ShopResp, len(l))
		for qIdx, qItem := range l {
			respList[qIdx] = responses.ShopResp{
				ShopId:              strconv.FormatInt(qItem.(models.Shops).ShopId, 10),
				ShopName:            qItem.(models.Shops).ShopName,
				ShopDescription:     qItem.(models.Shops).ShopDescription,
				ShopAssistantName:   qItem.(models.Shops).ShopAssistantName,
				ShopAssistantNumber: qItem.(models.Shops).ShopAssistantNumber,
				PhoneNumber:         qItem.(models.Shops).PhoneNumber,
				Email:               qItem.(models.Shops).Email,
				Image:               qItem.(models.Shops).Image,
				ShopLocation:        qItem.(models.Shops).ShopLocation,
				DateCreated:         qItem.(models.Shops).DateCreated,
				DateModified:        qItem.(models.Shops).DateModified,
				CreatedBy:           qItem.(models.Shops).CreatedBy,
				ModifiedBy:          qItem.(models.Shops).ModifiedBy,
				Active:              qItem.(models.Shops).Active,
			}
		}
		apiResp := responses.ShopsResponse{
			StatusCode: statusCode,
			StatusDesc: message,
			Result:     respList,
		}
		c.Data["json"] = apiResp
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Shops
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Shops	true		"body for Shops content"
// @Success 200 {object} models.Shops
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ShopsController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	req := requests.UpdateShopRequest{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)

	statusCode := 200
	message := "OK"
	v := models.Shops{
		ShopId:              id,
		ShopName:            req.Name,
		PhoneNumber:         req.PhoneNumber,
		Email:               req.Email,
		Image:               req.ImageUrl,
		ShopDescription:     req.Description,
		ShopLocation:        req.Location,
		ShopAssistantName:   req.AssistantName,
		ShopAssistantNumber: req.AssistantNumber,
	}
	if err := models.UpdateShopsById(&v); err == nil {
		apiResp := responses.ShopResponse{
			StatusCode: statusCode,
			StatusDesc: message,
			Result: responses.ShopResp{
				ShopId:              strconv.FormatInt(v.ShopId, 10),
				ShopName:            v.ShopName,
				ShopDescription:     v.ShopDescription,
				ShopAssistantName:   v.ShopAssistantName,
				ShopAssistantNumber: v.ShopAssistantNumber,
				PhoneNumber:         v.PhoneNumber,
				Email:               v.Email,
				Image:               v.Image,
				ShopLocation:        v.ShopLocation,
				DateCreated:         v.DateCreated,
				DateModified:        v.DateModified,
				CreatedBy:           v.CreatedBy,
				ModifiedBy:          v.ModifiedBy,
				Active:              v.Active,
			},
		}
		c.Data["json"] = apiResp
		// c.Data["json"] = "OK"
	} else {
		logs.Error("Error updating shop ", err.Error())
		statusCode = 500
		message = err.Error()
		c.Data["json"] = responses.ShopResponse{
			StatusCode: statusCode,
			StatusDesc: message,
			Result:     responses.ShopResp{},
		}
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Shops
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ShopsController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteShops(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
