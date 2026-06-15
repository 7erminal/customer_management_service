package functions

import (
	"bytes"
	"customer_management_service/api"
	"customer_management_service/structs/responses"
	"encoding/json"
	"io"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// containsIgnoreCase checks if substr is in s, case-insensitive.
func containsIgnoreCase(s, substr string) bool {
	return len(s) >= len(substr) &&
		strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

func setControllerJSON(c *beego.Controller, value interface{}) {
	if c == nil {
		return
	}
	if c.Data == nil {
		c.Data = map[interface{}]interface{}{}
	}
	c.Data["json"] = value
}

func GetCountryWithId(c *beego.Controller, countryId string) (responses.CountryResponseDTO, error) {
	host, _ := beego.AppConfig.String("systemBaseUrl")

	logs.Info("Request to get Country: ", countryId)

	request := api.NewRequest(
		host,
		"/v1/countries/"+countryId,
		api.GET)

	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		setControllerJSON(c, err.Error())
		return responses.CountryResponseDTO{}, err
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		setControllerJSON(c, err.Error())
		return responses.CountryResponseDTO{}, err
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	// var dataOri responses.UserOriResponseDTO
	var data responses.CountryResponseDTO
	if err := json.Unmarshal(read, &data); err != nil {
		setControllerJSON(c, err.Error())
		return responses.CountryResponseDTO{}, err
	}
	setControllerJSON(c, data)

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User.Branch.Country.DefaultCurrency)

	return data, nil
}

func GetCountryWithCode(c *beego.Controller, countryCode string) (responses.CountryResponseDTO, error) {
	host, _ := beego.AppConfig.String("systemBaseUrl")

	logs.Info("Request to get Country: ", countryCode)

	request := api.NewRequest(
		host,
		"/v1/countries/code/"+countryCode,
		api.GET)

	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		setControllerJSON(c, err.Error())
		return responses.CountryResponseDTO{}, err
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		setControllerJSON(c, err.Error())
		return responses.CountryResponseDTO{}, err
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	// var dataOri responses.UserOriResponseDTO
	var data responses.CountryResponseDTO
	if err := json.Unmarshal(read, &data); err != nil {
		setControllerJSON(c, err.Error())
		return responses.CountryResponseDTO{}, err
	}
	setControllerJSON(c, data)

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User.Branch.Country.DefaultCurrency)

	return data, nil
}

func GetCurrencyWithName(c *beego.Controller, currencyName string) (responses.CurrencyResponseDTO, error) {
	host, _ := beego.AppConfig.String("systemBaseUrl")

	logs.Info("Request to get Currency: ", currencyName)

	request := api.NewRequest(
		host,
		"/v1/currencies/"+currencyName,
		api.GET)

	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		setControllerJSON(c, err.Error())
		return responses.CurrencyResponseDTO{}, err
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		setControllerJSON(c, err.Error())
		return responses.CurrencyResponseDTO{}, err
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	// var dataOri responses.UserOriResponseDTO
	var data responses.CurrencyResponseDTO
	if err := json.Unmarshal(read, &data); err != nil {
		setControllerJSON(c, err.Error())
		return responses.CurrencyResponseDTO{}, err
	}
	setControllerJSON(c, data)

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User.Branch.Country.DefaultCurrency)

	return data, nil
}
