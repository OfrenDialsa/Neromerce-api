package category

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func RegisterRoutes(server *gin.Engine, injector *do.Injector) {
	// categoryController := do.MustInvoke[controller.CategoryController](injector)

	// categoryRoutes := server.Group("/api/category")
	// {
	// 	// TODO: add your endpoints here
	// }
}
