/*
This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.

SPDX-License-Identifier: MPL-2.0

File: structs.go
Description: Structs
Author: tengzl33t
*/

package internal

import "sync"

type ResponseStruct struct {
	valid      bool
	body       interface{}
	statusCode int
	marker     string
}

type MuToken struct {
	mu    sync.RWMutex
	token string
}
