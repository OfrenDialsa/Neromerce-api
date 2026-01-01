package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ofrendialsa/neromerce/middlewares"
	"github.com/ofrendialsa/neromerce/modules/auth/service"
	"github.com/ofrendialsa/neromerce/modules/user/controller"
	"github.com/ofrendialsa/neromerce/pkg/constants"
	"github.com/samber/do"
)

func RegisterRoutes(server *gin.Engine, injector *do.Injector) {
	userController := do.MustInvoke[controller.UserController](injector)
	jwtService := do.MustInvokeNamed[service.JWTService](injector, constants.JWTService)

	userRoutes := server.Group("/api/user")
	{
		userRoutes.GET("", middlewares.Authenticate(jwtService), middlewares.Authorize("admin"), userController.GetAllUser)
		userRoutes.GET("/me", middlewares.Authenticate(jwtService), userController.Me)
		userRoutes.PUT("/:id", middlewares.Authenticate(jwtService), userController.Update)
		userRoutes.DELETE("/:id", middlewares.Authenticate(jwtService), userController.Delete)
	}
}
