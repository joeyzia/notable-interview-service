package config

import (
	"net/http"
)

type Config struct {
	Detail string `json:"detail"`
	Port string `json:"port"`
}

// ClientDoer - API Request Interface
type ClientDoer interface {
	Do(*http.Request) (*http.Response, error)
}
