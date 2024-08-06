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

func init() {
	ferrite.Init()
}
