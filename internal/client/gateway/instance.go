package gateway

import (
	"fmt"
	"github.com/pkg/errors"
	"log/slog"
	"powerview/internal/client/gateway/endpoint"
	"powerview/internal/client/gateway/gatewayhttpclient"
)

type Instance struct {
	client            gatewayhttpclient.GatewayHttpClient
	EndpointProducers map[endpoint.Name]endpoint.ProducerFunc
}

func NewInstance(client gatewayhttpclient.GatewayHttpClient) (Instance, error) {
	return Instance{
		client:            client,
		EndpointProducers: endpoint.Producers,
	}, nil
}

func (i *Instance) DoStuff() error {
	resp, err := i.EndpointProducers[endpoint.HomeAutomations](i.client).Execute()
	if err != nil {
		return errors.Wrap(err, "failed to execute endpoint")
	}

	slog.Info(fmt.Sprintf("%v", resp))

	resp, err = i.EndpointProducers[endpoint.HomeAutomationId](i.client).Execute()
	if err != nil {
		return errors.Wrap(err, "failed to execute endpoint")
	}

	slog.Info(fmt.Sprintf("%v", resp))

	return nil
}
