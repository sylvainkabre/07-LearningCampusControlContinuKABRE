package config

import (
	"os"
)

func GetJWTSecretKey() string {
	return os.Getenv("JWT_SIGNATURE_KEY")
}
