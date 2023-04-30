package config

import (
	"time"

	"github.com/CafeKetab/auth-go/pkg/logger"
	"github.com/CafeKetab/auth-go/pkg/token"
)

func Default() *Config {
	return &Config{
		Logger: &logger.Config{
			Development: true,
			Level:       "debug",
		},
		Token: &token.Config{
			PrivatePem: "-----BEGIN PRIVATE KEY-----" +
				"MC4CAQAwBQYDK2VwBCIEINyMNS8h9M9HO73Tg1BPr53p//qlqylO+wPKN8GrlsX7" +
				"-----END PRIVATE KEY-----",
			PublicPem: "-----BEGIN PUBLIC KEY-----" +
				"MCowBQYDK2VwAyEAqQsZ5iRNP3kdpNn3V/db9o/WkYHY8kkwQqCZGcDvJ+g=" +
				"-----END PUBLIC KEY-----",
			Expiration: 30 * time.Minute,
		},
	}
}
