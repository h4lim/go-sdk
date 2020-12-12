package helpers

import (
	"github.com/h4lim/go-sdk/app/models"
	"github.com/h4lim/go-sdk/app/services"
	"github.com/h4lim/go-sdk/utils"
)

func SetClientApiModel(clientParty services.ClientParty, clientResponse services.ClientResponse) models.ClientApi {

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
