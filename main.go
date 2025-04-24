package main

import "C"
import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"slices"
	"sync"
	"time"
)

type responseStruct struct {
	valid      bool
	body       interface{}
	statusCode int
	marker     string
}

type muToken struct {
	mu    sync.RWMutex
	token string
}

func getApiVerLink(version uint8) string {
	return fmt.Sprintf("/tx_api/v%d", version)
}

func getApiEnvHost(apiEnv string) string {
	domainPart := "threatx.io"
	oldHostParts := map[string]string{
		"prod":    "",
		"pod":     "tx-us-east-2a",
		"qa":      "qa0",
		"dev":     "dev0",
		"staging": "staging0",
	}
	if _, ok := oldHostParts[apiEnv]; ok {
		return fmt.Sprintf("https://api%s.%s", oldHostParts[apiEnv], domainPart)
	}

	return fmt.Sprintf("https://%s.%s", apiEnv, domainPart)
}

func sites(apiEnv string, headers map[string]string, payloads []map[string]interface{}, token string, apiKey string) []responseStruct {
	availableCommands := []string{
		"list",
		"new",
		"get",
		"delete",
		"update",
		"unset",
	}
	methodName := "sites"
	url := fmt.Sprintf("%s%s/%s", getApiEnvHost(apiEnv), getApiVerLink(2), methodName)
	return validateRequests(url, headers, payloads, availableCommands, token, methodName, apiEnv, apiKey)
}

func validateRequests(
	url string, headers map[string]string, payloads []map[string]interface{}, availableCommands []string,
	token string, caller string, apiEnv string, apiKey string,
) []responseStruct {
	for _, payload := range payloads {
		if !slices.Contains(availableCommands, payload["command"].(string)) {
			panic(fmt.Sprintf("Incorrect command '%s' found in method '%s'", payload["command"].(string), caller))
		}
	}
	return processRequests(url, headers, payloads, token, apiEnv, apiKey)
}

func getHttpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
}

func login(apiEnv string, apiKey string) string {
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

func RunClient(function string, apiEnv string, apiKey string, headers map[string]string, payloads []map[string]interface{}) {
	token := login(apiEnv, apiKey)
	//token := "kO88yLep3Uo93qffH83TfUOzOgwkedGkglN8Z+LPbVl9C6ST/0R9DaLvRkH6mj+wO8qMAloxxDPVXmDJ5IhgITr4ftI5MPwOFXshboli+4OCz/qR99080cD60bWZsriPGb5MbkfOYa7fI4gYarpWKYpRXNjqPTcYwIbSIS9kWQbJ7L9gGo3kfwj14yaUUhNqu+2q7U42m9xn+p/K5CQPJFrjBB5b9WQwNanYrAsn0cF2vlEQOJaQb1tZiaoZ+Mdz/zchQ1Z9J+5SltPk0at9cd8ZsBBZykndFMiKiuK8IOeYpZW0nLqqyj9yLZ9/9LDYtTBHGtF3H6BRi284KJeacDX/QDNWVo3Yd3MSHIkVwFOEw5h9wqsaorhV6jzBfaf3Zd4sEnhy68oaimF+neKCskwru6UQw0pSrYHawGcSu2OzV2RNPUJCisKDcbWZSKtn2vBSc6tzZj3yGzCtnQoMo2R8ZhuP6TSZ41EMacKvHtqlM7uoMZoW9TFxqCYQm87J5WJ3fMUtBQm3u9vBWxDOnSoF+x7nD3H3Z/gDZ4LdHSHNTLPbX0OTL+oiqtTP1yprifY4dEwJBCPKmmXey5bQfClEI/3ImywDXggjcXuOHFjyXMuP80+RxrT5OC2mKxS6QTrMcZzfij12ygOXqjBhyfWpkygDPAGpzwyCCs1m9XTKodFMZqbFSnDUdwNW2CdkbBc76tcOjfeqdX77xrZZbcW0m2MPDCmlXmf2wrR7+r/8CvspMUmegi7n0wW/SMyb96tOsdpSLsJBA2gIXwba19LKCJS5C44MqoZ2j2GaudiVV8zfpks2237ESmnfQcG0YDnWktCzvBb9X8gOAVl78h9soXGcb6e5XkSVncQdAW9BG0x3onraSvqyk/LXmBxkBpUFv54cKzrSlUTjbcXrEVoxj5HqjDR4PHGfHZ3Dnx0kUzRy/btqDLYkE3vdLboTwHgQRH18/epzC4MliNDS8C9IY3K6YI6ANGFt6SnvEw/DOlGPlvVQb/oodGWSr/YKYkr34TZM135hzmcXfDz59vqNxb8mcXqRoMByrWB2oYN5J3P++ba1AL8HDmILiuZO6LmJmJsrefUvpjIgMCLkuATQtnASarNll/ORA//Z+/JVedJKn8QeOx74hWctA0IolmAnkC3Cc0lXA5ljvSgd6rKTSTKz1DZtPoRglwsdSQlULOJ5hgYDbPqwnC/+O+KPD+y6oCvX9nMGZXKLTR+3C3TlkZQgtAk11tCu4HWFyykuV+EqThTXXKFgVhF2C8C9Av/CaG8n2KPIgvu5M13hNJ0dKYTUaleY57Gw62Wb7oyZy0nDLcb8PTqbCw/qleMmVJkxGROXikIsDWkH9U5psfy74HcqmPe/6GWYcT9TGhE="

	functionMap := map[string]func(
		apiEnv string, headers map[string]string, payloads []map[string]interface{}, token string, apiKey string,
	) []responseStruct{
		"sites": sites,
	}
	if _, ok := functionMap[function]; !ok {
		panic("Function not found")
	}

	functionCall := functionMap[function](apiEnv, headers, payloads, token, apiKey)
	fmt.Printf("%v\n", functionCall)

}

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	log.Printf("took %s", elapsed)
}

func main() {
	defer timeTrack(time.Now())
	RunClient(
		"sites",
		"prod",
		"",
		nil,
		[]map[string]interface{}{{"command": "list", "customer_name": "soctest"}},
	)

}

func processRequests(
	url string, headers map[string]string, payloads []map[string]interface{}, token string,
	apiEnv string, apiKey string,
) []responseStruct {
	httpClient := getHttpClient()

	semaphore := make(chan struct{}, 100)

	var responses []responseStruct
	responsesChan := make(chan responseStruct)
	var wg sync.WaitGroup

	mutexToken := muToken{
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
	payload map[string]interface{}, mutexToken *muToken, url string, headers map[string]string,
	httpClient *http.Client, responseChan chan responseStruct, wg *sync.WaitGroup, semaphore chan struct{},
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
		mutexToken.token = login(apiEnv, apiKey)
		processPostRequest(payload, mutexToken, url, headers, httpClient, responseChan, wg, semaphore, apiEnv, apiKey)
		return
	}

	payloadMarker, _ := payload["marker_var"].(string)

	responseChan <- responseStruct{
		valid:      rawResponse.StatusCode == 200,
		body:       jsonData["Ok"],
		statusCode: rawResponse.StatusCode,
		marker:     payloadMarker,
	}

	wg.Done()
	<-semaphore
}
