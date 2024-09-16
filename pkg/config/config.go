package config

import (
	"fmt"
	"os"
	"path/filepath"
)

var AppConfig *Config
var jwtSecret = "your-256-bit-secret" // Replace with your own secret

type Config struct {
	AppName   string
	JWTSecret string
}

func New() *Config {

	appPath := os.Args[0]
	appName := filepath.Base(appPath)
	fmt.Printf("Application name: %s\n", appName)

	AppConfig = &Config{
		AppName:   appName,
		JWTSecret: jwtSecret,
	}

	return AppConfig
}
