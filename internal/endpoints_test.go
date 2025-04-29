/*
This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.

SPDX-License-Identifier: MPL-2.0

File: endpoints_test.go
Description: Tests for endpoints.go
Author: tengzl33t
*/

package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getEndpointMap(t *testing.T) {
	assert.Panics(t, func() {
		getEndpoint("non_existing")
	})

	sitesTestEndpoint := getEndpoint("sites")
	assert.IsType(t, 0, sitesTestEndpoint[0])
	assert.IsType(t, []string{}, sitesTestEndpoint[1])
}
