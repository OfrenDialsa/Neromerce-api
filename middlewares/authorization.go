package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ofrendialsa/neromerce/modules/user/dto"
	"github.com/ofrendialsa/neromerce/pkg/utils"
)

func Authorize(allowedRoles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		roleI, exists := ctx.Get("user_role")
		if !exists {
			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, "User not logged in", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}

		role := roleI.(string)
		for _, r := range allowedRoles {
			if r == role {
				ctx.Next()
				return
			}
		}

		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, dto.MESSAGE_FAILED_FORBIDDEN, nil)
		ctx.AbortWithStatusJSON(http.StatusForbidden, res)
	}
}
