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

	models := models.LogApi{
		LogID:        clientParty.UniqueID,
		Environment:  "Dev",
		ClientName:   clientParty.ClientName,
		Url:          clientParty.UrlApi.String(),
		Method:       request.Method,
		Header:       fmt.Sprintf("%v", request.Header),
		RequestBody:  requestBody,
		ResponseBody: string(clientResponse.ByteResponse),
		HttpCode:     clientResponse.HttpCode,
	}

	return models
}
