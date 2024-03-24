package client

import (
	"encoding/json"
	"errors"
	goerr "github.com/darkjinnee/go-err"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type ResponseJson struct {
	Message   string        `json:"message"`
	Data      interface{}   `json:"data"`
	Errors    []interface{} `json:"errors"`
	Exception string        `json:"exception"`
}

func Request(
	url string,
	method string,
	data interface{},
) (*ResponseJson, error) {
	dataJson, _ := json.Marshal(data)
	payload := strings.NewReader(string(dataJson))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	goerr.Log(
		err,
		"[Error] notifierbot.Request: Failed creating request",
	)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	goerr.Log(
		err,
		"[Error] certbot.Request: Failed request",
	)

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	goerr.Log(
		err,
		"[Error] client.Request: Failed to read response body",
	)

	response := new(ResponseJson)
	err = json.Unmarshal(body, &response)
	goerr.Log(
		err,
		"[Error] client.Request: Failed convert body",
	)

	if response.Errors != nil {
		for key, err := range response.Errors {
			log.Printf("[Error] client.ErrorOut: %s > %s", key, err)
		}
		return response, errors.New(response.Message)
	}
	if response.Exception != "" {
		return response, errors.New(response.Message)
	}

	return response, nil
}
