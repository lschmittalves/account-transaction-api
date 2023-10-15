package config

import (
	"os"
	"strconv"
)

func getEnvString(key string) string {
	return os.Getenv(key)
}

func getEnvInt(key string) int {
	s := getEnvString(key)
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return v
}

func getEnvBool(key string) bool {
	s := getEnvString(key)
	v, err := strconv.ParseBool(s)
	if err != nil {
		panic(err)
	}
	return v
}
