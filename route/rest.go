package route

import (
	"github.com/labstack/echo/v4"

	"github.com/bio426/monchisss/internal/auth"
	"github.com/bio426/monchisss/internal/store"
	"github.com/bio426/monchisss/internal/user"
)

func RegisterRest(app *echo.Echo) {
	restGroup := app.Group("/rest")

	authGoup := restGroup.Group("/auth")
	authGoup.POST("/login", auth.Controller.Login)
	authGoup.POST("/logout", auth.Controller.Logout)

	userGroup := restGroup.Group("/user")
	userGroup.GET("", user.Controller.List)
	userGroup.POST("", user.Controller.Create)

	storeGroup := restGroup.Group("/store")
	storeGroup.GET("", store.Controller.List)
	storeGroup.POST("", store.Controller.Create)
}
