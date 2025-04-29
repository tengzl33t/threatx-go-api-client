/*
This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.

SPDX-License-Identifier: MPL-2.0

File: getters_test.go
Description: Tests for getters.go
Author: tengzl33t
*/

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
