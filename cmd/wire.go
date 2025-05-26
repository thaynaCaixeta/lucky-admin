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
		provideRepository,
		provideHTTPServer,
		provideGameService,
		provideGameHandler,
		NewApp,
	)

	return nil, nil
}
