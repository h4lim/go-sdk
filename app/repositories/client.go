package repositories

import (
	"fmt"
	"github.com/h4lim/go-sdk/app/models"
	"github.com/h4lim/go-sdk/utils"
)

func InsertClientApi(data models.ClientApi)  {

	db, err := utils.DBModel.DBOpen()
	defer db.Close()

	if err != nil {
		fmt.Println(data.ServiceName, "Error occurred ", *err)
	}

	db.Create(&data)
}
