package config

type CacheConfig struct {
	Password string
	Host     string
	Port     string
	TTL      CacheConfigTTL
}

type CacheConfigTTL struct {
	Transactions int
}

func LoadCacheConfig() CacheConfig {
	return CacheConfig{
		Password: getEnvString("REDIS_PASSWORD"),
		Host:     getEnvString("REDIS_HOST"),
		Port:     getEnvString("REDIS_PORT"),
		TTL: CacheConfigTTL{
			Transactions: getEnvInt("REDIS_TTL_TRANSACTION"),
		},
	}
}
