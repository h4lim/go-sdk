package logging

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var GlobalLogMessage *TelegramLog

type TelegramLog struct {
	ServiceName string
	BotApiToken string
	GroupChatID string
	Debug       bool
	Environment string
}

type TelegramResponse struct {
	HttpCode     int
	ByteResponse []byte
}

type LogFunction interface {
	SendMessage(message string)
}

func (c *TelegramLog) SendMessage(message string) {

	if _, err := messageService(c, message); err != nil {
		fmt.Println(c.ServiceName, "Error occurred ", *err)
	}

}

func messageService(c *TelegramLog, message string) (*TelegramResponse, *error) {

	urlApi := setApiUrl(c.ServiceName,
		"https://api.telegram.org/bot"+c.BotApiToken+"/sendMessage")

	rawQuery := urlApi.Query()
	rawQuery.Set("chat_id", c.GroupChatID)
	rawQuery.Set("text", message)
	urlApi.RawQuery = rawQuery.Encode()

	request, err := http.NewRequest(http.MethodGet, urlApi.String(), nil)
	if err != nil {
		fmt.Println(c.ServiceName, "Error occurred %s ", err.Error())
		return nil, &err
	}

	request.Header.Set("Content-Type", "application/json")

	if c.Debug {
		fmt.Println(c.ServiceName, "===================================================== HIT TELEGRAM API ===================================================")
		fmt.Println(c.ServiceName, "HTTP METHOD %s ", request.Method)
		fmt.Println(c.ServiceName, "URL %s ", request.URL)
		fmt.Println(c.ServiceName, "HEADER %s ", request.Header)
		fmt.Println(c.ServiceName, "==========================================================================================================================")
	}

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(c.ServiceName, "Error occurred %s ", err.Error())
		return nil, &err
	}

	byteResult, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(c.ServiceName, "Error occurred %s ", err.Error())
		return nil, &err
	}

	telegramResponse := TelegramResponse{
		HttpCode:     response.StatusCode,
		ByteResponse: byteResult,
	}

	if c.Debug {
		fmt.Println(c.ServiceName, "================================================ RESPONSE "+c.ServiceName+" API ======================================")
		fmt.Println(c.ServiceName, "GET HTTP RESPONSE CODE ", telegramResponse.HttpCode)
		fmt.Println(c.ServiceName, "GET RESPONSE FROM "+strings.ToUpper(c.ServiceName)+" API %s", string(telegramResponse.ByteResponse))
		fmt.Println(c.ServiceName, "======================================================================================================================")

	}

	return &telegramResponse, nil
}

func setApiUrl(uniqueID string, urlApi string) *url.URL {
	url, err := url.Parse(urlApi)
	if err != nil {
		fmt.Println(uniqueID, "Error occurred %s ", err.Error())
		return nil
	}
	return url
}
