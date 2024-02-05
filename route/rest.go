package route

import (
	"github.com/labstack/echo/v4"

	"github.com/bio426/monchisss/internal/auth"
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

	waGroup := restGroup.Group("/wa")
    // url must be same for verify and recieve notifications
	waGroup.GET("/verify", wa.Controller.Verify)
	waGroup.POST("/verify", wa.Controller.Hook)
}
