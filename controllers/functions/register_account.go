package functions

import (
	"customer_management_service/api"
	"customer_management_service/structs/responses"
	"encoding/json"
	"io"
	"strconv"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func RegisterAccount(c *beego.Controller, userid int64) {
	host, err := beego.AppConfig.String("accountsBaseUrl")

	logs.Info("Sending user ID ", userid)

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
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	logs.Info("Raw response received is ", res)
	// data := map[string]interface{}{}
	var data responses.AccountDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	// logs.Info("Response received ", c.Data["json"])
	// logs.Info("Access token ", data["access_token"])
	// logs.Info("Expires in ", data["expires_in"])
	// logs.Info("Scope is ", data["scope"])
	// logs.Info("Token Type is ", data["token_type"])
	// logs.Info("Response received ", c.Data["json"])
	// logs.Info("Access token ", data.Access_token)
	// logs.Info("Expires in ", data.Expires_in)
	// logs.Info("Scope is ", data.Scope)
	// logs.Info("Token Type is ", data.Token_type)
}

func GenerateToken(c *beego.Controller, email string, role string) (resp responses.InviteHashResponseDTO) {
	host, _ := beego.AppConfig.String("authenticationBaseUrl")

	logs.Info("About to generate token ", email)

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
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	logs.Info("Raw response received is ", res)
	// data := map[string]interface{}{}
	var data responses.InviteHashResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	return data
}

func VerifyUserToken(c *beego.Controller, token string, nonce string, email string) (resp responses.StringResponseDTO) {
	host, _ := beego.AppConfig.String("authenticationBaseUrl")

	logs.Info("About to verify token ", token)

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
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	logs.Info("Raw response received is ", res)
	// data := map[string]interface{}{}
	var data responses.StringResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	return data
}
