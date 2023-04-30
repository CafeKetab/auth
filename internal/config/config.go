package config

import "github.com/CafeKetab/auth-go/pkg/logger"

type Config struct {
	Logger *logger.Config `koanf:"log"`
}
