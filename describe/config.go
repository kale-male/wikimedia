package describe

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	CacheTTL       time.Duration
	WikiAPIbaseURL string
	RequestTimeout time.Duration
}

func ReadConfig() Config {
	viper.AutomaticEnv()
	var cfg Config

	viper.SetDefault("CACHE_DURATION", 15*time.Minute)
	cfg.CacheTTL = viper.GetDuration("CACHE_DURATION")

	viper.SetDefault("WIKI_API_URL", "https://en.wikipedia.org")
	cfg.WikiAPIbaseURL = viper.GetString("WIKI_API_URL")

	viper.SetDefault("REQUEST_DURATION", 10*time.Second)
	cfg.RequestTimeout = viper.GetDuration("REQUEST_DURATION")

	return cfg
}
