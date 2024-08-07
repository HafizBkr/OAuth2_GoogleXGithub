package auth

import (
    "log"
    "os"
)

func GetEnv(key string) string {
    value, exists := os.LookupEnv(key)
    if !exists {
        log.Fatalf("%s not defined in .env file", key)
    }
    return value
}
