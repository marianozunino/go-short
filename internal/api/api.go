package api

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/marianozunino/go-short/internal/config"
	_ "github.com/mattn/go-sqlite3"

	"github.com/labstack/echo/v4"
	"github.com/marianozunino/go-short/internal/api/handlers"
	"github.com/marianozunino/go-short/internal/api/routes"
	"github.com/marianozunino/go-short/internal/store"
)

func Run() {
	e := echo.New()
	cfg, err := config.LoadConfig()

	db, err := sql.Open("sqlite3", cfg.DatabasePath)
	if err != nil {
		log.Fatal(err)
	}

	queries := store.New(db)
	h := handlers.NewUrlHandler(queries, *cfg)
	routes.SetupRoutes(e, h)

	e.Logger.Fatal(
		e.Start(
			fmt.Sprintf(":%d", cfg.Port),
		),
	)
}
