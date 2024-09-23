package config

type Config struct {
	Port string `env:"PORT"`

	DB_Host     string `env:"DB_HOST"`
	DB_User     string `env:"DB_USER"`
	DB_Password string `env:"DB_PASSWORD"`
	DB_Name     string `env:"DB_NAME"`
	DB_Port     string `env:"DB_PORT"`
	DB_SSLMode  string `env:"DB_SSLMODE"`
}
