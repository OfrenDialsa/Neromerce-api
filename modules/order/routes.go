package order

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func RegisterRoutes(server *gin.Engine, injector *do.Injector) {
	// orderController := do.MustInvoke[controller.OrderController](injector)

	// orderRoutes := server.Group("/api/order")
	// {
	// 	// TODO: add your endpoints here
	// }
}
