package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// Config holds all application configuration.
type Config struct {
	Rakuten RakutenConfig `mapstructure:"rakuten"`
}

type RakutenConfig struct {
	AppID       string `mapstructure:"app_id"`
	AffiliateID string `mapstructure:"affiliate_id"`
	AccessKey   string `mapstructure:"access_key"` // for Ichiba OpenAPI
	Origin      string `mapstructure:"origin"`     // e.g. https://www.yomitaku.com
}

// Load reads config from file and environment variables via viper.
// Priority: env vars > config file > defaults.
func Load() (*Config, error) {
	v := viper.New()

	// Config file
	home, err := os.UserHomeDir()
	if err == nil {
		cfgPath := filepath.Join(home, ".config", "raku-cli")
		v.SetConfigName("config")
		v.SetConfigType("yaml")
		v.AddConfigPath(cfgPath)
		// Ignore file-not-found errors; env vars may suffice
		_ = v.ReadInConfig()
	}

	// Environment variable bindings
	v.SetEnvPrefix("")
	_ = v.BindEnv("rakuten.app_id", "RAKUTEN_APP_ID")
	_ = v.BindEnv("rakuten.affiliate_id", "RAKUTEN_AFFILIATE_ID")
	_ = v.BindEnv("rakuten.access_key", "RAKUTEN_ACCESS_KEY")
	_ = v.BindEnv("rakuten.origin", "RAKUTEN_ORIGIN")

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}
	return &cfg, nil
}
