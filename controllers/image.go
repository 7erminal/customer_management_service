package controllers

import (
	"net/http"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"

	"customer_management_service/structs/responses"
)

// ImageController operations for Image
type ImageController struct {
	beego.Controller
}

// URLMapping ...
func (c *ImageController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("UploadImage", c.UploadImage)
}

// UploadImage ...
// @Title Upload Image
// @Description upload an image for the application
// @Param	Image		formData 	file	true		"The image file to upload"
// @Param	ItemID		query 	string	true		"The ID of the item to associate the image with"
// @Success 200 {object} responses.SystemImageResponseDTO
// @Failure 400,500 {object} responses.SystemErrorResponse
// @router /upload-image [post]
func (c *ImageController) UploadImage() {
	// var v models.Item_images

	logs.Info("Data received is ", c.Ctx.Input.Query("ItemID"))

	file, header, err := c.GetFile("Image")
	system := strings.ToLower(c.Ctx.Input.Query("System"))

	logs.Info("Data received is ", file)
	logs.Info("System received is ", system)

	if err != nil {
		// c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": "Failed to get image file."}
		logs.Error("Failed to get the file ", err)
		c.ServeJSON()
		return
	}
	defer file.Close()

	// Save the uploaded file
	fileName := header.Filename
	logs.Info("File Name Extracted is ", fileName)
	filePath := "/uploads/" + system + "/" + fileName // Define your file path
	viewHost, _ := beego.AppConfig.String("imagesBaseUrl")
	viewFilePath := viewHost + filePath
	logs.Info("File Path Extracted is ", filePath)
	host, _ := beego.AppConfig.String("imagesUploadBaseUrl")
	filePath = host + filePath
	logs.Info("Full file path is ", filePath)

	err = c.SaveToFile("Image", filePath)

	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		logs.Error("Error saving file", err)
		// c.Data["json"] = map[string]string{"error": "Failed to save the image file."}
		errorMessage := "Error: Failed to save the image file"

		resp := responses.StringResponseDTO{StatusCode: http.StatusInternalServerError, Value: errorMessage, StatusDesc: "Internal Server Error"}

		c.Data["json"] = resp
		c.ServeJSON()
		return
	}

	response := responses.StringResponseDTO{StatusCode: 200, Value: viewFilePath, StatusDesc: "Image uploaded successfully"}
	c.Data["json"] = response
	c.ServeJSON()
}

// Post ...
// @Title Create
// @Description create Image
// @Param	body		body 	models.Image	true		"body for Image content"
// @Success 201 {object} models.Image
// @Failure 403 body is empty
// @router / [post]
func (c *ImageController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Image by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Image
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ImageController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Image
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Image
// @Failure 403
// @router / [get]
func (c *ImageController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Image
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Image	true		"body for Image content"
// @Success 200 {object} models.Image
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ImageController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Image
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ImageController) Delete() {

}
