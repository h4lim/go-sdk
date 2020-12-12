package services

import (
	"fmt"
	"github.com/h4lim/go-sdk/app/models"
	"github.com/h4lim/go-sdk/utils"
)

func CountClientApi(clientParty ClientParty, clientResponse ClientResponse) {
	go insertClientApi(setClientApiModel(clientParty, clientResponse))
}

func insertClientApi(data models.ClientApi) {

	db, err := utils.DBModel.DBOpen()
	defer db.Close()

	if err != nil {
		fmt.Println(data.ClientName, "Error occurred ", *err)
	}

	db.Create(&data)
}

func setClientApiModel(clientParty ClientParty, clientResponse ClientResponse) models.ClientApi {

	models := models.ClientApi{
		Environment:  utils.GetRunMode(),
		ClientName:   clientParty.ClientName,
		Url:          clientParty.UrlApi.String(),
		RequestBody:  clientParty.RequestBody,
		ResponseBody: string(clientResponse.ByteResponse),
		HttpCode:     clientResponse.HttpCode,
	}

	return models
}
