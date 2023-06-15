package grpcapi

import (
	"github.com/watchtowerai/watchtower_go_libraries/pkg/modules/grpc/grpcserver"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

// RegistrarParams contains the list of services to be bound to API
type RegistrarParams struct {
	fx.In
}

// NewRegistrar provides a was of binding API and handlers (services)
func NewRegistrar(p RegistrarParams) grpcserver.APIRegistrarF {
	return func(s *grpc.Server) {

	}
}
