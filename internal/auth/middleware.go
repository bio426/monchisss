package auth

import (
	"net/http"
	"slices"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

const CtxTokenKey = "userToken"
const CtxUserIdKey = "userId"
const CtxUserRoleKey = "userRole"
const CtxUserStoreKey = "userStore"

func setContextUser(c echo.Context) {
	token := c.Get(CtxTokenKey).(*jwt.Token)
	claims, _ := token.Claims.(*CustomClaims)
	c.Set(CtxUserIdKey, claims.Id)
	c.Set(CtxUserRoleKey, claims.Role)
	c.Set(CtxUserStoreKey, claims.Store)
}

// This should be only used after the auth.Middleware has processed the request
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

var Middleware echo.MiddlewareFunc = echojwt.WithConfig(echojwt.Config{
	SigningKey:     []byte(JwtSecret),
	TokenLookup:    "cookie:" + CookieName,
	ContextKey:     CtxTokenKey,
	SuccessHandler: setContextUser,
	NewClaimsFunc: func(c echo.Context) jwt.Claims {
		return &CustomClaims{}
	},
})
