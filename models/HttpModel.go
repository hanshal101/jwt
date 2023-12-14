package models

import "net/http"

type Response struct {
	Error   error
	Status  http.ConnState
	Message string
}
