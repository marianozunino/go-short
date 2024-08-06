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

	db, err := sql.Open("sqlite3", config.DatabasePath.Value())
	if err != nil {
		log.Fatal(err)
	}

	queries := store.New(db)
	h := handlers.NewHandlers(queries)
	routes.SetupRoutes(e, h)

	e.Logger.Fatal(
		e.Start(
			fmt.Sprintf(":%s", config.Port.Value()),
		),
	)
}
