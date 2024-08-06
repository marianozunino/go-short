package config

import (
	"github.com/dogmatiq/ferrite"
)

var BaseURL = ferrite.
	URL("BASE_DOMAIN", "http://example.org").
	Required()

var Port = ferrite.
	NetworkPort("PORT", "1323").
	WithDefault("1323").
	Required()

var DatabasePath = ferrite.
	String("DATABASE_PATH", "./db.sqlite").
	WithDefault("./db.sqlite").
	Required()

func init() {
	ferrite.Init()
}
