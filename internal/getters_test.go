package internal

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

func Test_getApiEnvHost(t *testing.T) {
	envExamples := map[string]string{
		"non_predefined": "non_predefined.threatx.io",
		"prod":           "api.threatx.io",
		"pod":            "api-tx-us-east-2a.threatx.io",
	}
	for key, value := range envExamples {
		assert.Equal(t, "https://"+value, getApiEnvHost(key))
	}

	assert.Panics(t, func() {
		getApiEnvHost("")
	})

}

func Test_getApiVerLink(t *testing.T) {
	incorrectVersions := []uint8{0, 3, 15}
	for _, version := range incorrectVersions {
		assert.Panics(t, func() {
			getApiVerLink(version)
		})
	}

	correctVersions := []uint8{1, 2}
	for _, version := range correctVersions {
		assert.Equal(t, "/tx_api/v"+strconv.Itoa(int(version)), getApiVerLink(version))
	}
}

func Test_getHttpClient(t *testing.T) {
	httpClient := getHttpClient()

	assert.IsType(t, &http.Client{}, httpClient)
	assert.True(t, httpClient.Transport.(*http.Transport).TLSClientConfig.InsecureSkipVerify)
}
