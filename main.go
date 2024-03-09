package main

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/bio426/monchisss/datasource"
	"github.com/bio426/monchisss/route"
)

func main() {
	datasource.InitConfig()
	_, err := datasource.InitPostgres()
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.Debug = true
	e.HideBanner = true
	e.Validator = &CustomValidator{validator: validator.New()}

	if datasource.Config.PROD {
		e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
			Root:  "./dist",
			HTML5: true,
		}))
	} else {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     []string{"http://localhost:5173"},
			AllowCredentials: true,
		}))
		e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
			LogMethod:   true,
			LogURI:      true,
			LogStatus:   true,
			LogError:    true,
			HandleError: true,
			LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
				if v.Error == nil {
					fmt.Printf("%s %s %d\n", v.Method, v.URI, v.Status)
				} else {
					fmt.Printf("%s %s %d %s\n", v.Method, v.URI, v.Status, v.Error.Error())
				}
				return nil
			},
		}))
	}

	route.RegisterRest(e)

	if datasource.Config.PROD {
		e.Logger.Fatal(e.Start(":8080"))
	} else {
		e.Logger.Fatal(e.Start(":1323"))
	}
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
