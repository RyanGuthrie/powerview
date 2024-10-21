package endpoint

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"log/slog"
	"powerview/internal/client/gateway/gatewayhttpclient"
	"text/template"
)

var Producers = make(map[Name]ProducerFunc)

type ProducerFunc func(client gatewayhttpclient.GatewayHttpClient) Endpoint

type Endpoint interface {
	Verb() string
	Execute() (any, error)
}

type StandardPath interface {
	Path() string
}

type TemplatePath interface {
	Path() (*template.Template, error)
	PathTemplateArgs() any
}

type Name int

const (
	HomeAutomations  Name = iota
	HomeAutomationId Name = iota
)
const (
	hostname = "powerview-g3.local"
	port     = 80
)

func Execute[R any](client gatewayhttpclient.GatewayHttpClient, endpoint Endpoint) (any, error) {
	var parsedResp R

	path, err := pathFrom(endpoint)
	if err != nil {
		return parsedResp, err
	}

	fullPath := fmt.Sprintf("http://%s:%d%s", hostname, port, path)
	slog.Info(fmt.Sprintf("Calling endpoint: %s", fullPath))

	resp, err := client.Get(fullPath)
	if err != nil {
		return parsedResp, errors.Wrap(err, fmt.Sprintf("failed to %s %s", endpoint.Verb(), path))
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return parsedResp, errors.Wrap(err, "failed to read response body")
	}

	if err := json.Unmarshal(body, &parsedResp); err != nil {
		return parsedResp, errors.Wrap(err, "failed to unmarshal response body")
	}

	return parsedResp, nil
}

func pathFrom(endpoint Endpoint) (string, error) {
	var path string
	switch e := endpoint.(type) {
	case StandardPath:
		path = e.Path()
	case TemplatePath:
		pathTemplate, err := e.Path()
		if err != nil {
			return "", errors.Wrap(err, "failed to get path template")
		}

		var buf = bytes.Buffer{}
		if err := pathTemplate.Execute(&buf, e.PathTemplateArgs()); err != nil {
			return "", errors.Wrap(err, "failed to execute path template")
		}

		path = string(buf.Bytes())
	}
	return path, nil
}
