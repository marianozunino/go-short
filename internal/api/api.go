package api

import (
	"database/sql"
	"log"

	_ "github.com/marianozunino/go-short/internal/config"
	_ "github.com/mattn/go-sqlite3"

	"github.com/labstack/echo/v4"
	"github.com/marianozunino/go-short/internal/handler"
	"github.com/marianozunino/go-short/internal/store"
)

func Run() {
	e := echo.New()

	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	queries := store.New(db)

	h := handler.NewHandler(queries)

	e.GET("/", h.GetHomePage)
	e.GET("/:code", h.GetShortenURL)
	e.POST("/", h.PostShortenURL)

	e.Logger.Fatal(e.Start(":1323"))
}
