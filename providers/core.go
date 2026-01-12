package providers

import (
	"github.com/ofrendialsa/neromerce/config"
	authController "github.com/ofrendialsa/neromerce/modules/auth/controller"
	authRepo "github.com/ofrendialsa/neromerce/modules/auth/repository"
	authService "github.com/ofrendialsa/neromerce/modules/auth/service"
	categoryController "github.com/ofrendialsa/neromerce/modules/category/controller"
	categoryRepo "github.com/ofrendialsa/neromerce/modules/category/repository"
	categoryService "github.com/ofrendialsa/neromerce/modules/category/service"
	productController "github.com/ofrendialsa/neromerce/modules/product/controller"
	productRepo "github.com/ofrendialsa/neromerce/modules/product/repository"
	productService "github.com/ofrendialsa/neromerce/modules/product/service"
	userController "github.com/ofrendialsa/neromerce/modules/user/controller"
	"github.com/ofrendialsa/neromerce/modules/user/repository"
	userService "github.com/ofrendialsa/neromerce/modules/user/service"
	"github.com/ofrendialsa/neromerce/pkg/constants"
	"github.com/samber/do"
	"gorm.io/gorm"
)

func InitDatabase(injector *do.Injector) {
	do.ProvideNamed(injector, constants.DB, func(i *do.Injector) (*gorm.DB, error) {
		return config.SetUpDatabaseConnection(), nil
	})
}

func RegisterDependencies(injector *do.Injector) {
	InitDatabase(injector)

	do.ProvideNamed(injector, constants.JWTService, func(i *do.Injector) (authService.JWTService, error) {
		return authService.NewJWTService(), nil
	})

	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)
	jwtService := do.MustInvokeNamed[authService.JWTService](injector, constants.JWTService)

	userRepository := repository.NewUserRepository(db)
	refreshTokenRepository := authRepo.NewRefreshTokenRepository(db)
	categoryRepository := categoryRepo.NewCategoryRepository(db)
	productRepository := productRepo.NewProductRepository(db)

	userService := userService.NewUserService(userRepository, db)
	authService := authService.NewAuthService(userRepository, refreshTokenRepository, jwtService, db)
	categoryService := categoryService.NewCategoryService(categoryRepository)
	productService := productService.NewProductService(productRepository, db)

	do.Provide(
		injector, func(i *do.Injector) (userController.UserController, error) {
			return userController.NewUserController(i, userService), nil
		},
	)

	do.Provide(
		injector, func(i *do.Injector) (authController.AuthController, error) {
			return authController.NewAuthController(i, authService), nil
		},
	)

	do.Provide(
		injector, func(i *do.Injector) (categoryController.CategoryController, error) {
			return categoryController.NewCategoryController(categoryService), nil
		},
	)

	do.Provide(
		injector, func(i *do.Injector) (productController.ProductController, error) {
			return productController.NewProductController(i, productService), nil
		},
	)
}
