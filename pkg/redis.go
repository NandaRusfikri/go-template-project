package pkg

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

type Redis struct {
	Config *RedisConfig
	Client *redis.Client
}

func InitRedis(config *RedisConfig) *Redis {
	return &Redis{
		Config: config,
	}
}

func (r *Redis) New() (err error) {
	log.Info("Redis - New() - starting...")

	if r.Client != nil {
		_ = r.Client.Close()
	}

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", r.Config.Host, r.Config.Port),
		Password: r.Config.Password,
		DB:       r.Config.DB,
	})

	result, err := client.Ping().Result()
	if err != nil {
		return err
	}
	log.Info("Redis - New() - result Ping: ", result)

	r.Client = client
	log.Info("Redis - New() - finished.")
	return nil
}

func (r *Redis) Get(key string) (res string, err error) {
	if key == "" {
		return res, errors.New("key required")
	}

	res, err = r.Client.Get(key).Result()
	if err != nil {
		return res, err
	}

	return res, nil
}

func GetRedisConfig() *RedisConfig {
	redisHost := os.Getenv("REDIS_HOST")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		log.Error("GetRedisConfig() - error while parsing redis db: ", err)
	}
	redisPORT, err := strconv.Atoi(os.Getenv("REDIS_PORT"))
	if err != nil {
		log.Error("GetRedisConfig() - error while parsing  redisPORT: ", err)
	}
	config := &RedisConfig{
		Host:     redisHost,
		Port:     redisPORT,
		Password: redisPassword,
		DB:       redisDB,
	}

	return config
}
