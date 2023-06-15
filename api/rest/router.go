package restapi

import (
	"github.com/gin-gonic/gin"
	"github.com/watchtowerai/watchtower_go_libraries/pkg/modules/router"
)

// NewRegisterEndpoints returns function which can be used for instrumenting gin router
func NewRegisterEndpoints() router.RegisterEndpointsF {
	return registerEndpoints
}

func registerEndpoints(r gin.IRouter) {

}
