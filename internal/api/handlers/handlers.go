package handlers

import "github.com/marianozunino/go-short/internal/store"

type Handlers struct {
	UrlHandler UrlHandler
}

func NewHandlers(q *store.Queries) Handlers {
	return Handlers{
		UrlHandler: NewUrlHandler(q),
	}
}

