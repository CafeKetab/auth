package crypto

type Config struct {
	Secret string `koanf:"secret"`
	Salt   string `koanf:"salt"`
}
