package middlewares

import (
	"os"
	"regexp"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GuardMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		urlsNotNeedAuthorization := []string{
			"/api/user/signin",
			"/api/user",
			"/api/docs/",
			"/api/question/filter",
		}

		currentURL := context.Request().URL.Path

		for _, urlPattern := range urlsNotNeedAuthorization {
			if strings.HasPrefix(currentURL, urlPattern) {
				return next(context)
			}
			if strings.HasSuffix(urlPattern, "/*") {
				urlPrefix := strings.TrimSuffix(urlPattern, "/*")
				matched, err := regexp.MatchString("^"+regexp.QuoteMeta(urlPrefix), currentURL)
				if err == nil && matched {
					return next(context)
				}
			}
		}

		authHeader := context.Request().Header.Get("Authorization")

		if authHeader == "" {
			return context.JSON(401, map[string]string{
				"message": "Unauthorized",
			})
		}

		jwtSecretKey := os.Getenv("JWT_SECRET")

		token, err := jwt.Parse(authHeader[7:], func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecretKey), nil
		})

		if err != nil {
			return context.JSON(401, map[string]string{
				"message": "Unauthorized",
			})
		}

		if !token.Valid {
			return context.JSON(401, map[string]string{
				"message": "Unauthorized",
			})
		}

		return next(context)
	}
}
