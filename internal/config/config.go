package config

import (
	"github.com/CafeKetab/auth/internal/ports/grpc"
	"github.com/CafeKetab/auth/pkg/crypto"
	"github.com/CafeKetab/auth/pkg/logger"
	"github.com/CafeKetab/auth/pkg/token"
)

type Config struct {
	Logger *logger.Config `koanf:"logger"`
	Grpc   *grpc.Config   `koanf:"grpc"`
	Token  *token.Config  `koanf:"token"`
	Crypto *crypto.Config `koanf:"crypto"`
}
