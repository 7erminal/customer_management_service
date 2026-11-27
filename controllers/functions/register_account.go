package functions

import (
	"customer_management_service/api"
	"customer_management_service/logging"
	"customer_management_service/structs/responses"
	"encoding/json"
	"io"
	"strconv"

	beego "github.com/beego/beego/v2/server/web"
)

func RegisterAccount(c *beego.Controller, userid int64) {
	ctx := c.Ctx.Request.Context()
	log := logging.For(ctx)
	host, err := beego.AppConfig.String("accountsBaseUrl")

	log.Info("Sending user ID to accounts service", "user_id", userid)

	request := api.NewRequest(
		host,
		"/v1/accounts",
		api.POST)
	request.Params["UserId"] = strconv.Itoa(int(userid))
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequestWithContext(ctx)
	if err != nil {
		log.Error("accounts service request failed", "error", err.Error())
		c.Data["json"] = err.Error()
		return
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error("reading accounts service response failed", "error", err.Error())
		c.Data["json"] = err.Error()
		return
	}

	var data responses.AccountDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data
}

func GenerateToken(c *beego.Controller, email string, role string) (resp responses.InviteHashResponseDTO) {
	ctx := c.Ctx.Request.Context()
	log := logging.For(ctx)
	host, _ := beego.AppConfig.String("authenticationBaseUrl")

	log.Info("About to generate invite token")

	request := api.NewRequest(
		host,
		"/v1/auth/token/invite",
		api.POST)
	request.Params["Email"] = email
	request.Params["Role"] = role
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequestWithContext(ctx)
	if err != nil {
		log.Error("authentication service request failed", "error", err.Error())
		c.Data["json"] = err.Error()
		return
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error("reading authentication service response failed", "error", err.Error())
		c.Data["json"] = err.Error()
		return
	}

	var data responses.InviteHashResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	return data
}

func VerifyUserToken(c *beego.Controller, token string, nonce string, email string) (resp responses.InviteDecodeResponseDTO) {
	ctx := c.Ctx.Request.Context()
	log := logging.For(ctx)
	host, _ := beego.AppConfig.String("authenticationBaseUrl")

	log.Info("About to verify invite token")

	request := api.NewRequest(
		host,
		"/v1/auth/token/invite/verify",
		api.POST)
	request.Params["Token"] = token
	request.Params["Nonce"] = nonce
	request.Params["Email"] = email
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequestWithContext(ctx)
	if err != nil {
		log.Error("authentication service request failed", "error", err.Error())
		c.Data["json"] = err.Error()
		return
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error("reading authentication service response failed", "error", err.Error())
		c.Data["json"] = err.Error()
		return
	}

	var data responses.InviteDecodeResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	return data
}
