package main

import (
	"context"
	"embed"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"go-web-starter/controller"
)

type DB struct {}

//go:embed assets/img/* assets/fontello/font/* assets/fontello/css/* tmp/css/styles.css
var assets embed.FS

var staticResourcesHandler = echo.WrapHandler(http.FileServer(http.FS(assets)))
var staticResourcesRewrite = middleware.Rewrite(controller.AssetsRewriteRules(assets))

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8000"
    }

    e := echo.New()

    e.GET("/*", staticResourcesHandler, staticResourcesRewrite)

    userController := controller.UserController{}
    e.Use(withUser)
    e.GET("/", userController.ShowDashboard)

    e.Logger.Fatal(e.Start(":"+port))
}

func withUser(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        ctx := context.WithValue(c.Request().Context(), "user", "test@test.ch")
        c.SetRequest(c.Request().WithContext(ctx))
        return next(c)
    }
}
