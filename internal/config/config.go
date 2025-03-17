package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

// Config represents the application configuration
type Config struct {
	Port         int    `mapstructure:"port"`          // Port to listen on
	DatabasePath string `mapstructure:"database_path"` // Path to the database file
	BaseDomain   string `mapstructure:"base_domain"`   // Base domain for the shortener
}

// LoadConfig loads configuration from file and environment variables
func LoadConfig() (*Config, error) {
	v := viper.New()

	// Set defaults
	v.SetDefault("port", 8080)
	v.SetDefault("database_path", "./db.sqlite")
	v.SetDefault("base_domain", "https://example.com")

	// Load from config file
	v.SetConfigType("json")

	// Read the config file, ignore errors if file doesn't exist
	if err := v.ReadInConfig(); err != nil {
		// Only log if it's not a file not found error
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Printf("Warning: Error reading config file: %v\n", err)
		}
	}

	// Load from environment variables
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Unmarshal into Config struct
	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// Perform validation
	if err := validateConfig(&config); err != nil {
		return nil, err
	}

	log.Printf("Loaded config: %s", config.String())

	return &config, nil
}

// validateConfig ensures configuration values are valid and applies corrections if needed
func validateConfig(cfg *Config) error {
	// Ensure port is valid
	if cfg.Port <= 0 || cfg.Port > 65535 {
		return fmt.Errorf("invalid port: %d", cfg.Port)
	}

	// Ensure database path is valid
	if cfg.DatabasePath == "" {
		return fmt.Errorf("database path is empty")
	}

	// Ensure base domain is valid
	if cfg.BaseDomain == "" {
		return fmt.Errorf("base domain is empty")
	}

	return nil
}

// String returns a string representation of the config for logging
func (c *Config) String() string {
	return fmt.Sprintf(
		"Config{ Port: %d, DatabasePath: %s, BaseDomain: %s }",
		c.Port,
		c.DatabasePath,
		c.BaseDomain,
	)
}
