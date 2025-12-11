package category

import (
	"github.com/gin-gonic/gin"
	"github.com/ofrendialsa/neromerce/middlewares"
	"github.com/ofrendialsa/neromerce/modules/auth/service"
	"github.com/ofrendialsa/neromerce/modules/category/controller"
	"github.com/ofrendialsa/neromerce/pkg/constants"
	"github.com/samber/do"
)

func RegisterRoutes(server *gin.Engine, injector *do.Injector) {
	categoryController := do.MustInvoke[controller.CategoryController](injector)
	jwtService := do.MustInvokeNamed[service.JWTService](injector, constants.JWTService)

	categoryRoutes := server.Group("/api/category")
	{
		categoryRoutes.GET("", categoryController.GetAll)
		categoryRoutes.POST("", middlewares.Authenticate(jwtService), categoryController.Create)
		categoryRoutes.GET("/:id", categoryController.GetCategoryByID)
		categoryRoutes.DELETE("/:id", middlewares.Authenticate(jwtService), categoryController.Delete)
	}
}
