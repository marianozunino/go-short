package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	hxhttp "github.com/maragudk/gomponents-htmx/http"
	"github.com/marianozunino/go-short/internal/config"
	"github.com/marianozunino/go-short/internal/store"
	"github.com/marianozunino/go-short/internal/utils"
	"github.com/marianozunino/go-short/internal/view"
)

type UrlHandler struct {
	db *store.Queries
}

func NewUrlHandler(q *store.Queries) UrlHandler {
	return UrlHandler{q}
}

func (UrlHandler) GetHomePage(c echo.Context) error {
	return view.Form().Render(
		c.Request().Context(),
		c.Response().Writer,
	)
}

func (h UrlHandler) GetShortenURL(c echo.Context) error {
	code := c.Param("code")

	if code == "" {
		return view.NotFound(code).Render(c.Request().Context(), c.Response().Writer)
	}

	urlModel, err := h.db.GetUrlByCode(c.Request().Context(), code)

	if err != nil {
		return view.NotFound(code).Render(c.Request().Context(), c.Response().Writer)
	}

	err = h.db.IncrementUrlHitsById(c.Request().Context(), urlModel.ID)
	if err != nil {
		return view.ErrorPartial("Something went wrong").Render(c.Request().Context(), c.Response().Writer)
	}

	// redirect
	return c.Redirect(http.StatusFound, urlModel.Url)
}

func (h UrlHandler) PostShortenURL(c echo.Context) error {
	url := c.FormValue("url")

	if url == "" {
		return view.ErrorPartial("URL cannot be empty").Render(c.Request().Context(), c.Response().Writer)
	}

	if !utils.IsValidURL(url) {
		return view.ErrorPartial("Invalid URL provided").Render(c.Request().Context(), c.Response().Writer)
	}

	hxhttp.SetPushURL(c.Response().Writer.Header(), c.Request().URL.RequestURI())

	urlModel, err := h.db.CreateUrl(c.Request().Context(), store.CreateUrlParams{
		Url:  url,
		Code: utils.GenerateShortKey(),
		Md5:  utils.Md5(url),
	})

	if err != nil {
		return view.ErrorPartial("Something went wrong").Render(c.Request().Context(), c.Response().Writer)
	}

	return view.Partial(
		fmt.Sprintf("%s/%s", config.BaseURL.Value(), urlModel.Code),
	).Render(
		c.Request().Context(),
		c.Response().Writer,
	)
}
