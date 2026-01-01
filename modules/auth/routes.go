package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/ofrendialsa/neromerce/middlewares"
	"github.com/ofrendialsa/neromerce/modules/auth/controller"
	"github.com/ofrendialsa/neromerce/modules/auth/service"
	"github.com/ofrendialsa/neromerce/pkg/constants"
	"github.com/samber/do"
)

func RegisterRoutes(server *gin.Engine, injector *do.Injector) {
	authController := do.MustInvoke[controller.AuthController](injector)
	jwtService := do.MustInvokeNamed[service.JWTService](injector, constants.JWTService)

	authRoutes := server.Group("/api/auth")
	{
		authRoutes.POST("/register", authController.Register)
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/refresh", authController.RefreshToken)
		authRoutes.POST("/logout", middlewares.Authenticate(jwtService), authController.Logout)
		authRoutes.POST("/send-verification-email", authController.SendVerificationEmail)
		authRoutes.POST("/verify-email", authController.VerifyEmail)
		authRoutes.POST("/send-password-reset", authController.SendPasswordReset)
		authRoutes.POST("/reset-password", authController.ResetPassword)
	}
}
