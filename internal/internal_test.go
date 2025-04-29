/*
This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.

SPDX-License-Identifier: MPL-2.0

File: internal_test.go
Description: Tests for internal.go
Author: tengzl33t
*/

package internal

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func Test_getURL(t *testing.T) {
	examples := map[string][3]string{
		"https://test.threatx.io/tx_api/v1/example":   {"test", "1", "example"},
		"https://test1.threatx.io/tx_api/v2/example2": {"test1", "2", "example2"},
	}

	for key, values := range examples {
		uintValue, err := strconv.ParseUint(values[1], 10, 8)
		if err != nil {
			panic(err)
		}
		assert.Equal(t, key, getURL(values[0], uint8(uintValue), values[2]))
	}

	assert.Panics(t, func() {
		getURL("", 0, "")
	})
}
