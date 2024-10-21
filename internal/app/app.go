package app

import (
	"fmt"
	"net/http"
	"powerview/internal/client/gateway"
)

type App struct {
	server  *http.Server
	gateway gateway.Instance
}

func NewApp(gateway gateway.Instance, server *http.Server) (App, error) {
	return App{
		server:  server,
		gateway: gateway,
	}, nil
}

func (a *App) StartAndWait() error {
	var err error

	err = a.gateway.DoStuff()
	if err != nil {
		return err
	}

	if err = a.server.ListenAndServe(); err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}

	return nil
}
