package utils

import (
	"fmt"
	"github.com/h4lim/go-sdk/app/models"
	"github.com/h4lim/go-sdk/app/services"
)

func CountClientApi(clientParty services.ClientParty, clientResponse services.ClientResponse) {
	go insertClientApi(setClientApiModel(clientParty, clientResponse))
}

func insertClientApi(data models.ClientApi) {

	db, err := DBModel.DBOpen()
	defer db.Close()

	if err != nil {
		fmt.Println(data.ClientName, "Error occurred ", *err)
	}

	db.Create(&data)
}

func setClientApiModel(clientParty services.ClientParty, clientResponse services.ClientResponse) models.ClientApi {

	models := models.ClientApi{
		Environment:  GetRunMode(),
		ClientName:   clientParty.ClientName,
		Url:          clientParty.UrlApi.String(),
		RequestBody:  clientParty.RequestBody,
		ResponseBody: string(clientResponse.ByteResponse),
		HttpCode:     clientResponse.HttpCode,
	}

	return models
}
