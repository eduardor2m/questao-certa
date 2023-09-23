package middlewares

import (
	"regexp"
	"strings"

	"github.com/eduardor2m/questao-certa/internal/adapters/delivery/http/dicontainer"
	"github.com/labstack/echo/v4"
)

func Admin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		urlsNotNeedAuthorization := []string{
			// "/api/user/signin",
			// "/api/user",
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

		userServices := dicontainer.GetUserServices()

		userType, err := userServices.VerifyUserIsLoggedOrAdmin(authHeader[7:])

		if err != nil {
			return context.JSON(401, map[string]string{
				"message": "user not admin, user type is: " + *userType,
			})
		}

		if *userType != "admin" {
			return context.JSON(401, map[string]string{
				"message": "user not admin, user type is: " + *userType,
			})
		}

		return next(context)
	}
}
