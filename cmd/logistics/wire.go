//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	internal "github.com/coopnorge/interview-backend/internal/logistics"
	"github.com/coopnorge/interview-backend/internal/logistics/config"
	"github.com/coopnorge/interview-backend/internal/logistics/services/client"
	"github.com/coopnorge/interview-backend/internal/logistics/services/operator"

	"github.com/google/wire"
)

// newWire create new DI
func newWire(cfg *config.ClientAppConfig) (*internal.ServiceInstance, func(), error) {
	panic(wire.Build(
		client.ServiceSetForClient,
		operator.ServiceSetForOperator,
		internal.NewServiceInstance,
	))
}
