package logging

import (
	"fmt"
	"github.com/h4lim/go-sdk/app/models"
	"github.com/h4lim/go-sdk/app/services"
	"github.com/h4lim/go-sdk/utils"
)

func (ppl *GoSDK) Success(logID string, clientParty services.ClientParty,
	clientResponse services.ClientResponse, args ...interface{}) {
	go insertClientApi(setClientApiModel(clientParty, clientResponse))
	args = append([]interface{}{"[" + ppl.Hostname + "] [" + logID + "]"}, args...)
	ppl.Logger.Notice("notice", args)
}

func (ppl *GoSDK) Successf(logID string, stringFormat string,
	clientParty services.ClientParty, clientResponse services.ClientResponse, args ...interface{}) {
	go insertClientApi(setClientApiModel(clientParty, clientResponse))
	ppl.Logger.Noticef("["+ppl.Hostname+"] ["+logID+"] "+stringFormat, args...)
}

func (ppl *GoSDK) Failed(logID string,
	clientParty services.ClientParty, clientResponse services.ClientResponse, args ...interface{}) {
	go insertClientApi(setClientApiModel(clientParty, clientResponse))
	args = append([]interface{}{"[" + ppl.Hostname + "] [" + logID + "]"}, args...)
	ppl.Logger.Notice("notice", args)
}

func (ppl *GoSDK) Failedf(logID string, stringFormat string,
	clientParty services.ClientParty, clientResponse services.ClientResponse, args ...interface{}) {
	go insertClientApi(setClientApiModel(clientParty, clientResponse))
	ppl.Logger.Noticef("["+ppl.Hostname+"] ["+logID+"] "+stringFormat, args...)
}

func insertClientApi(data models.ClientApi) {

	db, err := utils.DBModel.DBOpen()
	defer db.Close()

	if err != nil {
		fmt.Println(data.ClientName, "Error occurred ", *err)
	}

	db.Create(&data)
}

func setClientApiModel(clientParty services.ClientParty, clientResponse services.ClientResponse) models.ClientApi {

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
