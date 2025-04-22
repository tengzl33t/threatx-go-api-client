package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

type responseStruct struct {
	valid      bool
	body       interface{}
	statusCode int
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

func sites(apiEnv string, payloads []map[string]interface{}) []responseStruct {
	//availableCommands := []string{
	//	"list",
	//	"new",
	//	"get",
	//	"delete",
	//	"update",
	//	"unset",
	//}

	url := fmt.Sprintf("%s%s/sites", getApiEnvHost(apiEnv), getApiVerLink(2))
	return processRequests(url, payloads)
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

	response := processRequests(url, jsonValues)[0]
	loginStatus := response.body.(map[string]interface{})["status"].(bool)
	if !loginStatus {
		panic("Could not login with key provided")
	}

	return response.body.(map[string]interface{})["token"].(string)
}

func RunClient(function string, apiEnv string, apiKey string, payloads []map[string]interface{}) {
	token := login(apiEnv, apiKey)
	for _, payload := range payloads {
		payload["token"] = token
	}

	functionMap := map[string]func(apiEnv string, payloads []map[string]interface{}) []responseStruct{
		"sites": sites,
	}
	if _, ok := functionMap[function]; !ok {
		println("Function not found")
		return
	}

	functionCall := functionMap[function](apiEnv, payloads)
	fmt.Printf("%v\n", functionCall)

}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func main() {
	RunClient(
		"sites",
		"prod",
		"",
		[]map[string]interface{}{{"command": "list", "customer_name": "threatx"}},
	)

}

func processRequests(url string, payloads []map[string]interface{}) []responseStruct {
	httpClient := getHttpClient()

	semaphore := make(chan struct{}, 100)

	var responses []responseStruct
	var wg sync.WaitGroup

	for _, payload := range payloads {
		semaphore <- struct{}{}

		wg.Add(1)
		go func() {
			preparedJsonPayload, _ := json.Marshal(payload)
			rawResponse, err := httpClient.Post(
				url, "application/json",
				bytes.NewBuffer(preparedJsonPayload),
			)
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
				panic(err)
			}
			println(jsonData["Ok"])

			responseStruct := responseStruct{
				valid:      rawResponse.StatusCode == 200,
				body:       jsonData["Ok"],
				statusCode: rawResponse.StatusCode,
			}

			responses = append(responses, responseStruct)
			wg.Done()
			<-semaphore
		}()
	}
	wg.Wait()

	return responses

}
