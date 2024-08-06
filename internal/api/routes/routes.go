package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/marianozunino/go-short/internal/api/handlers"
)

func SetupRoutes(e *echo.Echo, h handlers.Handlers) {
	e.GET("/", h.UrlHandler.GetHomePage)
	e.GET("/:code", h.UrlHandler.GetShortenURL)
	e.POST("/", h.UrlHandler.PostShortenURL)
}
