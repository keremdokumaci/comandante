package storage

import "github.com/keremdokumaci/comandante/src/models"

type RedisStorage struct{}

func NewRedisStorage() *RedisStorage {
	return &RedisStorage{}
}

func (rs *RedisStorage) Write(key string, value string) error {
	return nil
}

func (rs *RedisStorage) GetAll() (models.ArrConfigurationVariable, error) {
	return nil, nil
}
