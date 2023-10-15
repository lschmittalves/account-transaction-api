package config

type DBConfig struct {
	User     string
	Password string
	Driver   string
	Name     string
	Host     string
	Port     string
}

func LoadDBConfig() DBConfig {
	return DBConfig{
		User:     getEnvString("DB_USER"),
		Password: getEnvString("DB_PASSWORD"),
		Name:     getEnvString("DB_NAME"),
		Host:     getEnvString("DB_HOST"),
		Port:     getEnvString("DB_PORT"),
	}
}
