package redis

import (
	"os"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
)

var redisClient *redis.Client
var redisAddr string
var redisPassword string
var redisDb int

// GetClient get redis client.
func GetClient() *redis.Client {
	if redisClient == nil {
		_ = initEnv()

		redisClient = redis.NewClient(&redis.Options{
			Addr:     redisAddr,
			Password: redisPassword,
			DB:       redisDb,
		})
	}

	return redisClient
}

func loadEnv() error {
	err := godotenv.Load("/go/src/github.com/ph-piment/golang-pkg/.env")
	if err != nil {
		return err
	}
	return nil
}

func getDb() (int, error) {
	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		return 0, err
	}
	return db, nil
}

func initEnv() error {
	err := loadEnv()
	if err != nil {
		return err
	}

	redisAddr = os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT")
	redisPassword = os.Getenv("REDIS_PASSWORD")
	redisDb, err = getDb()
	if err != nil {
		return err
	}
	return nil
}
