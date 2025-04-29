/*
This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.

SPDX-License-Identifier: MPL-2.0

File: internal.go
Description: Internal functions
Author: tengzl33t
*/

package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strings"
	"sync"
)

func getURL(apiEnv string, apiVer uint8, methodName string) string {
	return fmt.Sprintf("%s%s/%s", getApiEnvHost(apiEnv), getApiVerLink(apiVer), methodName)
}

func SendRequests(
	endpoint string, apiEnv string, headers map[string]string,
	payloads []map[string]interface{}, token string, apiKey string,
) []ResponseStruct {
	endpointEntry, ok := getEndpointMap()[endpoint]
	if !ok {
		panic(fmt.Sprintf("Endpoint '%s' not found", endpoint))
	}

	methodName := strings.ReplaceAll(endpoint, "v2", "")

	for _, payload := range payloads {
		if !slices.Contains(endpointEntry[1].([]string), payload["command"].(string)) {
			panic(fmt.Sprintf("Incorrect command '%s' found in method '%s'", payload["command"].(string), methodName))
		}
	}

	return processRequests(getURL(apiEnv, uint8(endpointEntry[0].(int)), methodName), headers, payloads, token, apiEnv, apiKey)
}

func Login(apiEnv string, apiKey string) string {
	url := fmt.Sprintf("%s%s/login", getApiEnvHost(apiEnv), getApiVerLink(1))

	jsonValue := map[string]interface{}{"command": "login", "api_token": apiKey}

	var jsonValues []map[string]interface{}
	jsonValues = append(jsonValues, jsonValue)

	response := processRequests(url, nil, jsonValues, "", "", "")[0]
	loginStatus := response.body.(map[string]interface{})["status"].(bool)
	if !loginStatus {
		panic("Could not login with key provided")
	}

	return response.body.(map[string]interface{})["token"].(string)
}

func processRequests(
	url string, headers map[string]string, payloads []map[string]interface{}, token string,
	apiEnv string, apiKey string,
) []ResponseStruct {
	httpClient := getHttpClient()

	semaphore := make(chan struct{}, 100)

	var responses []ResponseStruct
	responsesChan := make(chan ResponseStruct)
	var wg sync.WaitGroup

	mutexToken := MuToken{
		token: token,
	}

	for _, payload := range payloads {
		semaphore <- struct{}{}
		wg.Add(1)
		go processPostRequest(payload, &mutexToken, url, headers, httpClient, responsesChan, &wg, semaphore, apiEnv, apiKey)
	}

	go func() {
		wg.Wait()
		close(responsesChan)
	}()

	for respChanItem := range responsesChan {
		responses = append(responses, respChanItem)
	}

	return responses

}

func processPostRequest(
	payload map[string]interface{}, mutexToken *MuToken, url string, headers map[string]string,
	httpClient *http.Client, responseChan chan ResponseStruct, wg *sync.WaitGroup, semaphore chan struct{},
	apiEnv string, apiKey string,
) {
	mutexToken.mu.RLock()
	payload["token"] = mutexToken.token
	mutexToken.mu.RUnlock()

	preparedJsonPayload, _ := json.Marshal(payload)
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(preparedJsonPayload))
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "ThreatX-Go-API-Client")

	for key, value := range headers {
		request.Header.Set(key, value)
	}

	rawResponse, err := httpClient.Do(request)
	if err != nil {
		panic(err)
	}

	preparedBody, err := io.ReadAll(rawResponse.Body)
	if err != nil {
		panic(err)
	}

	var jsonData map[string]interface{}

	err = json.Unmarshal(preparedBody, &jsonData)
	if err != nil {
		// TODO: maybe validate if response is from TX API
		panic(fmt.Sprintf(
			"Could not parse the API response.\nError: %s\nRequest ID: %s.",
			err.Error(),
			rawResponse.Header["X-Request-Id"][0],
		))
	}

	payloadError, _ := jsonData["Error"].(string)
	if payloadError == "Token Expired. Please re-authenticate." {
		mutexToken.token = Login(apiEnv, apiKey)
		processPostRequest(payload, mutexToken, url, headers, httpClient, responseChan, wg, semaphore, apiEnv, apiKey)
		return
	}

	payloadMarker, _ := payload["marker_var"].(string)

	responseChan <- ResponseStruct{
		valid:      rawResponse.StatusCode == 200,
		body:       jsonData["Ok"],
		statusCode: rawResponse.StatusCode,
		marker:     payloadMarker,
	}

	wg.Done()
	<-semaphore
}
