package config

import (
	"github.com/MetaerMarket/goconf"
)

func Load(path string) {
	goconf.LoadConfig(path)
}
