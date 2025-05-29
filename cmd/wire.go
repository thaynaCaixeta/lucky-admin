//go:build wireinject
// +build wireinject

package main

import (
	"context"

	"github.com/google/wire"
)

func injectAppDependencies(ctx context.Context) (Application, error) {
	wire.Build(
		provideAppConfig,
		provideDynamoConfig,
		provideDynamoClient,
		provideRepository,
		provideHttpServerConfig,
		provideHTTPServer,
		provideGameService,
		provideGameHandler,
		NewApp,
	)

	return nil, nil
}
