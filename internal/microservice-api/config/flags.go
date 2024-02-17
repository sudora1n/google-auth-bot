package config

import (
	"github.com/spf13/pflag"
)

var configPath *string

func ParseFlags() {
	configPath = pflag.StringP("config", "c", "./config/api.env", "Path to the config file")
}
