package config

import (
	"github.com/dogmatiq/ferrite"
)

var BaseURL = ferrite.
	URL("BASE_DOMAIN", "http://example.org").
	Required()

func init() {
	ferrite.Init()
}
