package products

import (
	"github.com/gin-gonic/gin"
	"github.com/ofrendialsa/neromerce/middlewares"
	"github.com/ofrendialsa/neromerce/modules/auth/service"
	"github.com/ofrendialsa/neromerce/modules/product/controller"
	"github.com/ofrendialsa/neromerce/pkg/constants"
	"github.com/samber/do"
)

func RegisterRoutes(server *gin.Engine, injector *do.Injector) {
	productController := do.MustInvoke[controller.ProductController](injector)
	jwtService := do.MustInvokeNamed[service.JWTService](injector, constants.JWTService)
	authAccess := middlewares.Authenticate(jwtService)
	roleAccess := middlewares.Authorize("admin")

	productRoutes := server.Group("/api/products")
	{
		productRoutes.GET("", productController.GetAll)
		productRoutes.POST("", authAccess, roleAccess, productController.Create)
		// productRoutes.DELETE("/:id", authAccess, roleAccess, productController.Delete)
	}
}
