package internal

import (
	"fmt"
)

func Sites(apiEnv string, headers map[string]string, payloads []map[string]interface{}, token string, apiKey string) []ResponseStruct {
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
