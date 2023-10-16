package config

type DBConfig struct {
	User               string
	Password           string
	Driver             string
	Name               string
	Host               string
	Port               string
	MaxLifetime        int
	MaxConnections     int
	MaxIdleConnections int
	LogMode            bool
}

func LoadDBConfig() DBConfig {
	return DBConfig{
		User:               getEnvString("DB_USER"),
		Password:           getEnvString("DB_PASSWORD"),
		Driver:             getEnvString("DB_DRIVER"),
		Name:               getEnvString("DB_NAME"),
		Host:               getEnvString("DB_HOST"),
		Port:               getEnvString("DB_PORT"),
		MaxLifetime:        getEnvInt("DB_MAX_LIFETIME"),
		MaxConnections:     getEnvInt("DB_MAX_OPEN_CONNS"),
		MaxIdleConnections: getEnvInt("DB_MAX_IDLE_CONNS"),
		LogMode:            getEnvBool("DB_LOG"),
	}
}
