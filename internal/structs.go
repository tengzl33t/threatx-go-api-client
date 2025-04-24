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
