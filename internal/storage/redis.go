package storage

import (
	"github.com/go-redis/redis"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(address, password string) (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0, // Redis veritabanı seçimi
	})

	// Ping Redis sunucusuna bağlantıyı test eder
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return &RedisClient{client}, nil
}

func (rc *RedisClient) Set(key string, value interface{}) error {
	err := rc.client.Set(key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (rc *RedisClient) Get(key string) (string, error) {
	val, err := rc.client.Get(key).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	}
	return val, nil
}

func (rc *RedisClient) Delete(key string) error {
	err := rc.client.Del(key).Err()
	if err != nil {
		return err
	}
	return nil
}
