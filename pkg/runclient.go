// SPDX-License-Identifier: MPL-2.0

package pkg

import (
	"errors"
	"fmt"
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

	responses := internal.SendRequests(endpoint, apiEnv, headers, payloads, token, apiKey)

	fmt.Printf("%v\n", responses)
}
