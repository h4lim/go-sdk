package services

import (
	"fmt"
	"github.com/h4lim/go-sdk/app/models"
	"github.com/h4lim/go-sdk/utils"
	"net/http"
)

func CountClientApi(clientParty ClientParty, request http.Request, requestBody string, clientResponse ClientResponse) {
	go insertClientApi(setClientApiModel(clientParty, request, requestBody, clientResponse))
}

func insertClientApi(data models.LogApi) {

	db, err := utils.DBModel.DBOpen()
	defer db.Close()

	if err != nil {
		fmt.Println(data.ClientName, "Error occurred ", *err)
	}

	db.Create(&data)
}

func setClientApiModel(clientParty ClientParty, request http.Request, requestBody string, clientResponse ClientResponse) models.LogApi {

	url := clientParty.UrlApi.String()
	if clientParty.HiddenLog.Url {
		url = ""
	}

	header := fmt.Sprintf("%v", request.Header)
	if clientParty.HiddenLog.Header {
		header = ""
	}

	if clientParty.HiddenLog.RequestBody {
		requestBody = ""
	}

	responseBody := string(clientResponse.ByteResponse)
	if clientParty.HiddenLog.ResponseBody {
		responseBody = ""
	}

	models := models.LogApi{
		LogID:        clientParty.UniqueID,
		Environment:  utils.GetRunMode(),
		ClientName:   clientParty.ClientName,
		Url:          url,
		Method:       request.Method,
		Header:       header,
		RequestBody:  requestBody,
		ResponseBody: responseBody,
		HttpCode:     clientResponse.HttpCode,
	}

	return models
}
