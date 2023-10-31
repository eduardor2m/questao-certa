package http

import (
	"net/http"
	"os"

	"github.com/eduardor2m/questao-certa/internal/adapters/delivery/http/middlewares"
	"github.com/eduardor2m/questao-certa/internal/adapters/delivery/http/routes"
	"github.com/eduardor2m/questao-certa/tools/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type API interface {
	Serve()
	loadRoutes()
}

type Options struct {
}

type api struct {
	options      *Options
	group        *echo.Group
	echoInstance *echo.Echo
}

// NewAPI
// @title Questão Certa API
// @version 1.0
// @description API para gerenciamento de questões e respostas
// @contact.name Eduardo Melo
// @contact.email deveduardomelo@gmail.com
// @BasePath /api
// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
func NewAPI(options *Options) API {
	echoInstance := echo.New()
	echoInstance.HideBanner = true

	logger.Info("Starting API")

	echoInstance.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/api/docs/index.html")
	})

	return &api{
		options:      options,
		group:        echoInstance.Group("/api"),
		echoInstance: echoInstance,
	}
}

func (instance *api) Serve() {
	instance.loadRoutes()
	instance.echoInstance.Use(instance.getCORSSettings())

	port := os.Getenv("PORT")
	instance.echoInstance.Logger.Fatal(instance.echoInstance.Start(":" + port))

}

func (instance *api) loadRoutes() {
	router := routes.New()

	router.Load(instance.group)
}

func (instance *api) getCORSSettings() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:         middlewares.OriginInspectSkipper,
		AllowOriginFunc: middlewares.VerifyOrigin,
		AllowMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodPatch,
		},
	})
}
