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
	c.Mapping("GetAllBranches", c.GetAllBranches)
	c.Mapping("UploadShopLogo", c.UploadShopLogo)
}

// Post ...
// @Title Post
// @Description create Shops
// @Param	body		body 	models.Shops	true		"body for Shops content"
// @Success 201 {int} models.Shops
// @Failure 403 body is empty
// @router / [post]
func (c *ShopsController) Post() {
	var v models.Shops
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if _, err := models.AddShops(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = v
	} else {
		c.Data["json"] = err.Error()
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
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetShopsById(id)
	if err != nil {
		logs.Error("Error fetching shop details ", err.Error())
		resp := models.ShopResponseDTO{StatusCode: 608, Shop: nil, StatusDesc: "Error fetching shop details " + err.Error()}
		c.Data["json"] = resp
	} else {
		resp := models.ShopResponseDTO{StatusCode: 200, Shop: v, StatusDesc: "Successfully fetched shop details "}
		c.Data["json"] = resp
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
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// GetAllBranches ...
// @Title Get All Branches
// @Description get Branches
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} responses.BranchesResponseDTO
// @Failure 403
// @router /branches [get]
func (c *ShopsController) GetAllBranches() {
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

	l, err := models.GetAllBranches(query, fields, sortby, order, offset, limit)
	if err != nil {
		var resp = responses.BranchesResponseDTO{StatusCode: 606, Branches: nil, StatusDesc: "Error fetching users" + err.Error()}
		c.Data["json"] = resp
	} else {
		var resp = responses.BranchesResponseDTO{StatusCode: 200, Branches: &l, StatusDesc: "Branches fetched"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Shops
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	requests.ShopRequest	true		"body for Shops content"
// @Success 200 {object} models.Shops
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ShopsController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := requests.ShopRequest{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	logs.Info("Request received ", v)
	name := c.Ctx.Input.Query("Name")
	email := c.Ctx.Input.Query("Email")
	imageUrl := c.Ctx.Input.Query("ImageUrl")
	phonenumber := c.Ctx.Input.Query("PhoneNumber")
	assistantName := c.Ctx.Input.Query("AssistantName")
	assistantNumber := c.Ctx.Input.Query("AssistantNumber")
	logs.Info("Name is ", name, " Email ", email, " Image ", imageUrl, " and phone number ", phonenumber)

	if shop, err := models.GetShopsById(id); err == nil {
		if imageUrl == "" {
			imageUrl = shop.Image
		}
		q := models.Shops{ShopId: id, ShopName: name, Email: email, Image: imageUrl, PhoneNumber: phonenumber, ShopAssistantName: assistantName, ShopAssistantNumber: assistantNumber, DateCreated: time.Now(), DateModified: time.Now(), Active: 1}
		if err := models.UpdateShopsById(&q); err == nil {
			// c.Data["json"] = "OK"
			resp := models.ShopResponseDTO{StatusCode: 200, Shop: &q, StatusDesc: "Successfully updated shop details "}
			c.Data["json"] = resp
		} else {
			c.Data["json"] = err.Error()
			logs.Error("Error updating shop details ", err.Error())
			resp := models.ShopResponseDTO{StatusCode: 608, Shop: nil, StatusDesc: "Error updating shop details " + err.Error()}
			c.Data["json"] = resp
		}
	} else {
		c.Data["json"] = err.Error()
		logs.Error("Error updating shop details ", err.Error())
		resp := models.ShopResponseDTO{StatusCode: 608, Shop: nil, StatusDesc: "Error updating shop details " + err.Error()}
		c.Data["json"] = resp
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

// Upload Shop Logo ...
// @Title Upload Shop Logo
// @Description Upload shop's logo
// @Param	Image		formData 	file	true		"Item Image"
// @Success 200 {int} responses.StringResponseFDTO
// @Failure 403 body is empty
// @router /upload-shop-logo [post]
func (c *ShopsController) UploadShopLogo() {
	// var v models.Item_images
	file, header, err := c.GetFile("Image")
	logs.Info("Data received is ", file)

	contentType := c.Ctx.Input.Header("Content-Type")
	logs.Info("Content-Type of incoming request:", contentType)

	if err != nil {
		// c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": "Failed to get image file."}
		logs.Error("Failed to get the file ", err)
		c.ServeJSON()
		return
	}
	defer file.Close()

	// Check the file size
	fileInfo := header.Size
	logs.Info("Received file size:", fileInfo)

	// Save the uploaded file
	fileName := filepath.Base(header.Filename)
	logs.Info("File Name Extracted is ", fileName, "Time now is ", time.Now().Format("20060102150405"))
	filePath := "/uploads/shops/" + time.Now().Format("20060102150405") + fileName // Define your file path
	logs.Info("File Path Extracted is ", filePath)
	err = c.SaveToFile("Image", "../images"+filePath)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		logs.Error("Error saving file", err)
		// c.Data["json"] = map[string]string{"error": "Failed to save the image file."}
		errorMessage := "Error: Failed to save the image file"

		resp := responses.StringResponseDTO{StatusCode: http.StatusInternalServerError, Value: errorMessage, StatusDesc: "Internal Server Error"}

		c.Data["json"] = resp
		c.ServeJSON()
		return
	} else {
		resp := responses.StringResponseDTO{StatusCode: 200, Value: filePath, StatusDesc: "Images uploaded successfully"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}
