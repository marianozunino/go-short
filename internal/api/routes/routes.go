package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/marianozunino/go-short/internal/api/handlers"
)

func SetupRoutes(e *echo.Echo, h handlers.UrlHandler) {
	e.GET("/", h.GetHomePage)
	e.GET("/:code", h.GetShortenURL)
	e.POST("/", h.PostShortenURL)
}
