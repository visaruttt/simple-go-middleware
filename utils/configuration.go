package utils

import (
	"os"
	"strconv"
)

func GetConfig(paramKey string) string {
	prod := os.Getenv("PROD")
	isProd, _ := strconv.ParseBool(prod)
	if !isProd {
		return os.Getenv(paramKey+"_TEST")
	}
	return os.Getenv(paramKey+"_PROD")
}
