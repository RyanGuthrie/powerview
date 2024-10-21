package endpoint

import (
	"fmt"
	"net/http"
	"powerview/internal/client/gateway/gatewayhttpclient"
	"text/template"
)

var _ Endpoint = (*HomeAutomationIdEndpoint)(nil)

type HomeAutomationIdEndpoint struct {
	client gatewayhttpclient.GatewayHttpClient
}

func init() {
	Producers[HomeAutomationId] = NewHomeAutomationId
}

func NewHomeAutomationId(client gatewayhttpclient.GatewayHttpClient) Endpoint {
	return &HomeAutomationIdEndpoint{client: client}
}

func (h *HomeAutomationIdEndpoint) Verb() string {
	return http.MethodGet
}

func (h *HomeAutomationIdEndpoint) Path() (*template.Template, error) {
	return template.New("HomeAutomationsId").Parse("/home/automations/{{.ID}}")
}

func (h *HomeAutomationIdEndpoint) PathTemplateArgs() any {
	return struct {
		ID string
	}{
		ID: "24",
	}
}

func (h *HomeAutomationIdEndpoint) Execute() (any, error) {
	response, err := Execute[HomeAutomation](h.client, h)
	if err != nil {
		return "", err
	}

	fmt.Println(response)

	return "", nil
}
