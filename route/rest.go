package route

import (
	"github.com/labstack/echo/v4"

	"github.com/bio426/monchisss/internal/auth"
	"github.com/bio426/monchisss/internal/order"
	"github.com/bio426/monchisss/internal/product"
	"github.com/bio426/monchisss/internal/store"
	"github.com/bio426/monchisss/internal/user"
	"github.com/bio426/monchisss/internal/wa"
)

func RegisterRest(app *echo.Echo) {
	restGroup := app.Group("/rest")

	authGoup := restGroup.Group("/auth")
	authGoup.POST("/login", auth.Controller.Login)
	authGoup.POST("/logout", auth.Controller.Logout)

	userGroup := restGroup.Group("/user")
	userGroup.GET("", user.Controller.List)
	userGroup.POST("", user.Controller.Create)
	userGroup.GET("/owner/inactive", user.Controller.GetInactiveOwners)

	storeGroup := restGroup.Group("/store")
	storeGroup.GET("", store.Controller.List)
	storeGroup.POST("", store.Controller.Create)
	storeGroup.GET("/:id/user", store.Controller.ListUsers)
	storeGroup.POST("/:id/user", store.Controller.CreateUser)

	productGroup := restGroup.Group("/product")
	productGroup.Use(auth.Middleware, auth.MiddlewareWithRoles([]string{"owner"}))
	productGroup.GET("", product.Controller.List)
	productGroup.POST("", product.Controller.Create)
	productGroup.GET("/category", product.Controller.ListCategory)
	productGroup.POST("/category", product.Controller.CreateCategory)

	orderGroup := restGroup.Group("/order")
	orderGroup.GET("/:id", order.Controller.CategoriesById)

	waGroup := restGroup.Group("/wa")
	// url must be same for verify and recieve notifications
	waGroup.GET("/verify", wa.Controller.Verify)
	waGroup.POST("/verify", wa.Controller.Hook)
}
