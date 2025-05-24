//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
)

func injectAppDependencies() (Application, error) {
	wire.Build(
		provideAppConfig,
		providePostgresConfig,
		provideHttpServerConfig,
		provideDBConnection,
		provideHTTPServer,
		NewApp,
	)

	return nil, nil
}
