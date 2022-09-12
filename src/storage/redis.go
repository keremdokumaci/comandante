package storage

import (
	"github.com/go-redis/redis/v8"
	"github.com/keremdokumaci/comandante/src/models"
)

type RedisStorage struct {
	Options *redis.Options
}

func NewRedisStorage(options *redis.Options) *RedisStorage {
	return &RedisStorage{
		Options: options,
	}
}

func (rs *RedisStorage) Write(key string, value string) error {
	return nil
}

func (rs *RedisStorage) GetAll() (models.ArrConfigurationVariable, error) {
	return nil, nil
}
