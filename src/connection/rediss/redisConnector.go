package rediss

import (
	"context"
	"os"
	"time"
	"user-service/src/helpers"

	"github.com/go-redis/redis/v8"
)

var log = helpers.GetLogger()

// RedisInstance func
func RedisInstance() *redis.Client {

	// Load environment variables
	redisUrl, err := helpers.GetEnvStringVal("REDIS_URL")
	if err != nil {
		log.Error("Failed to load environment variable : REDIS_URL")
		println("Failed to load environment variable : REDIS_URL " + err.Error())
		os.Exit(1)
	}

	redisPassword, err := helpers.GetEnvStringVal("REDIS_PASSWORD")
	if err != nil {
		log.Error("Failed to load environment variable : REDIS_PASSWORD")
		println("Failed to load environment variable : REDIS_PASSWORD " + err.Error())
		log.Debug(err.Error())
		os.Exit(1)
	}

	log.Info("Connecting to Redis Instance : " + redisUrl)
	println("Connecting to Redis Instance : " + redisUrl)
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		Password: redisPassword,
		DB:       0, // use default DB
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		println("Failed to connect to Redis Instance!!!")
		log.Fatal(err)
		println(err)
	} else {
		log.Info("Connected to Redis Instance!!!")
		println("Connected to Redis Instance!!!")
	}

	return rdb
}

var redisClient *redis.Client = RedisInstance()
var ctx = context.Background()

func GetRedisClientConnection() *redis.Client {
	return redisClient
}

func Set(key string, value string) {
	err := GetRedisClientConnection().Set(ctx, key, value, 0).Err()
	if err != nil {
		log.Error("Error in Redis Set.")
		println("Error in Redis Set.")
		log.Debug(err.Error())
	}
}

func HSet(key string, values []string) {
	err := GetRedisClientConnection().HSet(ctx, key, values).Err()
	if err != nil {
		log.Error("Error in Redis HSet.")
		println("Error in Redis HSet.")
		log.Debug(err.Error())
	}
}

func HGetAll(key string) *redis.StringStringMapCmd {
	return GetRedisClientConnection().HGetAll(ctx, key)

}

func Get(key string) (string, error) {
	val, err := GetRedisClientConnection().Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func SetWithExp(key string, value string, exp time.Duration) error {
	err := GetRedisClientConnection().Set(ctx, key, value, exp).Err()
	if err != nil {
		return err
	}
	return nil
}

func Delete(key string) (int64, error) {
	val, err := GetRedisClientConnection().Del(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return val, nil
}

func GetAllKeys() ([]string, error) {
	keys, err := GetRedisClientConnection().Keys(ctx, "*").Result()
	if err != nil {
		return nil, err
	}
	return keys, nil
}
