package config

import (
	"github.com/CafeKetab/auth/pkg/crypto"
	"github.com/CafeKetab/auth/pkg/logger"
	"github.com/CafeKetab/auth/pkg/token"
)

type Config struct {
	Logger *logger.Config `koanf:"logger"`
	Token  *token.Config  `koanf:"token"`
	Crypto *crypto.Config `koanf:"crypto"`
}
