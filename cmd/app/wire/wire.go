//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"powerview/internal/app"
	"powerview/internal/client/gateway"
	"powerview/internal/client/gateway/gatewayhttpclient"
)

// Run the `wire` command to update the auto-generated dependency injection code
//go:generate wire

func CreateApp() (app.App, error) {
	wire.Build(
		app.NewApp,
		server.NewServer,
		gatewayhttpclient.NewGatewayHttpClient,
		gateway.NewInstance,
	)

	return app.App{}, nil
}
