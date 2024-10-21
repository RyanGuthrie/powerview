package endpoint

import (
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"powerview/internal/client/gateway/gatewayhttpclient"
)

var _ Endpoint = (*HomeAutomationsEndpoint)(nil)
var _ StandardPath = (*HomeAutomationsEndpoint)(nil)

type HomeAutomationsEndpoint struct {
	client gatewayhttpclient.GatewayHttpClient
}

func init() {
	Producers[HomeAutomations] = NewHomeAutomation
}
func NewHomeAutomation(client gatewayhttpclient.GatewayHttpClient) Endpoint {
	return &HomeAutomationsEndpoint{client: client}
}

func (h *HomeAutomationsEndpoint) Verb() string {
	return http.MethodGet
}

func (h *HomeAutomationsEndpoint) Path() string {
	return "/home/automations"
}

func (h *HomeAutomationsEndpoint) PathTemplateArgs() any {
	return struct{}{}
}

func (h *HomeAutomationsEndpoint) Execute() (any, error) {
	response, err := Execute[HomeAutomationResponse](h.client, h)
	if err != nil {
		return "", errors.Wrap(err, "failed to GET /home/automations")
	}

	fmt.Println(response)

	return response, nil
}

type HomeAutomationResponse []HomeAutomation

type HomeAutomation struct {
	Id          int  `json:"id"`
	Type        int  `json:"type"`
	Enabled     bool `json:"enabled"`
	Days        int  `json:"days"`
	Hour        int  `json:"hour"`
	Min         int  `json:"min"`
	BleId       int  `json:"bleId"`
	SceneId     int  `json:"sceneId"`
	ErrorShdIds any  `json:"errorShd_Ids"`
}
