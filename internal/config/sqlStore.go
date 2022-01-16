package config

type sqlStore struct {
	DatabaseURL string `env:"DATABASE_URL"`
}
