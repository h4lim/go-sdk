package services

import (
	"fmt"
	"github.com/h4lim/go-sdk/logging"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var log = logging.MustGetLogger("go-sdk")

const (
	INTERVAL_MINUTE = "MINUTE"
	INTERVAL_SECOND = "SECOND"
)

type ThirdParty interface {
	HitClient(url url.URL) (*ClientResponse, *error)
}

type ClientParty struct {
	UniqueID    string
	ClientName  string
	HttpMethod  string
	Debug       bool
	UrlApi      url.URL
	HttpClient  http.Client
	Headers     []map[string]string
	RequestBody io.Reader
	LogApi      bool
	ClientRetry *ClientRetry
}

type ClientRetry struct {
	MaxRetry      int
	HttpToRetry   *[]map[int]string
	Interval      string
	BeginInterval int
	EndInterval   int
}

type ClientResponse struct {
	HttpCode     int
	ByteResponse []byte
}

func (c *ClientParty) HitClient() (*ClientResponse, *error) {

	request, err := http.NewRequest(c.HttpMethod, c.UrlApi.String(), c.RequestBody)
	if err != nil {
		log.Errorf(c.UniqueID, "Error occurred %s ", err.Error())
		return nil, &err
	}

	requestBody := fmt.Sprintf("%v", c.RequestBody)

	for _, element := range c.Headers {
		for key, value := range element {
			if strings.ToLower(key) == "baseauth" {
				baseAuth := strings.Split(value, ":")
				request.SetBasicAuth(baseAuth[0], baseAuth[1])
			} else {
				request.Header.Set(key, value)
			}
		}
	}

	if c.Debug {
		log.Debugf(c.UniqueID, "===================================================== HIT "+strings.ToUpper(c.ClientName)+" API =======================================================")
		log.Debugf(c.UniqueID, "HTTP METHOD %s ", request.Method)
		log.Debugf(c.UniqueID, "URL %s ", request.URL)
		log.Debugf(c.UniqueID, "HEADER %s ", request.Header)
		if request.Body != nil {
			log.Debugf(c.UniqueID, "BODY %s ", request.Body)
		}
		log.Debugf(c.UniqueID, "==========================================================================================================================")
	}

	if c.ClientRetry != nil {
		clientResponse, err := httpRetry(c, request, requestBody)
		if err != nil {
			return nil, err
		}
		if clientResponse != nil {
			return clientResponse, nil
		}
	}

	response, err := c.HttpClient.Do(request)
	if err != nil {
		log.Errorf(c.UniqueID, "Error occurred %s ", err.Error())
		return nil, &err
	}

	byteResult, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Errorf(c.UniqueID, "Error occurred %s ", err.Error())
		return nil, &err
	}
	printClientResponse(c.Debug, c.UniqueID, response.StatusCode, c.ClientName, string(byteResult))

	clientResponse := ClientResponse{
		HttpCode:     response.StatusCode,
		ByteResponse: byteResult,
	}

	if c.LogApi {
		CountClientApi(*c, *request, requestBody, clientResponse)
	}

	return &clientResponse, nil
}

func httpRetry(c *ClientParty, request *http.Request, requestBody string) (*ClientResponse, *error) {

	var response *http.Response
	var errResponse error
	var byteResult []byte
	var byteError error

	interval := time.Duration(c.ClientRetry.BeginInterval)
	for i := 0; i < c.ClientRetry.MaxRetry; i++ {
		for _, element := range *c.ClientRetry.HttpToRetry {
			for key, value := range element {

				response, errResponse = c.HttpClient.Do(request)
				if errResponse != nil {
					log.Errorf(c.UniqueID, "Error occurred %s ", errResponse.Error())
					return nil, &errResponse
				}

				byteResult, byteError = ioutil.ReadAll(response.Body)
				if byteError != nil {
					log.Errorf(c.UniqueID, "Error occurred %s ", byteError.Error())
					return nil, &byteError
				}

				if response.StatusCode != key {
					printClientResponse(c.Debug, c.UniqueID, response.StatusCode, c.ClientName, string(byteResult))
					clientResponse := ClientResponse{
						HttpCode:     response.StatusCode,
						ByteResponse: byteResult,
					}

					if c.LogApi {
						CountClientApi(*c, *request, requestBody, clientResponse)
					}

					log.Debugf(c.UniqueID, "No retry from client")
					return &clientResponse, nil
				}

				log.Warningf(c.UniqueID, "Retry occurred when http code and message, retry for ", key, value, i+1)
				log.Warningf(c.UniqueID, "Interval - ", interval)
				printClientResponse(c.Debug, c.UniqueID, response.StatusCode, c.ClientName, string(byteResult))
				if c.ClientRetry.Interval == INTERVAL_MINUTE {
					time.Sleep(interval * time.Minute)
				} else {
					time.Sleep(interval * time.Second)
				}
				interval *= time.Duration(c.ClientRetry.EndInterval)
			}
		}

	}

	log.Debugf(c.UniqueID, "Retry finished")

	return nil, nil
}

func printClientResponse(debug bool, uniqueID string, responseCode int, clientName string, strResult string) {
	if debug {
		log.Debugf(uniqueID, "================================================ RESPONSE "+strings.ToUpper(clientName)+" API ================================================")
		log.Debugf(uniqueID, "GET HTTP RESPONSE CODE ", responseCode)
		log.Debugf(uniqueID, "GET RESPONSE FROM "+strings.ToUpper(clientName)+" API %s", strResult)
		log.Debugf(uniqueID, "======================================================================================================================")
	}
}
