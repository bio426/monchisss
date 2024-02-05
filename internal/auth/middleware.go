package auth

import (
	"net/http"
	"slices"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

const CtxTokenKey = "userToken"
const CtxUserIdKey = "userId"
const CtxUserRoleKey = "userRole"

func setContextUser(c echo.Context) {
	user := c.Get(CtxTokenKey).(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId, _ := claims.GetIssuer()
	userRole, _ := claims.GetSubject()
	c.Set(CtxUserIdKey, userId)
	c.Set(CtxUserRoleKey, userRole)
}

func MiddlewareWithRoles(permittedRoles []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userRole := c.Get(CtxUserRoleKey).(string)
			if !slices.Contains(permittedRoles, userRole) {
				return echo.NewHTTPError(http.StatusForbidden)
			}
			return next(c)
		}
	}
}

func MiddlewareWithSkipper(skippedPaths []string) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:     []byte(JwtSecret),
		TokenLookup:    "cookie:" + CookieName,
		ContextKey:     CtxTokenKey,
		SuccessHandler: setContextUser,
		Skipper: func(c echo.Context) bool {
			for _, path := range skippedPaths {
				if strings.HasSuffix(c.Path(), path) {
					return true
				}
			}
			return false
		},
	})
}

var Middleware echo.MiddlewareFunc = echojwt.WithConfig(echojwt.Config{
	SigningKey:     []byte(JwtSecret),
	TokenLookup:    "cookie:" + CookieName,
	ContextKey:     CtxTokenKey,
	SuccessHandler: setContextUser,
})
