package main

import (
	"embed"

	"github.com/spf13/viper"
	grpcapi "github.com/watchtowerai/service_template/api/grpc"
	restapi "github.com/watchtowerai/service_template/api/rest"
	"github.com/watchtowerai/service_template/clients"
	"github.com/watchtowerai/watchtower_go_libraries/pkg/clients/bugsnag"
	"github.com/watchtowerai/watchtower_go_libraries/pkg/clients/datadog"
	"github.com/watchtowerai/watchtower_go_libraries/pkg/config"
	log "github.com/watchtowerai/watchtower_go_libraries/pkg/log"
	grpcmw "github.com/watchtowerai/watchtower_go_libraries/pkg/middleware/grpc"
	"github.com/watchtowerai/watchtower_go_libraries/pkg/modules/grpc/grpcserver"
	"github.com/watchtowerai/watchtower_go_libraries/pkg/modules/httpserver"
	"github.com/watchtowerai/watchtower_go_libraries/pkg/modules/logging"
	"github.com/watchtowerai/watchtower_go_libraries/pkg/modules/recovery"
	"github.com/watchtowerai/watchtower_go_libraries/pkg/modules/router"
	recovery2 "github.com/watchtowerai/watchtower_go_libraries/pkg/recovery"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc/health"
)

// main starts the service process.
func main() {
	app := fx.New(
		fx.Provide(
			// general constructors
			newConfig,
			logging.NewZapLoggerWithBugsnag,
			log.NewSugaredLoggerFactory,
			bugsnag.NewConfiguration,
			bugsnag.NewBugsnagNotifier,
			recovery.NewBugsnagHandler,
			newSugarLogger,
			newRecovery,

			// server specific constructor - select one group
			// 1) rest http
			router.NewMiddlewareProvider,
			router.NewRouter,
			restapi.NewRegisterEndpoints,

			// 2) grpc
			grpcmw.NewProvider,
			grpcapi.NewRegistrar,
			health.NewServer,
		),
		clients.Module,
		fx.Invoke(
			datadog.NewTracer,

			// server specific - select one
			grpcserver.NewGRPCServer,
			httpserver.NewServer,
		),
	)
	app.Run()
}

//go:embed config
var cfg embed.FS

func newConfig() (*viper.Viper, error) {
	return config.NewEmbedConfig(cfg)
}

// newSugarLogger helper
func newSugarLogger(zapLog *zap.Logger) *zap.SugaredLogger {
	return zapLog.Sugar()
}

func newRecovery(bsh *recovery.BugsnagHandler) recovery2.Handler {
	return bsh
}
