package config

import (
	"fmt"
	"os"
)
import "github.com/joho/godotenv"

func LoadEnv(env string) {
	err := godotenv.Load(env)
	if err != nil {
		panic(err)
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
func GetMongoConnectionString() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s",
		ApiConfig.MongoUser,
		ApiConfig.MongoPass,
		ApiConfig.MongoHost,
		ApiConfig.MongoPort)

}
