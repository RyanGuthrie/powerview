package gatewayhttpclient

import (
	"net/http"
)

type GatewayHttpClient = *http.Client

func NewGatewayHttpClient() (GatewayHttpClient, error) {
	return http.DefaultClient, nil
}
