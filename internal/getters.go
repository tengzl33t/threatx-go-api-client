package internal

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

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

func getHttpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
}
