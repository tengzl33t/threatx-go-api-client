/*
This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.

SPDX-License-Identifier: MPL-2.0

File: getters.go
Description: Get functions
Author: tengzl33t
*/

package internal

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"strconv"
)

func getApiVerLink(version uint8) string {
	if version < 1 || version > 2 {
		panic("Invalid version number " + strconv.Itoa(int(version)))
	}
	return fmt.Sprintf("/tx_api/v%d", version)
}

func getApiEnvHost(apiEnv string) string {
	if apiEnv == "" {
		panic("No API environment specified")
	}

	domainPart := "threatx.io"
	oldHostParts := map[string]string{
		"prod":    "",
		"pod":     "tx-us-east-2a",
		"qa":      "qa0",
		"dev":     "dev0",
		"staging": "staging0",
	}
	if _, ok := oldHostParts[apiEnv]; ok {
		delimiter := ""
		if apiEnv != "prod" {
			delimiter = "-"
		}

		return fmt.Sprintf("https://api%s%s.%s", delimiter, oldHostParts[apiEnv], domainPart)
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
