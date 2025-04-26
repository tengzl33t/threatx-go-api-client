// SPDX-License-Identifier: MPL-2.0

package pkg

import (
	"errors"
	"fmt"
	"strings"
	"threatx-go-api-client/internal"
)

func validateRunArguments(endpoint string, apiEnv string, apiKey string) error {
	if endpoint == "" {
		return errors.New("no endpoint provided")
	}
	if apiEnv == "" {
		return errors.New("no API environment provided")
	}
	if apiKey == "" {
		return errors.New("no API key provided")
	}
	return nil
}

func RunClient(endpoint string, apiEnv string, apiKey string, headers map[string]string, payloads []map[string]interface{}) {
	err := validateRunArguments(endpoint, apiEnv, apiKey)
	if err != nil {
		panic(err)
	}

	token := internal.Login(apiEnv, apiKey)

	endpointMap := map[string]func(
		apiEnv string, headers map[string]string, payloads []map[string]interface{}, token string, apiKey string,
	) []internal.ResponseStruct{
		"Sites": internal.Sites,
	}

	capitalizedEndpointString := strings.ToUpper(endpoint[:1]) + strings.ToLower(endpoint[1:])

	if _, ok := endpointMap[capitalizedEndpointString]; !ok {
		panic(fmt.Sprintf("Function '%s' not found", capitalizedEndpointString))
	}

	functionCall := endpointMap[capitalizedEndpointString](apiEnv, headers, payloads, token, apiKey)
	fmt.Printf("%v\n", functionCall)
}
